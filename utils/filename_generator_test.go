package utils

import (
	"testing"
)

func TestNewFilenameGenerator(t *testing.T) {
	t.Run("初期ファイル生成", func(t *testing.T) {
		got := NewFilenameGenerator()
		want := "0"
		if got.CurrentName != want {
			t.Errorf("NewFilenameGenerator() == %v, want %s", got.CurrentName, want)
		}
	})
}

func TestIncrementFilename(t *testing.T) {
	t.Run("ファイル名を次に進める", func(t *testing.T) {
		got := NewFilenameGenerator()
		for i := 0; i < 7; i++ {
			got.Increment()
		}
		want := "7"
		if got.CurrentName != want {
			t.Errorf("NewFilenameGenerator() == %v, want %s", got.CurrentName, want)
		}
	})
}
