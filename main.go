package main

import (
	"errors"
	"flag"
	"fmt"
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

	switch mode(l, n, b) {
	case "l":
		// split file by line
		if err := validatePositive(l); err != nil {
			panic(err)
		}
		readFileByLine(filename, l)
	case "n":
		// split file by chunk
		if err := validatePositive(n); err != nil {
			panic(err)
		}
		readFileByChunk(filename, n)
	case "b":
		// split file by byte
		if err := validatePositive(n); err != nil {
			panic(err)
		}
		readFileByBytes(filename, b)
	case "noArgs":
		// no args
	default:
		panic("only one option can be used: l or n or b")
	}

	fmt.Println("pass")
}

func mode(l, n, b int) string {
	if l != 0 && n == 0 && b == 0 {
		return "l"
	} else if l == 0 && n != 0 && b == 0 {
		return "n"
	} else if l == 0 && n == 0 && b != 0 {
		return "b"
	} else if l == 0 && n == 0 && b == 0 {
		return "noArgs"
	} else {
		return ""
	}
}

func validatePositive(number int) (err error) {
	if number < 0 {
		return errors.New("number must be positive")
	}
	return nil
}
