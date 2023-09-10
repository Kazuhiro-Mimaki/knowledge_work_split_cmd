package filename_generator

import "split_cmd/parser"

type FilenameGenerator interface {
	GetCurrentWithPrefix() string
	Increment() []rune
}

func New(cmdArgs *parser.CmdArgs) FilenameGenerator {
	if cmdArgs.IsNumeric {
		return NewNumericFilenameGenerator(cmdArgs.SuffixLength, cmdArgs.Prefix)
	} else {
		return NewAlphabetFilenameGenerator(cmdArgs.SuffixLength, cmdArgs.Prefix)
	}
}
