package validation

import (
	"testing"
)

func TestValidate(t *testing.T) {
	t.Run("正の数", func(t *testing.T) {
		got := IsPositive(13)
		if got != true {
			t.Errorf("IsPositive(13) == %v, want true", got)
		}
	})

	t.Run("0", func(t *testing.T) {
		got := IsPositive(0)
		if got != true {
			t.Errorf("IsPositive(0) == %v, want true", got)
		}
	})

	t.Run("負の数", func(t *testing.T) {
		got := IsPositive(-13)
		if got != false {
			t.Errorf("IsPositive(13) == %v, want false", got)
		}
	})
}
