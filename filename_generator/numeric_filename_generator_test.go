package filename_generator

import (
	"reflect"
	"testing"
)

func TestNumericFilenameGenerator(t *testing.T) {
	cases := map[string]struct {
		suffixLength int
		prefix       string
		want         string
	}{
		"prefixなし": {3, "", "000"},
		"prefix付き": {3, "prefix_", "prefix_000"},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			filenameGenerator := NewNumericFilenameGenerator(tt.suffixLength, tt.prefix)
			if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), tt.want) {
				t.Errorf("filenameGenerator.GetCurrentWithPrefix() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), tt.want)
			}
		})
	}
}

func TestNumericFilenameIncrement(t *testing.T) {
	cases := map[string]struct {
		current []rune
		want    string
	}{
		"正常系":      {[]rune{'0', '0'}, "01"},
		"最後尾が0の場合": {[]rune{'0', '9'}, "10"},
		"全て0の場合":   {[]rune{'9', '9'}, "000"},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			filenameGenerator := NumericFilenameGenerator{current: tt.current, prefix: ""}
			filenameGenerator.Increment()
			if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), tt.want) {
				t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), tt.want)
			}
		})
	}
}
