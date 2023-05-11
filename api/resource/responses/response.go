package responses

import (
	"encoding/json"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"net/http"
)

type Data struct {
	Success bool    `json:"success"`
	Message *string `json:"message"`
}

type NotificationResponse struct {
	Data *models.Notification `json:"data"`
}
type NotificationsResponse struct {
	Data *models.Notifications `json:"data"`
}
type UserDeviceResponse struct {
	Data *models.UserDevice `json:"data"`
}
type UserDevicesResponse struct {
	Data *models.UserDevices `json:"data"`
}
type AuthResponse struct {
	Data *models.Authentication `json:"data"`
}

type SuccessResponse struct {
	Data Data `json:"data"`
}

type ConflictResponse struct {
	Data Data `json:"conflict"`
}

type Response struct {
	Data interface{} `json:"data"`
}

type UserResponse struct {
	Data *models.UserResponse `json:"user"`
}

func successResponse(w http.ResponseWriter, r *http.Request) {
	response := &SuccessResponse{Data: Data{
		Success: true,
	}}
	js, err := json.Marshal(response)
	if err != nil {
		//logger.Log(err.Error(), logger.Error)
		RespondInternalServerError(w)
	}

	Respond(w, http.StatusOK, js)
}

func errorResponse(w http.ResponseWriter, r *http.Request, message *string) {
	response := &ErrorResponse{Data: Data{
		Success: false,
		Message: message,
	}}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		//logger.Log(err.Error(), logger.Error)
		RespondInternalServerError(w)
	}

	Respond(w, http.StatusOK, jsonResponse)
}
func conflictResponse(w http.ResponseWriter, r *http.Request, message *string) {
	response := &ErrorResponse{Data: Data{
		Success: false,
		Message: message,
	}}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		//logger.Log(err.Error(), logger.Error)
		RespondInternalServerError(w)
	}

	Respond(w, http.StatusConflict, jsonResponse)
}

func AuthenticationResponse(a models.Authentication, w http.ResponseWriter) {
	response := &Response{Data: a}
	js, err := json.Marshal(response)
	if err != nil {
		//logger.Log(err.Error(), logger.Error)
		RespondInternalServerError(w)
	}

	Respond(w, http.StatusOK, js)
}

func RespondInternalServerError(w http.ResponseWriter) {
	errorResponse := NewInternalServerErrorErrorResponse()
	responseBody := marshalErrorResponse(errorResponse)
	Respond(w, http.StatusInternalServerError, responseBody)
}

func RespondNotFound(w http.ResponseWriter) {
	response := NewNotFoundResponse()
	responseBody := marshalErrorResponse(response)
	Respond(w, http.StatusNotFound, responseBody)
}

func RespondNotAuthorized(w http.ResponseWriter, msg string) {
	errorResponse := NewUnauthorizedErrorResponse(msg)
	responseBody := marshalErrorResponse(errorResponse)
	Respond(w, http.StatusUnauthorized, responseBody)
}

func RespondConflict(w http.ResponseWriter, msg string) {
	resp := NewConflictResponse(msg)
	responseBody := marshalErrorResponse(resp)
	Respond(w, http.StatusConflict, responseBody)
}

func RespondBadRequest(w http.ResponseWriter, msg string) {
	errorResponse := NewBadRequestErrorResponse(msg)
	responseBody := marshalErrorResponse(errorResponse)
	Respond(w, http.StatusBadRequest, responseBody)
}

func Respond(w http.ResponseWriter, statusCode int, responseBody []byte) {
	setHttpHeaders(w, statusCode)
	_, err := w.Write(responseBody)

	if err != nil {
		//logger.LogErr(err)
	}
}

func RespondOk(w http.ResponseWriter) {
	setHttpHeaders(w, http.StatusOK)
	_, err := w.Write([]byte("{\"status\":\"ok\"}"))

	if err != nil {
		RespondInternalServerError(w)
	}
}

func NewUserResponse(u *models.UserResponse, w http.ResponseWriter, r *http.Request) {
	response := &UserResponse{Data: u}
	js, err := json.Marshal(response)
	if err != nil {
		RespondInternalServerError(w)
	}

	Respond(w, http.StatusOK, js)
}

func userNotificationResponse(n *models.Notification, w http.ResponseWriter, r *http.Request) {
	response := &NotificationResponse{Data: n}
	js, err := json.Marshal(response)
	if err != nil {
		//logger.Log(err.Error(), logger.Error)
		RespondInternalServerError(w)
	}

	Respond(w, http.StatusOK, js)
}
func notificationsResponse(ns *models.Notifications, w http.ResponseWriter, r *http.Request) {
	response := &NotificationsResponse{Data: ns}
	js, err := json.Marshal(response)
	if err != nil {
		//logger.Log(err.Error(), logger.Error)
		RespondInternalServerError(w)
	}

	Respond(w, http.StatusOK, js)
}

func userDevicesResponse(ud *models.UserDevices, w http.ResponseWriter, r *http.Request) {
	response := &UserDevicesResponse{Data: ud}
	js, err := json.Marshal(response)
	if err != nil {
		//logger.Log(err.Error(), logger.Error)
		RespondInternalServerError(w)
	}

	Respond(w, http.StatusOK, js)
}

func setHttpHeaders(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
}

func marshalErrorResponse(error interface{}) []byte {
	body, err := json.Marshal(error)

	if err != nil {
		//logger.LogErr(err)
		return nil
	}

	return body
}
