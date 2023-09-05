package filename_generator

import (
	"reflect"
	"testing"
)

func TestAlphabetFilenameGenerator(t *testing.T) {
	t.Run("正常系 (引数指定がない場合)", func(t *testing.T) {
		filenameGenerator := New(0, "", ALPHABET)
		want := "aa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithSuffix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithSuffix() == %v, want %v", filenameGenerator.GetCurrentWithSuffix(), want)
		}
	})

	t.Run("正常系 (引数指定がある場合)", func(t *testing.T) {
		filenameGenerator := New(3, "", ALPHABET)
		want := "aaa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithSuffix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithSuffix() == %v, want %v", filenameGenerator.GetCurrentWithSuffix(), want)
		}
	})

	t.Run("正常系 (suffix指定がある場合)", func(t *testing.T) {
		filenameGenerator := New(3, "suffix_", ALPHABET)
		want := "suffix_aaa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithSuffix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithSuffix() == %v, want %v", filenameGenerator.GetCurrentWithSuffix(), want)
		}
	})
}

func TestAlphabetFilenameIncrement(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		filenameGenerator := FilenameGenerator{currentRunes: []rune{'a', 'a', 'a'}, suffix: "", mode: ALPHABET}
		filenameGenerator.Increment()
		want := "aab"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithSuffix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithSuffix(), want)
		}
	})

	t.Run("正常系 (最後尾がzの場合)", func(t *testing.T) {
		filenameGenerator := FilenameGenerator{currentRunes: []rune{'a', 'a', 'z'}, suffix: "", mode: ALPHABET}
		filenameGenerator.Increment()
		want := "aba"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithSuffix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithSuffix(), want)
		}
	})

	t.Run("正常系 (全てzの場合)", func(t *testing.T) {
		filenameGenerator := FilenameGenerator{currentRunes: []rune{'z', 'z', 'z'}, suffix: "", mode: ALPHABET}
		filenameGenerator.Increment()
		want := "aaaa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithSuffix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithSuffix(), want)
		}
	})

	t.Run("正常系 (suffix指定がある場合)", func(t *testing.T) {
		filenameGenerator := FilenameGenerator{currentRunes: []rune{'z', 'z', 'z'}, suffix: "suffix_", mode: ALPHABET}
		filenameGenerator.Increment()
		want := "suffix_aaaa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithSuffix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithSuffix(), want)
		}
	})
}

func TestNumericFilenameGenerator(t *testing.T) {
	t.Run("正常系 (引数指定がない場合)", func(t *testing.T) {
		filenameGenerator := New(0, "", NUMERIC)
		want := "00"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithSuffix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithSuffix() == %v, want %v", filenameGenerator.GetCurrentWithSuffix(), want)
		}
	})

	t.Run("正常系 (引数指定がある場合)", func(t *testing.T) {
		filenameGenerator := New(3, "", NUMERIC)
		want := "000"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithSuffix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithSuffix() == %v, want %v", filenameGenerator.GetCurrentWithSuffix(), want)
		}
	})

	t.Run("正常系 (suffix指定がある場合)", func(t *testing.T) {
		filenameGenerator := New(3, "suffix_", NUMERIC)
		want := "suffix_000"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithSuffix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithSuffix() == %v, want %v", filenameGenerator.GetCurrentWithSuffix(), want)
		}
	})
}

func TestNumericFilenameIncrement(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		filenameGenerator := FilenameGenerator{currentRunes: []rune{'0', '0', '0'}, suffix: "", mode: NUMERIC}
		filenameGenerator.Increment()
		want := "001"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithSuffix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithSuffix(), want)
		}
	})

	t.Run("正常系 (最後尾が9の場合)", func(t *testing.T) {
		filenameGenerator := FilenameGenerator{currentRunes: []rune{'0', '0', '9'}, suffix: "", mode: NUMERIC}
		filenameGenerator.Increment()
		want := "010"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithSuffix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithSuffix(), want)
		}
	})

	t.Run("正常系 (全て9の場合)", func(t *testing.T) {
		filenameGenerator := FilenameGenerator{currentRunes: []rune{'9', '9', '9'}, suffix: "", mode: NUMERIC}
		filenameGenerator.Increment()
		want := "0000"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithSuffix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithSuffix(), want)
		}
	})

	t.Run("正常系 (suffix指定がある場合)", func(t *testing.T) {
		filenameGenerator := FilenameGenerator{currentRunes: []rune{'9', '9', '9'}, suffix: "suffix_", mode: NUMERIC}
		filenameGenerator.Increment()
		want := "suffix_0000"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithSuffix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithSuffix(), want)
		}
	})
}
