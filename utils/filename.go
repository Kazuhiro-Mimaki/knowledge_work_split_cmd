package utils

type IFilenameGenerator interface {
	Increment() []rune
	GetCurrentRunes() []rune
}

type FilenameGenerator struct {
	CurrentRunes []rune
}

func NewFilenameGenerator(defaultRuneCount int, isNumeric bool) IFilenameGenerator {
	if isNumeric {
		return NewNumericFilenameGenerator(defaultRuneCount)
	} else {
		return NewAlphabetFilenameGenerator(defaultRuneCount)
	}
}
