package utils

import (
	"errors"
)

func ValidatePositive(number int) (err error) {
	if number < 0 {
		return errors.New("number must be positive")
	}
	return nil
}

func ValidateCmdArgs(args []string) (err error) {
	if 0 < len(args) && len(args) < 3 {
		return nil
	}
	return errors.New("invalid arguments")
}
