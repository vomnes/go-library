package lib

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789_-"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// GetRandomString create a random string with a length of n characters
// with the characters include in letterBytes
func GetRandomString(n int) string {
	b := make([]byte, n)
	src := rand.NewSource(time.Now().UnixNano())
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func CaptureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

// UniqueTimeToken generates a unique base64 token using a key and time.Now() as string
func UniqueTimeToken(key string) string {
	now := time.Now()
	data := []byte(key + "&" + now.String())
	return base64.StdEncoding.EncodeToString(data)
}

func InterfaceToByte(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func SWAPStrings(str1, str2 *string) {
	*str1, *str2 = *str2, *str1
}
