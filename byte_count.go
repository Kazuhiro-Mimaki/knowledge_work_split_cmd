package main

import (
	"bufio"
	"log"
	"os"
)

func readFileByBytes(filename string, byteCount int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	filenameGenerator := NewFilenameGenerator()

	for {
		chunk := make([]byte, byteCount)
		n, err := reader.Read(chunk)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal(err)
		}

		writeBytes(chunk[:n], "./tmp_dir/"+filenameGenerator.CurrentName)

		filenameGenerator.Increment()

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

	writer := bufio.NewWriter(file)

	_, err = writer.Write(chunks)
	if err != nil {
		log.Fatal(err)
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
