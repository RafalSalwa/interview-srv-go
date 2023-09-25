package models

import (
	"encoding/json"
	"fmt"
	"strings"

	intrvproto "github.com/RafalSalwa/interview-app-srv/proto/grpc"

	"github.com/RafalSalwa/interview-app-srv/pkg/jwt"
	"github.com/jinzhu/copier"
	phpserialize "github.com/kovetskiy/go-php-serialize"
)

func (r *UserResponse) FromDBResponse(user *UserDBResponse) error {
	err := copier.Copy(r, &user)
	if err != nil {
		return fmt.Errorf("from db to response error: %w", err)
	}
	return nil
}

func (m *UserDBModel) FromCreateUserReq(cur *SignUpUserRequest) error {
	err := copier.Copy(m, &cur)
	m.Username = cur.Email
	if err != nil {
		return fmt.Errorf("from create to db model error: %w", err)
	}

	return nil
}

func (r *UserResponse) FromProtoUserDetails(pbu *intrvproto.UserDetails) {
	r.Id = pbu.GetId()
	r.Username = pbu.GetUsername()
	r.Firstname = pbu.GetFirstname()
	r.Lastname = pbu.GetLastname()
	r.Email = pbu.GetEmail()
	r.Verified = pbu.GetVerified()
	r.Active = pbu.GetActive()
	r.CreatedAt = pbu.GetCreatedAt().AsTime()

	ll := pbu.GetLastLogin().AsTime()
	r.LastLogin = &ll
}

func (r *UserResponse) FromDBModel(um *UserDBModel) error {
	err := copier.Copy(r, &um)
	r.Username = um.Username
	r.CreatedAt = um.CreatedAt
	if err != nil {
		return fmt.Errorf("from response to db error: %w", err)
	}
	return nil
}

func (r *UserResponse) AssignTokenPair(tp *jwt.TokenPair) {
	r.AccessToken = tp.AccessToken
	r.RefreshToken = tp.RefreshToken
}

func (r *UserResponse) FromProtoSignIn(pbu *intrvproto.SignInUserResponse) {
	r.AccessToken = pbu.AccessToken
	r.RefreshToken = pbu.RefreshToken
}
func (r *UserResponse) FromProtoSignUp(pbu *intrvproto.SignUpUserResponse) error {
	r.Username = pbu.GetUsername()
	r.VerificationCode = pbu.GetVerificationToken()
	c := pbu.GetCreatedAt().AsTime()
	r.CreatedAt = c
	return nil
}

func (r *UserResponse) FromProtoUserResponse(pu *intrvproto.UserResponse) error {
	err := copier.Copy(r, &pu)
	r.Id = pu.User.Id
	r.Username = pu.User.Username

	if err != nil {
		return fmt.Errorf("from response to db error: %w", err)
	}
	return nil
}

func (r *UserDBResponse) FromProtoUserResponse(pw *intrvproto.UserResponse) error {
	err := copier.Copy(r, pw.User)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserDBResponse) FromProtoUserDetails(pw *intrvproto.UserDetails) error {
	err := copier.Copy(r, pw)
	if err != nil {
		return err
	}
	return nil
}
func (r SignUpUserRequest) ToProto() *intrvproto.SignUpUserInput {
	u := &intrvproto.SignUpUserInput{}

	u.Email = r.Email
	u.Password = r.Password
	u.PasswordConfirm = r.PasswordConfirm

	return u
}
func (r *UserDBResponse) FromDBModel(user *UserDBResponse) error {
	err := copier.Copy(r, user)
	if err != nil {
		return err
	}
	return nil
}
func (m *UserMongoModel) FromDBModel(user *UserDBModel) error {
	err := copier.Copy(m, user)
	if err != nil {
		return err
	}
	return nil
}
func getRolesList(r string) []string {
	i := strings.Index(r, "roles")
	type RoleItem struct {
		Roles []string
	}
	var val RoleItem

	if i > -1 {
		err := json.Unmarshal([]byte(r), &val)
		if err != nil {
			fmt.Println(err)
		}
		return val.Roles
	}

	roles, err := phpserialize.Decode(r)

	if err != nil {
		return nil
	}
	v, ok := roles.(map[interface{}]interface{})
	decodedRoles := make([]string, len(v))
	if ok {
		for _, s := range v {
			decodedRoles = append(decodedRoles, fmt.Sprintf("%v", s))
		}
	}

	return decodedRoles
}
