package main

import (
	"errors"
	"log"
	"os"

	"split_cmd/filename_generator"
	"split_cmd/parser"
	"split_cmd/split"
)

func main() {
	var err error

	cmdArgs, err := parser.ParseCmdArgs()
	if err != nil {
		log.Fatal(err)
	}

	filenameGenerator := filename_generator.New(cmdArgs)

	file, err := os.Open(cmdArgs.Filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	statFile, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fileSize := int(statFile.Size())

	switch cmdArgs.SplitType {
	case parser.Line:
		err = split.ByLine(file, cmdArgs.LineCount, filenameGenerator)
	case parser.Chunk:
		err = split.ByChunk(file, fileSize, cmdArgs.ChunkCount, filenameGenerator)
	case parser.Byte:
		err = split.ByByteCount(file, cmdArgs.ByteCount, filenameGenerator)
	default:
		err = errors.New("only one option can be used: 'l' or 'n' or 'b' or 'a'")
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Print("success")
	return
}
