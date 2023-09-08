package filename_generator

type FilenameGenerator interface {
	GetCurrentWithPrefix() string
	Increment() []rune
}

func New(suffixLength int, prefix string, isNumeric bool) FilenameGenerator {
	if isNumeric {
		return NewNumericFilenameGenerator(suffixLength, prefix)
	} else {
		return NewAlphabetFilenameGenerator(suffixLength, prefix)
	}
}
