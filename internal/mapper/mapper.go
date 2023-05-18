package mapper

import (
	"encoding/json"
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/jinzhu/copier"
	phpserialize "github.com/kovetskiy/go-php-serialize"
	"strings"
)

const dbTimeFormat = "2006-01-02 15:04:05"

func MapUserDBModelToUserResponse(user *models.UserDBModel) *models.UserResponse {
	userResponse := &models.UserResponse{}

	err := copier.Copy(&userResponse, &user)
	if err != nil {
		return nil
	}

	userResponse.Roles = getRolesList(user.RolesJson)

	return userResponse
}

func MapUserDBResponseToUserResponse(user *models.UserDBResponse) *models.UserResponse {
	userResponse := &models.UserResponse{}
	err := copier.Copy(&userResponse, &user)
	if err != nil {
		return nil
	}

	userResponse.Roles = getRolesList(user.RolesJson)

	return userResponse
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
