package http

import (
	"github.com/RafalSalwa/interview/mapper"
	"github.com/RafalSalwa/interview/model"
	"github.com/RafalSalwa/interview/utils/logger"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func (h handler) postNotification() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondBadRequest(w, "")
			return
		}
		userId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		notification, err := mapper.MapNotificationFormToNotification(r, model.Notification{})
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondBadRequest(w, "")
			return
		}

		notification.CreatedAt = time.Now()
		userDevice, err := h.userSqlService.GetLatestDevice(userId)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		notification.DeviceId = userDevice.Id
		id, err := h.sqlService.CreateNotification(notification)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		notification.Id = id

		userNotificationResponse(notification, w, r)
	}
}
