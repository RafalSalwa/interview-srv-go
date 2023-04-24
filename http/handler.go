package http

import (
	"net/http"

	"github.com/RafalSalwa/interview/service"
	"github.com/gorilla/mux"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Handler interface {
	getUserById() HandlerFunc
	postUsersIdEdit() HandlerFunc
	postUsersIdChangePassword() HandlerFunc
	postUsersLogin() HandlerFunc
	postUsersRegistration() HandlerFunc
	postUsersExist() HandlerFunc
	getUserDevices() HandlerFunc
	postUserDevices() HandlerFunc
	postNotification() HandlerFunc
}

type handler struct {
	Router         *mux.Router
	userSqlService service.UserSqlService
}

func NewHandler(r *mux.Router, us service.UserSqlService) Handler {
	return handler{r, us}
}
