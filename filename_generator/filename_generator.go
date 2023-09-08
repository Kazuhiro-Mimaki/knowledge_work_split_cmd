package filename_generator

import "errors"

type Mode int

const (
	Alphabet Mode = iota + 1
	Numeric
)

type FilenameGenerator interface {
	GetCurrentWithPrefix() string
	Increment() []rune
}

func New(suffixLength int, prefix string, mode Mode) (FilenameGenerator, error) {
	switch mode {
	case Alphabet:
		return NewAlphabetFilenameGenerator(suffixLength, prefix), nil
	case Numeric:
		return NewNumericFilenameGenerator(suffixLength, prefix), nil
	default:
		return nil, errors.New("invalid mode")
	}
}
