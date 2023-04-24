package handler

import (
	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/RafalSalwa/interview-app-srv/util/password"
	"github.com/gorilla/mux"
	phpserialize "github.com/kovetskiy/go-php-serialize"
	"net/http"
	"strconv"
)

func (h handler) GetUser() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := mux.Vars(r)["id"]

		intId, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			h.logger.Error().Err(err)
			responses.RespondBadRequest(w, "Wrong paramater Type. Required int")
			return
		}

		user, err := h.userSqlService.GetUserById(intId)
		if err != nil {
			h.logger.Error().Err(err)
			responses.RespondInternalServerError(w)
			return
		}

		if user == nil {
			responses.RespondNotFound(w)
			return
		}

		responses.NewUserResponse(user, w, r)
	}

}

func (h handler) PostUser() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			h.logger.Error().Err(err)
			responses.RespondBadRequest(w, "")
			return
		}

		pathUserId := mux.Vars(r)["id"]
		formFirstname := r.PostForm.Get("firstname")
		formLastname := r.PostForm.Get("lastname")

		if len(formFirstname) == 0 || len(formLastname) == 0 {
			responses.RespondBadRequest(w, "")
			return
		}

		userId, err := strconv.ParseInt(pathUserId, 10, 64)
		if err != nil {
			h.logger.Error().Err(err)
			responses.RespondBadRequest(w, "")
			return
		}

		user, err := h.userSqlService.GetUserById(userId)
		if err != nil {
			h.logger.Error().Err(err)
			responses.RespondInternalServerError(w)
			return
		}

		if user == nil {
			responses.RespondNotFound(w)
			return
		}

		user.Firstname = &formFirstname
		user.Lastname = &formLastname

		err = h.userSqlService.UpdateUser(user)
		if err != nil {
			h.logger.Error().Err(err)
			responses.RespondInternalServerError(w)
			return
		}

		responses.NewUserResponse(user, w, r)
	}

}

func (h handler) PasswordChange() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			h.logger.Error().Err(err)
			responses.RespondBadRequest(w, "")
			return
		}

		formUsername := r.PostForm.Get("username")
		formOldPassword := r.PostForm.Get("oldPassword")
		formNewPassword := r.PostForm.Get("newPassword")

		user, err := h.userSqlService.LoginUser(formUsername)

		if formOldPassword == formNewPassword {
			var message = "Passwords are the same. You cannot change password to old one"
			responses.RespondConflict(w, message)
			return
		}

		if err != nil {
			h.logger.Error().Err(err)
			responses.RespondInternalServerError(w)
			return
		}

		if user == nil {
			responses.RespondNotFound(w)
			return
		}

		if password.CheckPasswordHash(formOldPassword, user.Password) {
			hash, err := password.HashPassword(formNewPassword)
			if err != nil {
				h.logger.Error().Err(err)
				responses.RespondInternalServerError(w)
				return
			}

			user.Password = hash

			err = h.userSqlService.UpdateUserPassword(user)
			if err != nil {
				h.logger.Error().Err(err)
				responses.RespondInternalServerError(w)
				return
			}

			responses.NewUserResponse(user, w, r)
			return
		}

		responses.RespondNotFound(w)
		return
	}

}

func (h handler) UserRegistration() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			h.logger.Error().Err(err)
			responses.RespondBadRequest(w, "")
			return
		}

		formUsername := r.PostForm.Get("username")
		formPassword := r.PostForm.Get("password")

		hash, err := password.HashPassword(formPassword)
		if err != nil {
			h.logger.Error().Err(err)
			responses.RespondInternalServerError(w)
			return
		}

		data := make(map[interface{}]interface{})
		data[0] = "ROLE_USER"
		roles, err := phpserialize.Encode(data)
		if err != nil {
			h.logger.Error().Err(err)
			responses.RespondInternalServerError(w)
			return
		}

		user := &models.User{Username: formUsername, Password: hash, RolesJson: roles}
		createdUserId, err := h.userSqlService.CreateUser(user)
		if err != nil {
			h.logger.Error().Err(err)
			responses.RespondInternalServerError(w)
			return
		}

		user, err = h.userSqlService.GetUserById(createdUserId)
		if err != nil {
			h.logger.Error().Err(err)
			responses.RespondInternalServerError(w)
			return
		}

		if user == nil {
			responses.RespondNotFound(w)
			return
		}

		responses.NewUserResponse(user, w, r)
	}

}

func (h handler) UserExist() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			h.logger.Error().Err(err)
			responses.RespondBadRequest(w, "")
			return
		}

		formUsername := r.PostForm.Get("username")

		user, err := h.userSqlService.GetUserByUsername(formUsername)
		if err != nil {
			h.logger.Error().Err(err)
			responses.RespondInternalServerError(w)
			return
		}

		if user == nil {
			responses.RespondNotFound(w)
			return
		}

		responses.NewUserResponse(user, w, r)
	}

}

func (h handler) LogIn() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			h.logger.Error().Err(err)
			responses.RespondBadRequest(w, "")
			return
		}

		formUsername := r.PostForm.Get("username")

		user, err := h.userSqlService.GetUserByUsername(formUsername)
		if err != nil {
			h.logger.Error().Err(err)
			responses.RespondInternalServerError(w)
			return
		}

		if user == nil {
			responses.RespondNotFound(w)
			return
		}

		responses.NewUserResponse(user, w, r)
	}

}
