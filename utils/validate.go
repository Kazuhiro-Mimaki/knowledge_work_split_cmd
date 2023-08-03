package utils

import "errors"

func ValidatePositive(number int) (err error) {
	if number < 0 {
		return errors.New("number must be positive")
	}
	return nil
}
