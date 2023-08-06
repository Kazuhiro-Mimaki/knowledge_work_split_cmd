package validation

import (
	"errors"
)

func ValidateIsPositive(number int) (err error) {
	if number < 0 {
		return errors.New("number must be positive")
	}
	return nil
}

func ValidateCmdArgs(args []string) (err error) {
	if len(args) == 1 || len(args) == 2 {
		return nil
	}
	return errors.New("invalid arguments")
}
