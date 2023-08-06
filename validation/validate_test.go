package validation

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	t.Run("正の数", func(t *testing.T) {
		got := ValidateIsPositive(13)
		if got != nil {
			t.Errorf("ValidateIsPositive(13) == %v, want nil", got)
		}
	})

	t.Run("0", func(t *testing.T) {
		got := ValidateIsPositive(0)
		if got != nil {
			t.Errorf("ValidateIsPositive(0) == %v, want nil", got)
		}
	})

	t.Run("負の数", func(t *testing.T) {
		got := ValidateIsPositive(-13)
		want := errors.New("number must be positive")
		if got.Error() != want.Error() {
			t.Errorf("ValidateIsPositive(13) == %v, want %s", got, want)
		}
	})
}
