package responses

import (
	"net/http"

	grpc_codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func FromGRPCError(err *status.Status, w http.ResponseWriter) {
	switch err.Code() {
	case grpc_codes.AlreadyExists:
		RespondConflict(w, err.Message())
	case grpc_codes.NotFound:
		RespondNotFound(w)
	default:
		RespondBadRequest(w, err.Message())
	}
}
