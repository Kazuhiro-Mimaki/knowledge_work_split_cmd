package main

import (
	"errors"
	"log"

	"split_cmd/cmd"
	"split_cmd/filename_generator"
	"split_cmd/parser"
)

func main() {
	cmdArgs, err := parser.ParseCmdArgs()
	if err != nil {
		log.Fatal(err)
	}

	filenameGenerator := filename_generator.New(cmdArgs)

	switch cmdArgs.SplitType {
	case parser.Line:
		// split file by line
		if err := cmd.ExecuteByLine(cmdArgs.Filepath, cmdArgs.LineCount, filenameGenerator); err != nil {
			log.Fatal(err)
		}
	case parser.Chunk:
		// split file by chunk
		if err := cmd.ExecuteByChunk(cmdArgs.Filepath, cmdArgs.ChunkCount, filenameGenerator); err != nil {
			log.Fatal(err)
		}
	case parser.Byte:
		// split file by byte
		if err := cmd.ExecuteByteCount(cmdArgs.Filepath, cmdArgs.ByteCount, filenameGenerator); err != nil {
			log.Fatal(err)
		}
	default:
		err := errors.New("only one option can be used: 'l' or 'n' or 'b' or 'a'")
		log.Fatal(err)
	}

	log.Print("success")
	return
}
