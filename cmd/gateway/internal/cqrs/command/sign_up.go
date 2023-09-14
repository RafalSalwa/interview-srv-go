package command

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	intrvproto "github.com/RafalSalwa/interview-app-srv/proto/grpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

type SignUpUser struct {
	User models.SignUpUserRequest
}

type SignUpHandler struct {
	authClient intrvproto.AuthServiceClient
}

func NewSignUpHandler(authClient intrvproto.AuthServiceClient) SignUpHandler {
	return SignUpHandler{authClient: authClient}
}

func (h SignUpHandler) Handle(ctx context.Context, cmd SignUpUser) error {
	ctx, span := otel.GetTracerProvider().Tracer("cqrs-command").Start(ctx, "CQRS SignUpUser")
	defer span.End()

	_, err := h.authClient.SignUpUser(ctx, &intrvproto.SignUpUserInput{
		Email:           cmd.User.Email,
		Password:        cmd.User.Password,
		PasswordConfirm: cmd.User.PasswordConfirm,
	})
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	return nil
}
