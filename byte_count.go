package main

import (
	"bufio"
	"log"
	"os"
)

func ExecuteByteCount(filename string, byteCount int) {
	readFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	reader := bufio.NewReader(readFile)

	filenameGenerator := NewFilenameGenerator()

	for {
		chunks, cursor, err := readChunksByByteCount(reader, byteCount)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal(err)
		}

		writeFile, err := os.Create("./tmp_dir/" + filenameGenerator.CurrentName)
		if err != nil {
			log.Fatal(err)
		}

		writer := bufio.NewWriter(writeFile)
		err = writeChunks(writer, chunks)
		if err != nil {
			log.Fatal(err)
		}

		filenameGenerator.Increment()

		writeFile.Close()

		// Check for EOF
		if err == nil && cursor < byteCount {
			break
		}
	}
}

func readChunksByByteCount(reader *bufio.Reader, byteCount int) ([]byte, int, error) {
	chunks := make([]byte, byteCount)
	cursor, err := reader.Read(chunks)
	return chunks[:cursor], cursor, err
}

func writeChunks(writer *bufio.Writer, chunks []byte) error {
	_, err := writer.Write(chunks)
	err = writer.Flush()
	return err
}
