package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
    var (
        l int
    )

    flag.IntVar(&l, "l", 0, "line_count")
    flag.Parse()
    
    fileArgs := flag.Args()
    filename := fileArgs[0]

    err := validateL(l)
    if err != nil {
        panic(err)
    }

    readFile(filename, l)

    fmt.Println("pass")
}

func validateL(l int) (err error) {
    if l < 0 {
        return errors.New("validation error of l")
    }
    return nil
}

type Buffer struct {
    count int
    text string
}

func readFile(filename string, l int) {
    lastIndex := 1
    buffer := Buffer{0, ""}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
    defer file.Close()

	fileScanner := bufio.NewScanner(file)

	// read line by line
	for fileScanner.Scan() {
        buffer.text += fileScanner.Text() + "\n"
        buffer.count += 1

        if (buffer.count == l) {
            writeFunc(buffer.text, "output" + strconv.Itoa(lastIndex) + ".txt")
            buffer.text = ""
            buffer.count = 0
            lastIndex += 1
        }
	}

    if (buffer.count > 0) {
        writeFunc(buffer.text, "output" + strconv.Itoa(lastIndex) + ".txt")
    }

	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
}

func writeFunc(text, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		log.Fatal(err)
	}
}