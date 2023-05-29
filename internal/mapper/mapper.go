package mapper

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/jinzhu/copier"
	phpserialize "github.com/kovetskiy/go-php-serialize"
)

const dbTimeFormat = "2006-01-02 15:04:05"

func UserCreateRequestToDBModel(ur *models.CreateUserRequest) *models.UserDBModel {
	um := &models.UserDBModel{}
	err := copier.Copy(um, ur)
	if err != nil {
		return nil
	}
	return um
}

func MapUserDBModelToUserResponse(user *models.UserDBModel) *models.UserResponse {
	userResponse := &models.UserResponse{}

	err := copier.Copy(&userResponse, &user)
	if err != nil {
		return nil
	}

	return userResponse
}

func MapUserDBResponseToUserResponse(user *models.UserDBResponse) *models.UserResponse {
	userResponse := &models.UserResponse{}
	err := copier.Copy(&userResponse, &user)
	if err != nil {
		return nil
	}

	userResponse.Roles = user.RolesJson

	return userResponse
}

func GetRolesList(r string) []string {
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
