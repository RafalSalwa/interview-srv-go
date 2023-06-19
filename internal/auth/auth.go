package auth

import (
	"net/http"
    "handlers"
)

func Authorization(h handlers.HandlerFunc) http.HandlerFunc {
	at, _ := NewAuthMethod(h, "basic")
	return at.middleware(h)
}
