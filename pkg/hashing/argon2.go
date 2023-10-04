package hashing

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidHash         = errors.New("argon2id: hash is not in the correct format")
	ErrIncompatibleVariant = errors.New("argon2id: incompatible variant of argon2")
	ErrIncompatibleVersion = errors.New("argon2id: incompatible version of argon2")
	ErrHashesMismatch      = errors.New("argon2id: hashes mismatch")
)

var defaultParams = &params{
	Memory:      64 * 1024,
	Iterations:  4,
	Parallelism: 4,
	SaltLength:  16,
	KeyLength:   32,
}

type params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

func Argon2ID(password string) string {
	salt := generateRandomBytes(defaultParams.SaltLength)

	key := argon2.IDKey([]byte(password), salt, defaultParams.Iterations, defaultParams.Memory, defaultParams.Parallelism, defaultParams.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Key := base64.RawStdEncoding.EncodeToString(key)

	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, defaultParams.Memory, defaultParams.Iterations, defaultParams.Parallelism, b64Salt, b64Key)
}

func Argon2IDComparePasswordAndHash(password, hash string) error {
	match, _, err := checkHash(password, hash)
	if err != nil {
		return err
	}

	if !match {
		return ErrHashesMismatch
	}
	return nil
}

func checkHash(password, hash string) (match bool, params *params, err error) {
	params, salt, key, err := decodeHash(hash)
	if err != nil {
		return false, nil, err
	}

	otherKey := argon2.IDKey([]byte(password), salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)

	keyLen := int32(len(key))
	otherKeyLen := int32(len(otherKey))

	if subtle.ConstantTimeEq(keyLen, otherKeyLen) == 0 {
		return false, params, nil
	}
	if subtle.ConstantTimeCompare(key, otherKey) == 1 {
		return true, params, nil
	}
	return false, params, nil
}

func generateRandomBytes(n uint32) []byte {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return b
}

func decodeHash(hash string) (p *params, salt, key []byte, err error) {
	vals := strings.Split(hash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	if vals[1] != "argon2id" {
		return nil, nil, nil, ErrIncompatibleVariant
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.SaltLength = uint32(len(salt))

	key, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.KeyLength = uint32(len(key))

	return p, salt, key, nil
}
