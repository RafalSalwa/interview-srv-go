package models

import (
	"encoding/json"
	"fmt"
	intrvproto "github.com/RafalSalwa/interview-app-srv/proto/grpc"
	"strings"

	"github.com/RafalSalwa/interview-app-srv/pkg/jwt"
	"github.com/jinzhu/copier"
	phpserialize "github.com/kovetskiy/go-php-serialize"
)

func (r *UserResponse) FromDBResponse(user *UserDBResponse) error {
	err := copier.Copy(r, &user)
	if err != nil {
		return fmt.Errorf("from db to response error: %w", err)
	}

	r.Roles = user.RolesJson

	return nil
}

func (m *UserDBModel) FromCreateUserReq(cur *CreateUserRequest) error {
	err := copier.Copy(m, &cur)
	if err != nil {
		return fmt.Errorf("from create to db model error: %w", err)
	}

	return nil
}

func (r *UserResponse) FromDBModel(um *UserDBModel) error {
	err := copier.Copy(r, &um)
	if err != nil {
		return fmt.Errorf("from response to db error: %w", err)
	}
	r.Roles = string(um.Roles)
	return nil
}

func (r *UserResponse) AssignTokenPair(tp *jwt.TokenPair) {
	r.Token = tp.AccessToken
	r.RefreshToken = tp.RefreshToken
}

func (r *UserResponse) FromProtoSignIn(pbu *intrvproto.SignInUserResponse) error {
	err := copier.Copy(r, &pbu)
	if err != nil {
		return fmt.Errorf("from response to db error: %w", err)
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
