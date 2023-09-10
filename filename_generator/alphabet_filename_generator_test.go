package filename_generator

import (
	"reflect"
	"testing"
)

func TestAlphabetFilenameGenerator(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		filenameGenerator := New(3, "", false)
		want := "aaa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithPrefix() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (prefix指定がある場合)", func(t *testing.T) {
		filenameGenerator := New(3, "prefix_", false)
		want := "prefix_aaa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithPrefix() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})
}

func TestAlphabetFilenameIncrement(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		filenameGenerator := AlphabetFilenameGenerator{current: []rune{'a', 'a'}, prefix: ""}
		filenameGenerator.Increment()
		want := "ab"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (最後尾がzの場合)", func(t *testing.T) {
		filenameGenerator := AlphabetFilenameGenerator{current: []rune{'a', 'z'}, prefix: ""}
		filenameGenerator.Increment()
		want := "ba"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (全てzの場合)", func(t *testing.T) {
		filenameGenerator := AlphabetFilenameGenerator{current: []rune{'z', 'z'}, prefix: ""}
		filenameGenerator.Increment()
		want := "aaa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (prefix指定がある場合)", func(t *testing.T) {
		filenameGenerator := AlphabetFilenameGenerator{current: []rune{'z', 'z'}, prefix: "prefix_"}
		filenameGenerator.Increment()
		want := "prefix_aaa"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})
}
