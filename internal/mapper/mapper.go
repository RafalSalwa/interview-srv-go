package mapper

import (
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func MapUserDeviceFormToUserDevice(r *http.Request, device models.UserDevice) (*models.UserDevice, error) {
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
