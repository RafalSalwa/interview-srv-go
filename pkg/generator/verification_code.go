package generator

import (
	crand "crypto/rand"
	"errors"
	"math/big"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func RandomString(length int) (string, error) {
	if length < 6 || length > 18 {
		return "", errors.New("code should be between 6 and 10 letters")
	}

	b := make([]byte, length)
	_, err := crand.Read(b)
	if err != nil {
		return "", err
	}

	for i := 0; i < length; i++ {
		num, err := crand.Int(crand.Reader, big.NewInt(int64(len(letterBytes))))
		if err != nil {
			return "", err
		}
		b[i] = letterBytes[num.Int64()]
	}

	return string(b), nil
}
