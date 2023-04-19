package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/RafalSalwa/interview/jwt"
	"github.com/RafalSalwa/interview/mapper"
	"github.com/RafalSalwa/interview/model"
	"github.com/RafalSalwa/interview/utils"
	"github.com/RafalSalwa/interview/utils/logger"
	"github.com/gorilla/mux"
	phpserialize "github.com/kovetskiy/go-php-serialize"
)

func (h handler) getUsersId() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := mux.Vars(r)["id"]

		intId, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondBadRequest(w, "Wrong paramater Type. Required int")
			return
		}

		user, err := h.userSqlService.GetUserById(intId)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		if user == nil {
			respondNotFound(w)
			return
		}

		userResponse(user, w, r)
	}

}

func (h handler) postUsersIdEdit() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondBadRequest(w, "")
			return
		}

		pathUserId := mux.Vars(r)["id"]
		formFirstname := r.PostForm.Get("firstname")
		formLastname := r.PostForm.Get("lastname")

		if len(formFirstname) == 0 || len(formLastname) == 0 {
			respondBadRequest(w, "")
			return
		}

		userId, err := strconv.ParseInt(pathUserId, 10, 64)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondBadRequest(w, "")
			return
		}

		user, err := h.userSqlService.GetUserById(userId)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		if user == nil {
			respondNotFound(w)
			return
		}

		user.Firstname = &formFirstname
		user.Lastname = &formLastname

		err = h.userSqlService.UpdateUser(user)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		userResponse(user, w, r)
	}

}

func (h handler) getUsersIdNotifications() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondBadRequest(w, "")
			return
		}

		user, err := h.userSqlService.GetUserById(userId)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		if user != nil && user.FirebaseToken != nil {
			notifications, err := h.sqlService.GetNotificationsByUserId(user.Id)
			if err != nil {
				logger.Log(err.Error(), logger.Error)
				respondInternalServerError(w)
				return
			}

			notificationsResponse(notifications, w, r)
			return
		}

		respondNotFound(w)
	}

}

func (h handler) postUsersIdChangePassword() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondBadRequest(w, "")
			return
		}

		formUsername := r.PostForm.Get("username")
		formOldPassword := r.PostForm.Get("oldPassword")
		formNewPassword := r.PostForm.Get("newPassword")

		user, err := h.userSqlService.LoginUser(formUsername)

		if formOldPassword == formNewPassword {
			var message = "Passwords are the same. You cannot change password to old one"
			conflictResponse(w, r, &message)
			return
		}

		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		if user == nil {
			respondNotFound(w)
			return
		}

		if utils.CheckPasswordHash(formOldPassword, user.Password) {
			hash, err := utils.HashPassword(formNewPassword)
			if err != nil {
				logger.Log(err.Error(), logger.Error)
				respondInternalServerError(w)
				return
			}

			user.Password = hash

			err = h.userSqlService.UpdateUserPassword(user)
			if err != nil {
				logger.Log(err.Error(), logger.Error)
				respondInternalServerError(w)
				return
			}

			userResponse(user, w, r)
			return
		}

		respondNotFound(w)
		return
	}

}

func (h handler) postUsersRegistration() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondBadRequest(w, "")
			return
		}

		formUsername := r.PostForm.Get("username")
		formPassword := r.PostForm.Get("password")

		hash, err := utils.HashPassword(formPassword)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		data := make(map[interface{}]interface{})
		data[0] = "ROLE_USER"
		roles, err := phpserialize.Encode(data)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		user := &model.User{Username: formUsername, Password: hash, RolesJson: roles}
		createdUserId, err := h.userSqlService.CreateUser(user)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		user, err = h.userSqlService.GetUserById(createdUserId)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		if user == nil {
			respondNotFound(w)
			return
		}

		userResponse(user, w, r)
	}

}

func (h handler) postUsersExist() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondBadRequest(w, "")
			return
		}

		formUsername := r.PostForm.Get("username")

		user, err := h.userSqlService.GetUserByUsername(formUsername)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		if user == nil {
			respondNotFound(w)
			return
		}

		userResponse(user, w, r)
	}

}

func (h handler) postUserDevices() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondBadRequest(w, "")
			return
		}

		userDevice, err := mapper.MapUserDeviceFormToUserDevice(r, model.UserDevice{})
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondBadRequest(w, "")
			return
		}

		userDevice.CreatedAt = time.Now()

		userDevice.LastLogin = time.Now()

		id, err := h.userSqlService.CreateUserDevice(userDevice)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		userDevice.Id = id

		userDeviceResponse(userDevice, w, r)
	}
}

func (h handler) getUserDevices() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondBadRequest(w, "")
			return
		}

		userDevices, err := h.userSqlService.GetDevicesByUserId(userId)

		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		if userDevices != nil && userDevices.Items != nil {
			userDevicesResponse(userDevices, w, r)
			return
		}

		respondNotFound(w)
	}
}

func (h handler) postUsersLogin() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondBadRequest(w, "")
			return
		}

		formUsername := r.PostForm.Get("username")
		formPassword := r.PostForm.Get("password")

		login, err := h.userSqlService.LoginUser(formUsername)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		if login == nil {
			respondNotFound(w)
			return
		}

		if utils.CheckPasswordHash(formPassword, login.Password) {
			user, err := h.userSqlService.GetUserById(login.Id)
			if err != nil {
				logger.Log(err.Error(), logger.Error)
				respondInternalServerError(w)
				return
			}
			tokenString, _ := jwt.CreateToken(formUsername)

			auth := model.Authentication{
				User: user,
				Token: jwt.Token{
					Username:    user.Username,
					TokenString: tokenString,
				},
			}
			authenticationResponse(auth, w)
			return
		}

		respondNotFound(w)
	}
}
