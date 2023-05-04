package services

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	mySql "github.com/RafalSalwa/interview-app-srv/sql"
	"github.com/RafalSalwa/interview-app-srv/util"

	phpserialize "github.com/kovetskiy/go-php-serialize"
)

type SqlServiceImpl struct {
	db mySql.DB
}

type UserSqlService interface {
	GetUserById(id int64) (user *models.UserDBResponse, err error)
	GetUserByUsername(username string) (user *models.UserResponse, err error)
	UpdateUser(user *models.UpdateUserRequest) (err error)
	LoginUser(user *models.LoginUserRequest) (*models.UserResponse, error)
	UpdateUserPassword(user *models.UpdateUserRequest) (err error)
	CreateUser(user *models.CreateUserRequest) (id int64, err error)
}

func NewMySqlService(db mySql.DB) *SqlServiceImpl {
	return &SqlServiceImpl{db}
}

func (s *SqlServiceImpl) GetUserById(id int64) (user *models.UserDBResponse, err error) {
	user = &models.UserDBResponse{}
	row := s.db.QueryRow("SELECT id,username,first_name,last_name,roles as roles_json,is_verified, 1 FROM `user` WHERE is_active = 1 AND id=?", id)
	err = row.Scan(&user.Id,
		&user.Username,
		&user.Firstname,
		&user.Lastname,
		&user.RolesJson,
		&user.Verified,
		&user.Active)

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

func (s *SqlServiceImpl) GetUserByUsername(username string) (user *models.UserResponse, err error) {
	user = &models.UserResponse{}
	row := s.db.QueryRow("SELECT id,username,roles as roles_json FROM `user` WHERE is_active = 1 AND username=?", username)
	err = row.Scan(&user.Id,
		&user.Username,
		&user.RolesJson)

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

func (s *SqlServiceImpl) UpdateUser(user *models.UpdateUserRequest) (err error) {
	sqlStatement := "UPDATE user SET first_name=?, last_name=? WHERE id=?;"
	_, err = s.db.ExecContext(getContext(), sqlStatement, &user.Firstname, &user.Lastname, user.Id)

	if err != nil {
		return err
	}

	return nil
}

func (s *SqlServiceImpl) LoginUser(u *models.LoginUserRequest) (*models.UserResponse, error) {
	user := models.UserResponse{}

	row := s.db.QueryRow("SELECT id,username,first_name,last_name,roles as roles_json FROM `user` WHERE (username=? OR email=?) AND (is_active = 1 AND is_verified = 1)", u.Username, u.Email)
	err := row.Scan(&user.Id,
		&user.Username,
		&user.Firstname,
		&user.RolesJson)

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

func (s *SqlServiceImpl) UpdateUserPassword(user *models.UpdateUserRequest) (err error) {
	sqlStatement := "UPDATE `user` SET password=? WHERE id=?;"
	_, err = s.db.ExecContext(getContext(), sqlStatement, user.Password, user.Id)

	if err != nil {
		return err
	}

	return nil
}

func (s *SqlServiceImpl) CreateUser(user *models.CreateUserRequest) (id int64, err error) {
	sqlStatement := "INSERT INTO `user` (`username`, `email`,`password`, `is_active`, `roles`) VALUES (?,?,?,1,{\"roles\": \"ROLE_USER\"});"
	rows, err := s.db.ExecContext(getContext(),
		sqlStatement,
		user.Username,
		user.Password,
		1)
	if err != nil {
		return 0, err
	}

	id, err = rows.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func getContext() context.Context {
	ctx := context.Background()
	var timeout, err = strconv.Atoi(util.Env("SQL_REQUEST_TIMEOUT_SECONDS", "60"))
	if err == nil {
		ctx, _ = context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	}
	return ctx
}
