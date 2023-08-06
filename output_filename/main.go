package output_filename

type FilenameGenerator interface {
	Increment() []rune
	GetCurrentRunes() []rune
	GetOutputFilePath() string
}

func NewFilenameGenerator(defaultRuneCount int, suffix string, isNumeric bool) FilenameGenerator {
	if isNumeric {
		return NewNumericFilenameGenerator(defaultRuneCount, suffix)
	} else {
		return NewAlphabetFilenameGenerator(defaultRuneCount, suffix)
	}
}
