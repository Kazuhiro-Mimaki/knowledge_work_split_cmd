package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFileByChunk(filename string, chunkCount int) {
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a buffered reader
	reader := bufio.NewReader(file)

	var fileSize int
	if statFile, err := file.Stat(); err == nil {
		size64 := statFile.Size()
		if int64(int(size64)) == size64 {
			fileSize = int(size64)
		}
	}

	lastIndex := 1

	for {
		chunk := make([]byte, fileSize/chunkCount)
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
		if err == nil && n < chunkCount {
			break
		}
	}
}

func writeChunk(chunks []byte, filename string) {
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
