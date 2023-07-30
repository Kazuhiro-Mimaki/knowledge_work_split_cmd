package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFileByBytes(filename string, byteCount int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a buffered reader
	reader := bufio.NewReader(file)

	lastIndex := 1

	for {
		chunk := make([]byte, byteCount)
		n, err := reader.Read(chunk)

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal(err)
		}

		// Print the chunk
		fmt.Printf("%s", chunk[:n])

		writeChunk(chunk[:n], "output"+strconv.Itoa(lastIndex)+".txt")

		lastIndex++

		// Check for EOF
		if err == nil && n < byteCount {
			break
		}
	}
}

func writeBytes(chunks []byte, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(chunks)
	if err != nil {
		log.Fatal(err)
	}
}
