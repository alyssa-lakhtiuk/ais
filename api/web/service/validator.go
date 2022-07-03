package service

import (
	"math/rand"
	"regexp"
	"time"
)

const random = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func ValidPhone(phone string) error {
	hasLetters, _ := regexp.MatchString(`[a-zA-Z]`, phone)
	if hasLetters {
		// err "Phone number can't contain letters"
	}
	correct, _ := regexp.MatchString(`[0-9]{10,12}`, phone)
	if !correct {
		// err "Phone number has to have less than 13 digits"
	}
	return nil
}

func IsUnsigned(num int) error {
	if num < 0 {
		// err invalid number
	}
	return nil
}

func GenerateRandomStr(length int) string {
	b := make([]byte, length)
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	for i := range b {
		b[i] = random[y1.Int63()%int64(len(random))]
	}
	return string(b)
}
