package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Buffer struct {
	count int
	text  string
}

func readFileByLine(filename string, lineCount int) {
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

		if buffer.count == lineCount {
			writeFunc(buffer.text, "output"+strconv.Itoa(lastIndex)+".txt")
			buffer.text = ""
			buffer.count = 0
			lastIndex += 1
		}
	}

	if buffer.count > 0 {
		writeFunc(buffer.text, "output"+strconv.Itoa(lastIndex)+".txt")
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
