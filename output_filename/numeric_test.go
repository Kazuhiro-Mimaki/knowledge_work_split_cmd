package output_filename

import (
	"reflect"
	"testing"
)

func TestNumericFilenameGenerator(t *testing.T) {
	t.Run("引数指定がない場合", func(t *testing.T) {
		filenameGenerator := NewNumericFilenameGenerator(0, "")
		want := []rune{'0', '0'}
		if !reflect.DeepEqual(filenameGenerator.GetCurrentRunes(), want) {
			t.Errorf("filenameGenerator.GetCurrentRunes() == %v, want %v", filenameGenerator.GetCurrentRunes(), want)
		}
	})

	t.Run("引数指定がある場合", func(t *testing.T) {
		filenameGenerator := NewNumericFilenameGenerator(3, "")
		want := []rune{'0', '0', '0'}
		if !reflect.DeepEqual(filenameGenerator.GetCurrentRunes(), want) {
			t.Errorf("filenameGenerator.GetCurrentRunes() == %v, want %v", filenameGenerator.GetCurrentRunes(), want)
		}
	})
}

func TestNumericIncrement(t *testing.T) {
	t.Run("正常ケース", func(t *testing.T) {
		filenameGenerator := NumericFilenameGenerator{currentRunes: []rune{'0', '0', '0'}, suffix: ""}
		filenameGenerator.Increment()
		want := []rune{'0', '0', '1'}
		if !reflect.DeepEqual(filenameGenerator.GetCurrentRunes(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentRunes(), want)
		}
	})

	t.Run("正常ケース (最後尾が9の場合)", func(t *testing.T) {
		filenameGenerator := NumericFilenameGenerator{currentRunes: []rune{'0', '0', '9'}, suffix: ""}
		filenameGenerator.Increment()
		want := []rune{'0', '1', '0'}
		if !reflect.DeepEqual(filenameGenerator.GetCurrentRunes(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentRunes(), want)
		}
	})

	t.Run("正常ケース (全て9の場合)", func(t *testing.T) {
		filenameGenerator := NumericFilenameGenerator{currentRunes: []rune{'9', '9', '9'}, suffix: ""}
		filenameGenerator.Increment()
		want := []rune{'0', '0', '0', '0'}
		if !reflect.DeepEqual(filenameGenerator.GetCurrentRunes(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentRunes(), want)
		}
	})
}
