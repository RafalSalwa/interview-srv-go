package responses

import (
    grpc_codes "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "net/http"
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
