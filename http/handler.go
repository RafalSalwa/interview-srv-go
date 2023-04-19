package http

import (
	"net/http"

	"github.com/RafalSalwa/interview/service"
	"github.com/gorilla/mux"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Handler interface {
	getUsersId() HandlerFunc
	getUsersIdNotifications() HandlerFunc
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
	sqlService     service.SqlService
	userSqlService service.UserSqlService
}

func NewHandler(r *mux.Router, s service.SqlService, us service.UserSqlService) Handler {
	return handler{r, s, us}
}
