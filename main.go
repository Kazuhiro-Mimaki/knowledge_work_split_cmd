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
		// split file by line
		if err := split.ByLine(file, cmdArgs.LineCount, filenameGenerator); err != nil {
			log.Fatal(err)
		}
	case parser.Chunk:
		// split file by chunk
		if err := split.ByChunk(file, fileSize, cmdArgs.ChunkCount, filenameGenerator); err != nil {
			log.Fatal(err)
		}
	case parser.Byte:
		// split file by byte
		if err := split.ByByteCount(file, cmdArgs.ByteCount, filenameGenerator); err != nil {
			log.Fatal(err)
		}
	default:
		err := errors.New("only one option can be used: 'l' or 'n' or 'b' or 'a'")
		log.Fatal(err)
	}

	log.Print("success")
	return
}
