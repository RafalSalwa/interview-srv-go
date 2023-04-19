package mapper

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/RafalSalwa/interview/model"
	"github.com/gorilla/mux"
)

func MapUserDeviceFormToUserDevice(r *http.Request, device model.UserDevice) (*model.UserDevice, error) {
	userId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		return nil, err
	}
	osType, err := strconv.ParseInt(r.PostForm.Get("osType"), 10, 64)
	if err != nil {
		return nil, err
	}
	sdkVersion, err := strconv.ParseInt(r.PostForm.Get("sdkVersion"), 10, 64)
	if err != nil {
		return nil, err
	}

	device.UserId = userId
	device.FirebaseToken = strings.TrimSpace(r.PostForm.Get("firebaseToken"))
	device.OsType = osType
	device.SdkVersion = sdkVersion
	device.Model = strings.TrimSpace(r.PostForm.Get("model"))
	device.Brand = strings.TrimSpace(r.PostForm.Get("brand"))

	return &device, nil
}

func MapNotificationFormToNotification(r *http.Request, notification model.Notification) (*model.Notification, error) {
	userId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		return nil, err
	}
	title := strings.TrimSpace(r.PostForm.Get("title"))
	body := strings.TrimSpace(r.PostForm.Get("body"))
	clickAction := strings.TrimSpace(r.PostForm.Get("clickAction"))
	if err != nil {
		return nil, err
	}
	channelId, err := strconv.ParseInt(r.PostForm.Get("channelId"), 10, 64)
	if err != nil {
		return nil, err
	}
	data := strings.TrimSpace(r.PostForm.Get("data"))
	notification.UserId = &userId
	notification.Title = title
	notification.Body = &body
	notification.ClickAction = &clickAction
	notification.ChannelId = &channelId
	notification.Data = &data
	notification.Sent = 0
	return &notification, nil
}
