package utils

import (
	"log"
	"strconv"
)

type FilenameGenerator struct {
	CurrentName string
}

func NewFilenameGenerator() FilenameGenerator {
	const DEFAULT_NAME = "0"
	return FilenameGenerator{CurrentName: DEFAULT_NAME}
}

func (f *FilenameGenerator) Increment() string {
	i, err := strconv.Atoi(f.CurrentName)
	if err != nil {
		log.Fatal(err)
	}
	f.CurrentName = strconv.Itoa(i + 1)

	return f.CurrentName
}
