package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/RafalSalwa/interview-app-srv/internal/generator"
	"github.com/RafalSalwa/interview-app-srv/internal/password"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"

	"github.com/RafalSalwa/interview-app-srv/internal/util"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	mySql "github.com/RafalSalwa/interview-app-srv/sql"
)

type SqlServiceImpl struct {
	db     *mySql.DB
	logger *logger.Logger
}

type UserSqlService interface {
	GetById(id int64) (user *models.UserDBResponse, err error)
	GetByCode(code string) (user *models.UserDBModel, err error)
	UsernameInUse(user *models.CreateUserRequest) bool
	StoreVerificationData(user *models.UserDBModel) bool
	UpdateUser(user *models.UpdateUserRequest) (err error)
	LoginUser(user *models.LoginUserRequest) (*models.UserResponse, error)
	UpdateUserPassword(user *models.UpdateUserRequest) (err error)
	CreateUser(user *models.CreateUserRequest) (*models.UserResponse, error)
}

func NewMySqlService(db *mySql.DB, l *logger.Logger) *SqlServiceImpl {
	return &SqlServiceImpl{db, l}
}

func (s SqlServiceImpl) GetByCode(code string) (*models.UserDBModel, error) {
	user := &models.UserDBModel{}

	row := s.db.QueryRow("SELECT id,verification_code FROM `user` WHERE verification_code = ?", code)
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

func (s *SqlServiceImpl) UsernameInUse(user *models.CreateUserRequest) bool {
	dbUser := &models.UserDBResponse{}
	row := s.db.QueryRow("SELECT id FROM `user` WHERE username=? OR email = ?", user.Username, user.Email)
	err := row.Scan(&dbUser.Id)

	if err == sql.ErrNoRows {
		return false
	}

	if err != nil {
		return false
	}

	return true
}

func (s *SqlServiceImpl) StoreVerificationData(user *models.UserDBModel) bool {
	_, err := s.db.Exec("UPDATE `user` SET is_verified = 1, is_active=1 WHERE id = ?", user.Id)
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

	// roles, err := phpserialize.Decode(user.RolesJson)

	if err != nil {
		return nil, err
	}

	// v, ok := roles.(map[interface{}]interface{})
	// if ok {
	//	for _, s := range v {
	//		user.Roles = append(user.Roles, fmt.Sprintf("%v", s))
	//	}
	//}

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

func (s *SqlServiceImpl) CreateUser(newUserRequest *models.CreateUserRequest) (*models.UserResponse, error) {
	if err := password.Validate(newUserRequest.Password, newUserRequest.PasswordConfirm); err != nil {
		return nil, err
	}

	if s.UsernameInUse(newUserRequest) {
		return nil, errors.New("create user: username already in use")
	}

	roles, err := json.Marshal(models.Roles{Roles: []string{"ROLE_USER"}})
	if err != nil {
		return nil, err
	}

	vcode, err := generator.RandomString(6)
	if err != nil {
		return nil, err
	}

	newUserRequest.Password, err = password.HashPassword(newUserRequest.Password)
	if err != nil {
		return nil, err
	}

	dbUser := &models.UserDBModel{
		Username:         newUserRequest.Username,
		Password:         newUserRequest.Password,
		Email:            newUserRequest.Email,
		Roles:            roles,
		VerificationCode: *vcode,
	}
	ctx := getContext()
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	sqlStatement := "INSERT INTO `user` (`username`, `password`, `email`, `roles`, `verification_code`, `is_verified`,`is_active`) VALUES (?,?,?,?,?,0,1);"
	rows, err := tx.ExecContext(ctx,
		sqlStatement,
		dbUser.Username,
		dbUser.Password,
		dbUser.Email,
		dbUser.Roles,
		dbUser.VerificationCode)

	if err != nil {
		if errTx := tx.Rollback(); errTx != nil {
			return nil, errTx
		}
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	ur := &models.UserResponse{
		Id:       id,
		Username: dbUser.Username,
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
