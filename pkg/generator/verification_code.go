package generator

import (
	crand "crypto/rand"
	"encoding/binary"
	"errors"
	"math/rand"
	"unsafe"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

type cryptSrc struct{}

func (s cryptSrc) Seed(seed int64) {}

func (s cryptSrc) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}
func (s cryptSrc) Uint64() (v uint64) {
	_ = binary.Read(crand.Reader, binary.BigEndian, &v)
	return v
}

func RandomString(l int) (*string, error) {
	if l < 6 || l > 18 {
		return nil, errors.New("code should be between 6 and 10 letters")
	}

	var src cryptSrc
	rnd := rand.New(src)
	b := make([]byte, l)

	for i, cache, remain := l-1, rnd.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rnd.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return (*string)(unsafe.Pointer(&b)), nil
}
