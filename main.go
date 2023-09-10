package main

import (
	"errors"
	"flag"
	"log"

	"split_cmd/cmd"
	"split_cmd/filename_generator"
	"split_cmd/validation"
)

func main() {
	var (
		l            int
		n            int
		b            int
		suffixLength int
		isNumeric    bool
	)

	flag.IntVar(&l, "l", 0, "line_count")
	flag.IntVar(&n, "n", 0, "chunk_count")
	flag.IntVar(&b, "b", 0, "byte_count")
	// デフォルトを2にする
	flag.IntVar(&suffixLength, "a", 2, "suffix_length")
	flag.BoolVar(&isNumeric, "d", false, "is_numeric")

	flag.Parse()

	if err := validation.ValidateCmdArgs(flag.Args()); err != nil {
		log.Fatal(err)
	}

	var (
		readFilePath string
		prefix       string
	)

	if len(flag.Args()) == 1 {
		readFilePath, prefix = flag.Arg(0), ""
	} else if len(flag.Args()) == 2 {
		readFilePath, prefix = flag.Arg(0), flag.Arg(1)
	}

	filenameGenerator := filename_generator.New(suffixLength, prefix, isNumeric)

	switch validation.Mode(l, n, b) {
	case "l":
		// split file by line
		if isPositive := validation.IsPositive(l); isPositive != true {
			log.Fatal("number must be positive")
		}
		if err := cmd.ExecuteByLine(readFilePath, l, filenameGenerator); err != nil {
			log.Fatal(err)
		}
	case "n":
		// split file by chunk
		if isPositive := validation.IsPositive(n); isPositive != true {
			log.Fatal("number must be positive")
		}
		if err := cmd.ExecuteByChunk(readFilePath, n, filenameGenerator); err != nil {
			log.Fatal(err)
		}
	case "b":
		// split file by byte
		if isPositive := validation.IsPositive(b); isPositive != true {
			log.Fatal("number must be positive")
		}
		if err := cmd.ExecuteByteCount(readFilePath, b, filenameGenerator); err != nil {
			log.Fatal(err)
		}
	case "noArgs":
		// 引数がない場合は1つのファイルに書き込む
		if err := cmd.ExecuteByChunk(readFilePath, 1, filenameGenerator); err != nil {
			log.Fatal(err)
		}
	default:
		err := errors.New("only one option can be used: 'l' or 'n' or 'b' or 'a'")
		log.Fatal(err)
	}

	log.Print("success")
	return
}
