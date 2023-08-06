package utils

type IFilenameGenerator interface {
	Increment() []rune
	GetCurrentRunes() []rune
	GetOutputFilePath() string
}

type FilenameGenerator struct {
	CurrentRunes []rune
}

func NewFilenameGenerator(defaultRuneCount int, suffix string, isNumeric bool) IFilenameGenerator {
	if isNumeric {
		return NewNumericFilenameGenerator(defaultRuneCount, suffix)
	} else {
		return NewAlphabetFilenameGenerator(defaultRuneCount, suffix)
	}
}
