package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/RafalSalwa/interview-app-srv/internal/generator"
	"github.com/RafalSalwa/interview-app-srv/internal/password"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"

	"github.com/RafalSalwa/interview-app-srv/internal/util"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	mySql "github.com/RafalSalwa/interview-app-srv/sql"

	phpserialize "github.com/kovetskiy/go-php-serialize"
)

type SqlServiceImpl struct {
	db     mySql.DB
	logger *logger.Logger
}

type UserSqlService interface {
	GetById(id int64) (user *models.UserDBResponse, err error)
	GetByCode(code string) (user *models.UserDBModel, err error)
	Exists(user *models.CreateUserRequest) bool
	Veryficate(user *models.UserDBModel) bool
	UpdateUser(user *models.UpdateUserRequest) (err error)
	LoginUser(user *models.LoginUserRequest) (*models.UserResponse, error)
	UpdateUserPassword(user *models.UpdateUserRequest) (err error)
	CreateUser(user *models.CreateUserRequest) (*models.UserResponse, error)
}

func (u SqlServiceImpl) GetByCode(code string) (*models.UserDBModel, error) {
	user := &models.UserDBModel{}
	row := u.db.QueryRow("SELECT id,verification_code FROM `user` WHERE verification_code = ?", code)
	err := row.Scan(&user.Id,
		&user.VerificationCode)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewMySqlService(db mySql.DB, l *logger.Logger) *SqlServiceImpl {
	return &SqlServiceImpl{db, l}
}

func (s *SqlServiceImpl) GetById(id int64) (user *models.UserDBResponse, err error) {
	user = &models.UserDBResponse{}
	row := s.db.QueryRow("SELECT id,username,first_name,last_name,password, roles as roles_json,is_verified, is_active, created_at FROM `user` WHERE is_active = 1 AND id=?", id)
	err = row.Scan(&user.Id,
		&user.Username,
		&user.Firstname,
		&user.Lastname,
		&user.Password,
		&user.RolesJson,
		&user.Verified,
		&user.Active,
		&user.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *SqlServiceImpl) Exists(user *models.CreateUserRequest) bool {
	dbuser := &models.UserDBResponse{}
	row := s.db.QueryRow("SELECT id FROM `user` WHERE username=? OR email = ?", user.Username, user.Email)
	err := row.Scan(&dbuser.Id)

	if err == sql.ErrNoRows {
		return false
	}

	if err != nil {
		return false
	}

	return true
}

func (s *SqlServiceImpl) Veryficate(user *models.UserDBModel) bool {
	fmt.Printf("Veryficate %#v\n", user)
	dbuser := &models.UserDBResponse{}
	_, err := s.db.Exec("UPDATE `user` SET is_verified = 1, is_active=1 WHERE id = ?", user.Id)
	fmt.Printf("Veryficate2 %#v\n", dbuser)
	if err == sql.ErrNoRows {
		return false
	}

	if err != nil {
		return false
	}

	return true
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

func (s *SqlServiceImpl) CreateUser(user *models.CreateUserRequest) (*models.UserResponse, error) {
	roles, _ := json.Marshal(models.Roles{Roles: []string{"ROLE_USER"}})
	vcode, _ := generator.RandomString(6)
	user.Password, _ = password.HashPassword(user.Password)
	dbUser := &models.UserDBModel{
		Username:         user.Username,
		Password:         user.Password,
		Email:            user.Email,
		Roles:            roles,
		VerificationCode: *vcode,
	}

	sqlStatement := "INSERT INTO `user` (`username`, `password`, `email`, `roles`, `verification_code`, `is_verified`,`is_active`) VALUES (?,?,?,?,?,0,1);"
	rows, err := s.db.ExecContext(getContext(),
		sqlStatement,
		dbUser.Username,
		dbUser.Password,
		dbUser.Email,
		dbUser.Roles,
		dbUser.VerificationCode)

	if err != nil {
		fmt.Println(err)
		s.logger.Error().Err(err).Msg("")
		return nil, err
	}

	id, err := rows.LastInsertId()
	ur := &models.UserResponse{
		Id:        id,
		Username:  dbUser.Username,
		CreatedAt: nil,
	}
	if err != nil {
		s.logger.Error().Err(err).Msg("")
		return nil, err
	}

	return ur, nil
}

func getContext() context.Context {
	ctx := context.Background()
	var timeout, err = strconv.Atoi(util.Env("SQL_REQUEST_TIMEOUT_SECONDS", "60"))
	if err == nil {
		ctx, _ = context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	}
	return ctx
}
