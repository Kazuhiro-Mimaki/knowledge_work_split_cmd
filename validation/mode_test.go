package validation

import (
	"testing"
)

func TestMode(t *testing.T) {
	t.Run("lモード", func(t *testing.T) {
		got := Mode(1, 0, 0)
		want := "l"
		if got != want {
			t.Errorf("Mode(1, 0, 0) == %v, want %s", got, want)
		}
	})

	t.Run("nモード", func(t *testing.T) {
		got := Mode(0, 1, 0)
		want := "n"
		if got != want {
			t.Errorf("Mode(0, 1, 0) == %v, want %s", got, want)
		}
	})

	t.Run("bモード", func(t *testing.T) {
		got := Mode(0, 0, 1)
		want := "b"
		if got != want {
			t.Errorf("Mode(0, 0, 1) == %v, want %s", got, want)
		}
	})

	t.Run("引数なし", func(t *testing.T) {
		got := Mode(0, 0, 0)
		want := "noArgs"
		if got != want {
			t.Errorf("Mode(0, 0, 0) == %v, want %s", got, want)
		}
	})

	t.Run("不正な入力(2つの入力値が存在する)", func(t *testing.T) {
		got := Mode(0, 1, 1)
		want := ""
		if got != want {
			t.Errorf("Mode(0, 1, 1) == %v, want %s", got, want)
		}
	})

	t.Run("不正な入力(全て入力値が存在する)", func(t *testing.T) {
		got := Mode(1, 1, 1)
		want := ""
		if got != want {
			t.Errorf("Mode(1, 1, 1) == %v, want %s", got, want)
		}
	})
}
