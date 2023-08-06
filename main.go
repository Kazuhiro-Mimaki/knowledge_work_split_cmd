package main

import (
	"errors"
	"flag"
	"log"

	"split_cmd/cmd"
	"split_cmd/utils"
)

func main() {
	var (
		l int
		n int
		b int
		a int
		d bool
	)

	flag.IntVar(&l, "l", 0, "line_count")
	flag.IntVar(&n, "n", 0, "chunk_count")
	flag.IntVar(&b, "b", 0, "byte_count")
	flag.IntVar(&a, "a", 0, "suffix_length")
	flag.BoolVar(&d, "d", false, "is_numeric_suffix")

	flag.Parse()

	if err := utils.ValidateCmdArgs(flag.Args()); err != nil {
		log.Fatal(err)
	}

	var (
		filename string
		suffix   string
	)

	if len(flag.Args()) == 1 {
		filename = flag.Arg(0)
	} else if len(flag.Args()) == 2 {
		filename, suffix = flag.Arg(0), flag.Arg(1)
	}

	filenameGenerator := utils.NewFilenameGenerator(a, d)

	switch utils.Mode(l, n, b) {
	case "l":
		// split file by line
		if err := utils.ValidatePositive(l); err != nil {
			log.Fatal(err)
		}
		if err := cmd.ExecuteByLine(filename, suffix, l, filenameGenerator); err != nil {
			log.Fatal(err)
		}
	case "n":
		// split file by chunk
		if err := utils.ValidatePositive(n); err != nil {
			log.Fatal(err)
		}
		if err := cmd.ExecuteByChunk(filename, suffix, n, filenameGenerator); err != nil {
			log.Fatal(err)
		}
	case "b":
		// split file by byte
		if err := utils.ValidatePositive(n); err != nil {
			log.Fatal(err)
		}
		if err := cmd.ExecuteByteCount(filename, suffix, b, filenameGenerator); err != nil {
			log.Fatal(err)
		}
	case "noArgs":
		// 引数がない場合は1つのファイルに書き込む
		if err := cmd.ExecuteByChunk(filename, suffix, 1, filenameGenerator); err != nil {
			log.Fatal(err)
		}
	default:
		err := errors.New("only one option can be used: 'l' or 'n' or 'b' or 'a'")
		log.Fatal(err)
	}

	log.Print("success")
	return
}
