package main

import (
	"bufio"
	"errors"
	"log"
	"os"
)

func ExecuteByChunk(filename string, chunkCount int) {
	readFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	reader := bufio.NewReader(readFile)

	fileSize, err := getFileSize(readFile)
	if err != nil {
		log.Fatal(err)
	}

	filenameGenerator := NewFilenameGenerator()

	byteCount, rest := fileSize/chunkCount, fileSize%chunkCount

	for i := 1; i < chunkCount; i++ {
		chunks, _, err := readChunksByChunkCount(reader, byteCount)
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
		err = writeChunksByChunkCount(writer, chunks)
		if err != nil {
			log.Fatal(err)
		}

		writeFile.Close()

		filenameGenerator.Increment()
	}

	chunks, _, err := readChunksByChunkCount(reader, byteCount+rest)
	if err != nil {
		log.Fatal(err)
	}

	writeFile, err := os.Create("./tmp_dir/" + filenameGenerator.CurrentName)
	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(writeFile)
	err = writeChunksByChunkCount(writer, chunks)
	if err != nil {
		log.Fatal(err)
	}

	writeFile.Close()
}

func getFileSize(file *os.File) (int, error) {
	var fileSize int
	statFile, err := file.Stat()
	if err != nil {
		return fileSize, err
	}
	size64 := statFile.Size()
	if int64(int(size64)) == size64 {
		fileSize = int(size64)
	} else {
		return fileSize, errors.New("File size is too big")
	}
	return fileSize, nil
}

func readChunksByChunkCount(reader *bufio.Reader, byteCount int) ([]byte, int, error) {
	chunks := make([]byte, byteCount)
	cursor, err := reader.Read(chunks)
	return chunks[:cursor], cursor, err
}

func writeChunksByChunkCount(writer *bufio.Writer, chunks []byte) error {
	_, err := writer.Write(chunks)
	err = writer.Flush()
	return err
}
