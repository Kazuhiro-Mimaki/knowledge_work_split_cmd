package validation

import (
	"errors"
)

func IsPositive(number int) bool {
	return number >= 0
}

func ValidateCmdArgs(args []string) (err error) {
	// ファイル名とprefixのみ引数で受け付けるため、1 or 2以外はエラーとして扱う
	if len(args) == 1 || len(args) == 2 {
		return nil
	}
	return errors.New("invalid arguments")
}
