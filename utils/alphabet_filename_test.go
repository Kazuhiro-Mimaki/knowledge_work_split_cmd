package utils

import (
	"reflect"
	"testing"
)

func TestAlphabetFilenameGenerator(t *testing.T) {
	t.Run("引数指定がない場合", func(t *testing.T) {
		filenameGenerator := NewAlphabetFilenameGenerator(0, "")
		want := []rune{'a', 'a'}
		if !reflect.DeepEqual(filenameGenerator.GetCurrentRunes(), want) {
			t.Errorf("filenameGenerator.GetCurrentRunes() == %v, want %v", filenameGenerator.GetCurrentRunes(), want)
		}
	})

	t.Run("引数指定がある場合", func(t *testing.T) {
		filenameGenerator := NewAlphabetFilenameGenerator(3, "")
		want := []rune{'a', 'a', 'a'}
		if !reflect.DeepEqual(filenameGenerator.GetCurrentRunes(), want) {
			t.Errorf("filenameGenerator.GetCurrentRunes() == %v, want %v", filenameGenerator.GetCurrentRunes(), want)
		}
	})
}

func TestAlphabetIncrement(t *testing.T) {
	t.Run("正常ケース", func(t *testing.T) {
		filenameGenerator := AlphabetFilenameGenerator{currentRunes: []rune{'a', 'a', 'a'}, suffix: ""}
		filenameGenerator.Increment()
		want := []rune{'a', 'a', 'b'}
		if !reflect.DeepEqual(filenameGenerator.GetCurrentRunes(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentRunes(), want)
		}
	})

	t.Run("正常ケース (最後尾がzの場合)", func(t *testing.T) {
		filenameGenerator := AlphabetFilenameGenerator{currentRunes: []rune{'a', 'a', 'z'}, suffix: ""}
		filenameGenerator.Increment()
		want := []rune{'a', 'b', 'a'}
		if !reflect.DeepEqual(filenameGenerator.GetCurrentRunes(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentRunes(), want)
		}
	})

	t.Run("正常ケース (全てzの場合)", func(t *testing.T) {
		filenameGenerator := AlphabetFilenameGenerator{currentRunes: []rune{'z', 'z', 'z'}, suffix: ""}
		filenameGenerator.Increment()
		want := []rune{'a', 'a', 'a', 'a'}
		if !reflect.DeepEqual(filenameGenerator.GetCurrentRunes(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentRunes(), want)
		}
	})
}
