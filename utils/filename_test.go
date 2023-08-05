package utils

import (
	"reflect"
	"testing"
)

func TestInitFileManager(t *testing.T) {
	t.Run("引数指定がない場合", func(t *testing.T) {
		filenameManager := NewFilenameManager(0)
		want := []rune{'a', 'a'}
		if !reflect.DeepEqual(filenameManager.CurrentRunes, want) {
			t.Errorf("filenameManager.CurrentRunes == %v, want %v", filenameManager.CurrentRunes, want)
		}
	})

	t.Run("引数指定がある場合", func(t *testing.T) {
		filenameManager := NewFilenameManager(3)
		want := []rune{'a', 'a', 'a'}
		if !reflect.DeepEqual(filenameManager.CurrentRunes, want) {
			t.Errorf("filenameManager.CurrentRunes == %v, want %v", filenameManager.CurrentRunes, want)
		}
	})
}

func TestIncrement(t *testing.T) {
	t.Run("正常ケース", func(t *testing.T) {
		filenameManager := FilenameManager{[]rune{'a', 'a', 'a'}}
		filenameManager.Increment()
		want := []rune{'a', 'a', 'b'}
		if !reflect.DeepEqual(filenameManager.CurrentRunes, want) {
			t.Errorf("filenameManager.Increment() == %v, want %v", filenameManager.CurrentRunes, want)
		}
	})

	t.Run("正常ケース (最後尾がzの場合)", func(t *testing.T) {
		filenameManager := FilenameManager{[]rune{'a', 'a', 'z'}}
		filenameManager.Increment()
		want := []rune{'a', 'b', 'a'}
		if !reflect.DeepEqual(filenameManager.CurrentRunes, want) {
			t.Errorf("filenameManager.Increment() == %v, want %v", filenameManager.CurrentRunes, want)
		}
	})

	t.Run("正常ケース (全てzの場合)", func(t *testing.T) {
		filenameManager := FilenameManager{[]rune{'z', 'z', 'z'}}
		filenameManager.Increment()
		want := []rune{'a', 'a', 'a', 'a'}
		if !reflect.DeepEqual(filenameManager.CurrentRunes, want) {
			t.Errorf("filenameManager.Increment() == %v, want %v", filenameManager.CurrentRunes, want)
		}
	})
}
