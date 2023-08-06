package output_filename

import (
	"reflect"
	"testing"
)

func TestAlphabetOutputFilenameGenerator(t *testing.T) {
	t.Run("正常系 (引数指定がない場合)", func(t *testing.T) {
		filenameGenerator := NewAlphabetFilenameGenerator(0, "")
		want := "aa"
		if !reflect.DeepEqual(filenameGenerator.GetOutputFilePath(), want) {
			t.Errorf("filenameGenerator.GetOutputFilePath() == %v, want %v", filenameGenerator.GetOutputFilePath(), want)
		}
	})

	t.Run("正常系 (引数指定がある場合)", func(t *testing.T) {
		filenameGenerator := NewAlphabetFilenameGenerator(3, "")
		want := "aaa"
		if !reflect.DeepEqual(filenameGenerator.GetOutputFilePath(), want) {
			t.Errorf("filenameGenerator.GetOutputFilePath() == %v, want %v", filenameGenerator.GetOutputFilePath(), want)
		}
	})

	t.Run("正常系 (suffix指定がある場合)", func(t *testing.T) {
		filenameGenerator := NewAlphabetFilenameGenerator(3, "suffix_")
		want := "suffix_aaa"
		if !reflect.DeepEqual(filenameGenerator.GetOutputFilePath(), want) {
			t.Errorf("filenameGenerator.GetOutputFilePath() == %v, want %v", filenameGenerator.GetOutputFilePath(), want)
		}
	})
}

func TestAlphabetIncrement(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		filenameGenerator := AlphabetFilenameGenerator{currentRunes: []rune{'a', 'a', 'a'}, suffix: ""}
		filenameGenerator.Increment()
		want := "aab"
		if !reflect.DeepEqual(filenameGenerator.GetOutputFilePath(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetOutputFilePath(), want)
		}
	})

	t.Run("正常系 (最後尾がzの場合)", func(t *testing.T) {
		filenameGenerator := AlphabetFilenameGenerator{currentRunes: []rune{'a', 'a', 'z'}, suffix: ""}
		filenameGenerator.Increment()
		want := "aba"
		if !reflect.DeepEqual(filenameGenerator.GetOutputFilePath(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetOutputFilePath(), want)
		}
	})

	t.Run("正常系 (全てzの場合)", func(t *testing.T) {
		filenameGenerator := AlphabetFilenameGenerator{currentRunes: []rune{'z', 'z', 'z'}, suffix: ""}
		filenameGenerator.Increment()
		want := "aaaa"
		if !reflect.DeepEqual(filenameGenerator.GetOutputFilePath(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetOutputFilePath(), want)
		}
	})

	t.Run("正常系 (suffix指定がある場合)", func(t *testing.T) {
		filenameGenerator := AlphabetFilenameGenerator{currentRunes: []rune{'z', 'z', 'z'}, suffix: "suffix_"}
		filenameGenerator.Increment()
		want := "suffix_aaaa"
		if !reflect.DeepEqual(filenameGenerator.GetOutputFilePath(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetOutputFilePath(), want)
		}
	})
}
