package utils

import (
	"reflect"
	"testing"
)

// string(rune(f.lastChar))
func TestFilename(t *testing.T) {
	t.Run("正常ケース", func(t *testing.T) {
		filename := NewFilename([]rune{'a', 'a', 'a'})
		got := filename.Increment()
		want := []rune{'a', 'a', 'b'}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("filename.Increment() == %v, want %v", got, want)
		}
	})

	t.Run("正常ケース (最後尾がzの場合)", func(t *testing.T) {
		filename := NewFilename([]rune{'a', 'a', 'z'})
		got := filename.Increment()
		want := []rune{'a', 'b', 'a'}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("filename.Increment() == %v, want %v", got, want)
		}
	})

	t.Run("正常ケース (全てzの場合)", func(t *testing.T) {
		filename := NewFilename([]rune{'z', 'z', 'z'})
		got := filename.Increment()
		want := []rune{'a', 'a', 'a', 'a'}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("filename.Increment() == %v, want %v", got, want)
		}
	})
}
