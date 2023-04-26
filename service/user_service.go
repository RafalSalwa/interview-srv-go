package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	mySql "github.com/RafalSalwa/interview-app-srv/sql"
	"github.com/RafalSalwa/interview-app-srv/util"
	"strconv"
	"time"

	phpserialize "github.com/kovetskiy/go-php-serialize"
)

type UserSqlService interface {
	GetUserById(id int64) (user *models.User, err error)
	GetUserByUsername(username string) (user *models.User, err error)
	UpdateUser(user *models.User) (err error)
	UpdateUserFirebaseToken(user *models.User) (err error)
	LoginUser(username string) (*models.User, error)
	UpdateUserPassword(user *models.User) (err error)
	CreateUser(user *models.User) (id int64, err error)
	CreateUserDevice(userDevice *models.UserDevice) (id int64, err error)
	GetDevicesByUserId(id int64) (userDevice *models.UserDevices, err error)
	GetLatestDevice(id int64) (userDevice *models.UserDevice, err error)
}

type SqlServiceImpl struct {
	db mySql.DB
}

func NewMySqlService(db mySql.DB) *SqlServiceImpl {
	return &SqlServiceImpl{db}
}

func (s *SqlServiceImpl) GetUserById(id int64) (user *models.User, err error) {
	user = &models.User{}
	row := s.db.QueryRow("SELECT id,username,first_name,last_name,roles as roles_json,firebase_token FROM `user` WHERE deleted_at IS NULL AND id=?", id)
	err = row.Scan(&user.Id,
		&user.Username,
		&user.Firstname,
		&user.Lastname,
		&user.RolesJson,
		&user.FirebaseToken)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	roles, err := phpserialize.Decode(user.RolesJson)

	if err != nil {
		return nil, err
	}

	v, ok := roles.(map[interface{}]interface{})
	if ok {
		for _, s := range v {
			user.Roles = append(user.Roles, fmt.Sprintf("%v", s))
		}
	}

	return user, nil
}

func (s *SqlServiceImpl) GetUserByUsername(username string) (user *models.User, err error) {
	user = &models.User{}
	row := s.db.QueryRow("SELECT id,username,first_name,last_name,roles as roles_json,firebase_token FROM `user` WHERE deleted_at IS NULL AND username=?", username)
	err = row.Scan(&user.Id,
		&user.Username,
		&user.Firstname,
		&user.Lastname,
		&user.RolesJson,
		&user.FirebaseToken)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	roles, err := phpserialize.Decode(user.RolesJson)

	if err != nil {
		return nil, err
	}

	v, ok := roles.(map[interface{}]interface{})
	if ok {
		for _, s := range v {
			user.Roles = append(user.Roles, fmt.Sprintf("%v", s))
		}
	}

	return user, nil
}

func (s *SqlServiceImpl) UpdateUser(user *models.User) (err error) {
	sqlStatement := "UPDATE user SET first_name=?, last_name=? WHERE id=?;"
	_, err = s.db.ExecContext(getContext(), sqlStatement, &user.Firstname, &user.Lastname, user.Id)

	if err != nil {
		return err
	}

	return nil
}

func (s *SqlServiceImpl) UpdateUserFirebaseToken(user *models.User) (err error) {
	sqlStatement := "UPDATE `user` SET firebase_token=? WHERE id=?;"
	_, err = s.db.ExecContext(getContext(), sqlStatement, user.FirebaseToken, user.Id)

	if err != nil {
		return err
	}

	return nil
}

func (s *SqlServiceImpl) LoginUser(username string) (*models.User, error) {
	user := models.User{}

	row := s.db.QueryRow("SELECT id,username,password,first_name,last_name,roles as roles_json,firebase_token FROM `user` WHERE (username=? OR email=?)", username, username)
	err := row.Scan(&user.Id,
		&user.Username,
		&user.Password,
		&user.Firstname,
		&user.Lastname,
		&user.RolesJson,
		&user.FirebaseToken)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	roles, err := phpserialize.Decode(user.RolesJson)

	if err != nil {
		return nil, err
	}

	v, ok := roles.(map[interface{}]interface{})
	if ok {
		for _, s := range v {
			user.Roles = append(user.Roles, fmt.Sprintf("%v", s))
		}
	}

	return &user, nil
}

func (s *SqlServiceImpl) UpdateUserPassword(user *models.User) (err error) {
	sqlStatement := "UPDATE `user` SET password=? WHERE id=?;"
	_, err = s.db.ExecContext(getContext(), sqlStatement, user.Password, user.Id)

	if err != nil {
		return err
	}

	return nil
}

func (s *SqlServiceImpl) CreateUser(user *models.User) (id int64, err error) {
	sqlStatement := "INSERT INTO `user` (`username`, `password`, `enabled`, `roles`) VALUES (?,?,?,?);"
	rows, err := s.db.ExecContext(getContext(),
		sqlStatement,
		user.Username,
		user.Password,
		1,
		user.RolesJson)
	if err != nil {
		return 0, err
	}

	id, err = rows.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *SqlServiceImpl) CreateUserDevice(userDevice *models.UserDevice) (id int64, err error) {
	sqlStatement :=
		"INSERT INTO `user_device` " +
			"(`user_id`,`firebase_token`,`os_type`,`sdk_version`,`model`,`brand`,`last_login`,`created_at`,`deleted_at`) " +
			"VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE last_login=NOW(),sdk_version=?, model=?, brand=?"
	rows, err := s.db.ExecContext(
		getContext(),
		sqlStatement,
		&userDevice.UserId,
		&userDevice.FirebaseToken,
		&userDevice.OsType,
		&userDevice.SdkVersion,
		&userDevice.Model,
		&userDevice.Brand,
		&userDevice.LastLogin,
		&userDevice.CreatedAt,
		&userDevice.DeletedAt,
		&userDevice.SdkVersion,
		&userDevice.Model,
		&userDevice.Brand,
	)

	if err != nil {
		return 0, err
	}

	id, err = rows.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *SqlServiceImpl) GetDevicesByUserId(id int64) (devices *models.UserDevices, err error) {
	devices = &models.UserDevices{}

	rows, err := s.db.QueryContext(getContext(), "SELECT `id`,`user_id`,`firebase_token`,`os_type`,`sdk_version`,`model`,`brand`,`last_login`,`created_at`,`deleted_at` "+
		"FROM `user_device` WHERE `user_id`=?", id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var device models.UserDevice
		err := rows.Scan(
			&device.Id,
			&device.UserId,
			&device.FirebaseToken,
			&device.OsType,
			&device.SdkVersion,
			&device.Model,
			&device.Brand,
			&device.LastLogin,
			&device.CreatedAt,
			&device.DeletedAt,
		)

		if err != nil {
			return nil, err
		}

		devices.Items = append(devices.Items, device)

	}
	return devices, nil
}

func (s *SqlServiceImpl) GetLatestDevice(id int64) (device *models.UserDevice, err error) {
	device = &models.UserDevice{}

	row := s.db.QueryRowContext(getContext(), "SELECT `id`,`user_id`,`firebase_token`,`os_type`,`sdk_version`,`model`,`brand`,`last_login`,`created_at`,`deleted_at` "+
		"FROM `user_device` WHERE `user_id`=? ORDER BY created_at DESC limit 1", id)
	if err != nil {
		return nil, err
	}

	err = row.Scan(
		&device.Id,
		&device.UserId,
		&device.FirebaseToken,
		&device.OsType,
		&device.SdkVersion,
		&device.Model,
		&device.Brand,
		&device.LastLogin,
		&device.CreatedAt,
		&device.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return device, nil
}

func getContext() context.Context {
	ctx := context.Background()
	var timeout, err = strconv.Atoi(util.Env("SQL_REQUEST_TIMEOUT_SECONDS", "60"))
	if err == nil {
		ctx, _ = context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	}
	return ctx
}
