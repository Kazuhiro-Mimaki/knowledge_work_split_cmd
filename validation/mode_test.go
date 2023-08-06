package validation

import (
	"testing"
)

func TestMode(t *testing.T) {
	t.Run("正常系 (lモード)", func(t *testing.T) {
		got := Mode(1, 0, 0)
		want := "l"
		if got != want {
			t.Errorf("Mode(1, 0, 0) == %v, want %s", got, want)
		}
	})

	t.Run("正常系 (nモード)", func(t *testing.T) {
		got := Mode(0, 1, 0)
		want := "n"
		if got != want {
			t.Errorf("Mode(0, 1, 0) == %v, want %s", got, want)
		}
	})

	t.Run("正常系 (bモード)", func(t *testing.T) {
		got := Mode(0, 0, 1)
		want := "b"
		if got != want {
			t.Errorf("Mode(0, 0, 1) == %v, want %s", got, want)
		}
	})

	t.Run("正常系 (引数なし)", func(t *testing.T) {
		got := Mode(0, 0, 0)
		want := "noArgs"
		if got != want {
			t.Errorf("Mode(0, 0, 0) == %v, want %s", got, want)
		}
	})

	t.Run("異常系 (不正な入力(2つの入力値が存在する))", func(t *testing.T) {
		got := Mode(0, 1, 1)
		want := ""
		if got != want {
			t.Errorf("Mode(0, 1, 1) == %v, want %s", got, want)
		}
	})

	t.Run("異常系 (全て入力値が存在する)", func(t *testing.T) {
		got := Mode(1, 1, 1)
		want := ""
		if got != want {
			t.Errorf("Mode(1, 1, 1) == %v, want %s", got, want)
		}
	})
}
