package output_filename

import (
	"reflect"
	"testing"
)

func TestNumericOutputFilenameGenerator(t *testing.T) {
	t.Run("引数指定がない場合", func(t *testing.T) {
		filenameGenerator := NewNumericFilenameGenerator(0, "")
		want := "00"
		if !reflect.DeepEqual(filenameGenerator.GetOutputFilePath(), want) {
			t.Errorf("filenameGenerator.GetOutputFilePath() == %v, want %v", filenameGenerator.GetOutputFilePath(), want)
		}
	})

	t.Run("引数指定がある場合", func(t *testing.T) {
		filenameGenerator := NewNumericFilenameGenerator(3, "")
		want := "000"
		if !reflect.DeepEqual(filenameGenerator.GetOutputFilePath(), want) {
			t.Errorf("filenameGenerator.GetOutputFilePath() == %v, want %v", filenameGenerator.GetOutputFilePath(), want)
		}
	})

	t.Run("suffix指定がある場合", func(t *testing.T) {
		filenameGenerator := NewNumericFilenameGenerator(3, "suffix_")
		want := "suffix_000"
		if !reflect.DeepEqual(filenameGenerator.GetOutputFilePath(), want) {
			t.Errorf("filenameGenerator.GetOutputFilePath() == %v, want %v", filenameGenerator.GetOutputFilePath(), want)
		}
	})
}

func TestNumericOutputFilenameIncrement(t *testing.T) {
	t.Run("正常ケース", func(t *testing.T) {
		filenameGenerator := NumericFilenameGenerator{currentRunes: []rune{'0', '0', '0'}, suffix: ""}
		filenameGenerator.Increment()
		want := "001"
		if !reflect.DeepEqual(filenameGenerator.GetOutputFilePath(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetOutputFilePath(), want)
		}
	})

	t.Run("正常ケース (最後尾が9の場合)", func(t *testing.T) {
		filenameGenerator := NumericFilenameGenerator{currentRunes: []rune{'0', '0', '9'}, suffix: ""}
		filenameGenerator.Increment()
		want := "010"
		if !reflect.DeepEqual(filenameGenerator.GetOutputFilePath(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetOutputFilePath(), want)
		}
	})

	t.Run("正常ケース (全て9の場合)", func(t *testing.T) {
		filenameGenerator := NumericFilenameGenerator{currentRunes: []rune{'9', '9', '9'}, suffix: ""}
		filenameGenerator.Increment()
		want := "0000"
		if !reflect.DeepEqual(filenameGenerator.GetOutputFilePath(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetOutputFilePath(), want)
		}
	})

	t.Run("正常ケース (suffix指定がある場合)", func(t *testing.T) {
		filenameGenerator := NumericFilenameGenerator{currentRunes: []rune{'9', '9', '9'}, suffix: "suffix_"}
		filenameGenerator.Increment()
		want := "suffix_0000"
		if !reflect.DeepEqual(filenameGenerator.GetOutputFilePath(), want) {
			t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetOutputFilePath(), want)
		}
	})
}
