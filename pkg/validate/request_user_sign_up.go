package validate

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func UserInput(r *http.Request, req interface{}) error {
	reqValidator := validator.New()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return fmt.Errorf("Cannot decode request data")
	}

	if err := reqValidator.Struct(req); err != nil {
		return fmt.Errorf("data validation failed with reason: %s", err.Error())
	}
	return nil
}
