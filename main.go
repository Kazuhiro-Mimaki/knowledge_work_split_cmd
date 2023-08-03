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
	)

	flag.IntVar(&l, "l", 0, "line_count")
	flag.IntVar(&n, "n", 0, "chunk_count")
	flag.IntVar(&b, "b", 0, "byte_count")

	flag.Parse()

	filename := flag.Args()[0]

	var allError error

	switch utils.Mode(l, n, b) {
	case "l":
		// split file by line
		if err := utils.ValidatePositive(l); err != nil {
			allError = err
		}
		allError = cmd.ExecuteByLine(filename, l)
	case "n":
		// split file by chunk
		if err := utils.ValidatePositive(n); err != nil {
			allError = err
		}
		allError = cmd.ExecuteByChunk(filename, n)
	case "b":
		// split file by byte
		if err := utils.ValidatePositive(n); err != nil {
			allError = err
		}
		allError = cmd.ExecuteByteCount(filename, b)
	case "noArgs":
		// 引数がない場合は1つのファイルに書き込む
		allError = cmd.ExecuteByChunk(filename, 1)
	default:
		allError = errors.New("only one option can be used: 'l' or 'n' or 'b'")
	}

	if allError != nil {
		log.Fatal(allError)
	}

	log.Print("success")
	return
}
