package filename_generator

import (
	"reflect"
	"testing"
)

func TestAlphabetFilenameGenerator(t *testing.T) {
	t.Run("正常系 (引数指定がない場合)", func(t *testing.T) {
		filenameGenerator, _ := New(0, "", ALPHABET)
		want := "aa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithPrefix() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (引数指定がある場合)", func(t *testing.T) {
		filenameGenerator, _ := New(3, "", ALPHABET)
		want := "aaa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithPrefix() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (prefix指定がある場合)", func(t *testing.T) {
		filenameGenerator, _ := New(3, "prefix_", ALPHABET)
		want := "prefix_aaa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithPrefix() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})
}

func TestAlphabetFilenameIncrement(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		filenameGenerator := FilenameGenerator{currentRunes: []rune{'a', 'a', 'a'}, prefix: "", mode: ALPHABET}
		filenameGenerator.Increment()
		want := "aab"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (最後尾がzの場合)", func(t *testing.T) {
		filenameGenerator := FilenameGenerator{currentRunes: []rune{'a', 'a', 'z'}, prefix: "", mode: ALPHABET}
		filenameGenerator.Increment()
		want := "aba"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (全てzの場合)", func(t *testing.T) {
		filenameGenerator := FilenameGenerator{currentRunes: []rune{'z', 'z', 'z'}, prefix: "", mode: ALPHABET}
		filenameGenerator.Increment()
		want := "aaaa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (prefix指定がある場合)", func(t *testing.T) {
		filenameGenerator := FilenameGenerator{currentRunes: []rune{'z', 'z', 'z'}, prefix: "prefix_", mode: ALPHABET}
		filenameGenerator.Increment()
		want := "prefix_aaaa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})
}

func TestNumericFilenameGenerator(t *testing.T) {
	t.Run("正常系 (引数指定がない場合)", func(t *testing.T) {
		filenameGenerator, _ := New(0, "", NUMERIC)
		want := "00"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithPrefix() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (引数指定がある場合)", func(t *testing.T) {
		filenameGenerator, _ := New(3, "", NUMERIC)
		want := "000"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithPrefix() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (prefix指定がある場合)", func(t *testing.T) {
		filenameGenerator, _ := New(3, "prefix_", NUMERIC)
		want := "prefix_000"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithPrefix() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})
}

func TestNumericFilenameIncrement(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		filenameGenerator := FilenameGenerator{currentRunes: []rune{'0', '0', '0'}, prefix: "", mode: NUMERIC}
		filenameGenerator.Increment()
		want := "001"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (最後尾が9の場合)", func(t *testing.T) {
		filenameGenerator := FilenameGenerator{currentRunes: []rune{'0', '0', '9'}, prefix: "", mode: NUMERIC}
		filenameGenerator.Increment()
		want := "010"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (全て9の場合)", func(t *testing.T) {
		filenameGenerator := FilenameGenerator{currentRunes: []rune{'9', '9', '9'}, prefix: "", mode: NUMERIC}
		filenameGenerator.Increment()
		want := "0000"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (prefix指定がある場合)", func(t *testing.T) {
		filenameGenerator := FilenameGenerator{currentRunes: []rune{'9', '9', '9'}, prefix: "prefix_", mode: NUMERIC}
		filenameGenerator.Increment()
		want := "prefix_0000"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})
}
