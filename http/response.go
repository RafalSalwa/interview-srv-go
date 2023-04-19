package http

import (
	"encoding/json"
	"net/http"

	"github.com/RafalSalwa/interview/model"
	"github.com/RafalSalwa/interview/utils"
	"github.com/RafalSalwa/interview/utils/logger"
)

type Data struct {
	Success bool    `json:"success"`
	Message *string `json:"message"`
}

type NotificationResponse struct {
	Data *model.Notification `json:"data"`
}
type NotificationsResponse struct {
	Data *model.Notifications `json:"data"`
}
type UserDeviceResponse struct {
	Data *model.UserDevice `json:"data"`
}
type UserDevicesResponse struct {
	Data *model.UserDevices `json:"data"`
}
type AuthResponse struct {
	Data *model.Authentication `json:"data"`
}

type SuccessResponse struct {
	Data Data `json:"data"`
}

type ErrorResponse struct {
	Data Data `json:"error"`
}

type ConflictResponse struct {
	Data Data `json:"conflict"`
}

type Response struct {
	Data interface{} `json:"data"`
}

type UserResponse struct {
	Data *model.User `json:"data"`
}

func successResponse(w http.ResponseWriter, r *http.Request) {
	response := &SuccessResponse{Data: Data{
		Success: true,
	}}
	js, err := json.Marshal(response)
	if err != nil {
		logger.Log(err.Error(), logger.Error)
		respondInternalServerError(w)
	}

	respond(w, http.StatusOK, js)
}

func errorResponse(w http.ResponseWriter, r *http.Request, message *string) {
	response := &ErrorResponse{Data: Data{
		Success: false,
		Message: message,
	}}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Log(err.Error(), logger.Error)
		respondInternalServerError(w)
	}

	respond(w, http.StatusOK, jsonResponse)
}
func conflictResponse(w http.ResponseWriter, r *http.Request, message *string) {
	response := &ErrorResponse{Data: Data{
		Success: false,
		Message: message,
	}}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Log(err.Error(), logger.Error)
		respondInternalServerError(w)
	}

	respond(w, http.StatusConflict, jsonResponse)
}

func authenticationResponse(a model.Authentication, w http.ResponseWriter) {
	response := &Response{Data: a}
	js, err := json.Marshal(response)
	if err != nil {
		logger.Log(err.Error(), logger.Error)
		respondInternalServerError(w)
	}

	respond(w, http.StatusOK, js)
}

func respondInternalServerError(w http.ResponseWriter) {
	errorResponse := utils.NewInternalServerErrorErrorResponse()
	responseBody := marshalErrorResponse(errorResponse)
	respond(w, http.StatusInternalServerError, responseBody)
}

func respondNotFound(w http.ResponseWriter) {
	response := utils.NewNotFoundResponse()
	responseBody := marshalErrorResponse(response)
	respond(w, http.StatusNotFound, responseBody)
}

func respondNotAuthorized(w http.ResponseWriter, msg string) {
	errorResponse := utils.NewUnauthorizedErrorResponse(msg)
	responseBody := marshalErrorResponse(errorResponse)
	respond(w, http.StatusUnauthorized, responseBody)
}

func respondBadRequest(w http.ResponseWriter, msg string) {
	errorResponse := utils.NewBadRequestErrorResponse(msg)
	responseBody := marshalErrorResponse(errorResponse)
	respond(w, http.StatusBadRequest, responseBody)
}

func respond(w http.ResponseWriter, statusCode int, responseBody []byte) {
	setHttpHeaders(w, statusCode)
	_, err := w.Write(responseBody)

	if err != nil {
		logger.LogErr(err)
	}
}

func userResponse(u *model.User, w http.ResponseWriter, r *http.Request) {
	response := &UserResponse{Data: u}
	js, err := json.Marshal(response)
	if err != nil {
		logger.Log(err.Error(), logger.Error)
		respondInternalServerError(w)
	}

	respond(w, http.StatusOK, js)
}

func userNotificationResponse(n *model.Notification, w http.ResponseWriter, r *http.Request) {
	response := &NotificationResponse{Data: n}
	js, err := json.Marshal(response)
	if err != nil {
		logger.Log(err.Error(), logger.Error)
		respondInternalServerError(w)
	}

	respond(w, http.StatusOK, js)
}
func notificationsResponse(ns *model.Notifications, w http.ResponseWriter, r *http.Request) {
	response := &NotificationsResponse{Data: ns}
	js, err := json.Marshal(response)
	if err != nil {
		logger.Log(err.Error(), logger.Error)
		respondInternalServerError(w)
	}

	respond(w, http.StatusOK, js)
}

func userDeviceResponse(ud *model.UserDevice, w http.ResponseWriter, r *http.Request) {
	response := &UserDeviceResponse{Data: ud}
	js, err := json.Marshal(response)
	if err != nil {
		logger.Log(err.Error(), logger.Error)
		respondInternalServerError(w)
	}

	respond(w, http.StatusOK, js)
}

func userDevicesResponse(ud *model.UserDevices, w http.ResponseWriter, r *http.Request) {
	response := &UserDevicesResponse{Data: ud}
	js, err := json.Marshal(response)
	if err != nil {
		logger.Log(err.Error(), logger.Error)
		respondInternalServerError(w)
	}

	respond(w, http.StatusOK, js)
}

func setHttpHeaders(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
}

func marshalErrorResponse(error interface{}) []byte {
	body, err := json.Marshal(error)

	if err != nil {
		logger.LogErr(err)
		return nil
	}

	return body
}
