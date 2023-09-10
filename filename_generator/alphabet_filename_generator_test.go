package filename_generator

import (
	"reflect"
	"testing"
)

func TestAlphabetFilenameGenerator(t *testing.T) {
	cases := map[string]struct {
		suffixLength int
		prefix       string
		want         string
	}{
		"prefixなし": {3, "", "aaa"},
		"prefix付き": {3, "prefix_", "prefix_aaa"},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			filenameGenerator := NewAlphabetFilenameGenerator(tt.suffixLength, tt.prefix)
			if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), tt.want) {
				t.Errorf("filenameGenerator.GetCurrentWithPrefix() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), tt.want)
			}
		})
	}
}

func TestAlphabetFilenameIncrement(t *testing.T) {
	cases := map[string]struct {
		current []rune
		want    string
	}{
		"正常系":      {[]rune{'a', 'a'}, "ab"},
		"最後尾がzの場合": {[]rune{'a', 'z'}, "ba"},
		"全てzの場合":   {[]rune{'z', 'z'}, "aaa"},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			filenameGenerator := AlphabetFilenameGenerator{current: tt.current, prefix: ""}
			filenameGenerator.Increment()
			if !reflect.DeepEqual(filenameGenerator.GetCurrentWithPrefix(), tt.want) {
				t.Errorf("filenameGenerator.Increment() == %v, want %v", filenameGenerator.GetCurrentWithPrefix(), tt.want)
			}
		})
	}
}
