package generator

import (
	"regexp"
	"testing"
)

func TestVerificationCode(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, err := RandomString(6)
			if err != nil {
				t.Errorf("encrypt() error = %v", err)
				return
			}
			letters := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
			if !letters(*code) {
				t.Error("verification code can contain only letters")
				return
			}
			code, err = RandomString(4)
			if err == nil {
				t.Errorf("cannot generate too short codes")
				return
			}
		})
	}
}
