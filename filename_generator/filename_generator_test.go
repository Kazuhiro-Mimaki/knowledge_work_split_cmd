package filename_generator

import (
	"reflect"
	"testing"
)

func TestAlphabetFilenameGenerator(t *testing.T) {
	t.Run("正常系 (引数指定がない場合)", func(t *testing.T) {
		filenameGenerator, _ := New(0, "", Alphabet)
		want := "aa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithPrefix() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (引数指定がある場合)", func(t *testing.T) {
		filenameGenerator, _ := New(3, "", Alphabet)
		want := "aaa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithPrefix() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (prefix指定がある場合)", func(t *testing.T) {
		filenameGenerator, _ := New(3, "prefix_", Alphabet)
		want := "prefix_aaa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithPrefix() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})
}

func TestAlphabetFilenameIncrement(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		filenameGenerator := AlphabetFilenameGenerator{current: []rune{'a', 'a', 'a'}, prefix: ""}
		filenameGenerator.Increment()
		want := "aab"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (最後尾がzの場合)", func(t *testing.T) {
		filenameGenerator := AlphabetFilenameGenerator{current: []rune{'a', 'a', 'z'}, prefix: ""}
		filenameGenerator.Increment()
		want := "aba"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (全てzの場合)", func(t *testing.T) {
		filenameGenerator := AlphabetFilenameGenerator{current: []rune{'z', 'z', 'z'}, prefix: ""}
		filenameGenerator.Increment()
		want := "aaaa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (prefix指定がある場合)", func(t *testing.T) {
		filenameGenerator := AlphabetFilenameGenerator{current: []rune{'z', 'z', 'z'}, prefix: "prefix_"}
		filenameGenerator.Increment()
		want := "prefix_aaaa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})
}

func TestNumericFilenameGenerator(t *testing.T) {
	t.Run("正常系 (引数指定がない場合)", func(t *testing.T) {
		filenameGenerator, _ := New(0, "", Numeric)
		want := "00"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithPrefix() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (引数指定がある場合)", func(t *testing.T) {
		filenameGenerator, _ := New(3, "", Numeric)
		want := "000"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithPrefix() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (prefix指定がある場合)", func(t *testing.T) {
		filenameGenerator, _ := New(3, "prefix_", Numeric)
		want := "prefix_000"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithPrefix() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})
}

func TestNumericFilenameIncrement(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		filenameGenerator := NumericFilenameGenerator{current: []rune{'0', '0', '0'}, prefix: ""}
		filenameGenerator.Increment()
		want := "001"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (最後尾が9の場合)", func(t *testing.T) {
		filenameGenerator := NumericFilenameGenerator{current: []rune{'0', '0', '9'}, prefix: ""}
		filenameGenerator.Increment()
		want := "010"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (全て9の場合)", func(t *testing.T) {
		filenameGenerator := NumericFilenameGenerator{current: []rune{'9', '9', '9'}, prefix: ""}
		filenameGenerator.Increment()
		want := "0000"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (prefix指定がある場合)", func(t *testing.T) {
		filenameGenerator := NumericFilenameGenerator{current: []rune{'9', '9', '9'}, prefix: "prefix_"}
		filenameGenerator.Increment()
		want := "prefix_0000"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})
}
