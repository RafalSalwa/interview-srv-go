package encdec

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

const key string = "RHCaDQMfmrBhuaKLDmBNGOEgERuGzwSi"

func Encrypt(plaintext string) (string, error) {
	c, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", nil
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", nil
	}

	nonce := make([]byte, gcm.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	return string(ciphertext), nil
}
