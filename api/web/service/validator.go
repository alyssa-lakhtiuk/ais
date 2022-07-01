package service

import "regexp"

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
