// Package path implements HTTP responses struct features and functions
package responses

import "net/http"

type ErrResponseBuilder struct {
	Response       ErrResponse
	ResponseWriter http.ResponseWriter
}

type ErrResponse struct {
	Code    int    `json:"code"`
	Reason  string `json:"reason"`
	Headers map[string]interface{}
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Data    Data   `json:"data"`
	Err     error
}

func NewErrorBuilder() *ErrResponseBuilder {
	return &ErrResponseBuilder{
		Response: ErrResponse{},
	}
}

func (rb *ErrResponseBuilder) SetResponseCode(statusCode int) *ErrResponseBuilder {
	rb.Response.Code = statusCode
	return rb
}

func (rb *ErrResponseBuilder) SetReason(reason string) *ErrResponseBuilder {
	rb.Response.Reason = reason
	return rb
}

func (rb *ErrResponseBuilder) AddHeader(key, val string) *ErrResponseBuilder {
	rb.Response.Headers[key] = val
	return rb
}

func (rb *ErrResponseBuilder) SetMessage(msg string) *ErrResponseBuilder {
	rb.Response.Message = msg
	return rb
}

func (rb *ErrResponseBuilder) SetWriter(w http.ResponseWriter) *ErrResponseBuilder {
	rb.ResponseWriter = w
	return rb
}

func (rb *ErrResponseBuilder) SetError(err error) *ErrResponseBuilder {
	rb.Response.Err = err
	return rb
}

func (rb *ErrResponseBuilder) Respond() error {
	body := marshalErrorResponse(rb.Response.Reason)

	rb.ResponseWriter.WriteHeader(rb.Response.Code)
	_, err := rb.ResponseWriter.Write(body)
	if err != nil {
		rb.Response.Err = err
		return err
	}
	return nil
}
