package filename_generator

import (
	"reflect"
	"testing"
)

func TestNumericFilenameGenerator(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		filenameGenerator := New(3, "", true)
		want := "000"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithPrefix() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (prefix指定がある場合)", func(t *testing.T) {
		filenameGenerator := New(3, "prefix_", true)
		want := "prefix_000"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.GetCurrentWithPrefix() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})
}

func TestNumericFilenameIncrement(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		filenameGenerator := NumericFilenameGenerator{current: []rune{'0', '0'}, prefix: ""}
		filenameGenerator.Increment()
		want := "01"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (最後尾が9の場合)", func(t *testing.T) {
		filenameGenerator := NumericFilenameGenerator{current: []rune{'0', '9'}, prefix: ""}
		filenameGenerator.Increment()
		want := "10"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (全て9の場合)", func(t *testing.T) {
		filenameGenerator := NumericFilenameGenerator{current: []rune{'9', '9'}, prefix: ""}
		filenameGenerator.Increment()
		want := "000"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})

	t.Run("正常系 (prefix指定がある場合)", func(t *testing.T) {
		filenameGenerator := NumericFilenameGenerator{current: []rune{'9', '9'}, prefix: "prefix_"}
		filenameGenerator.Increment()
		want := "prefix_000"
		if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), want)
		}
	})
}
