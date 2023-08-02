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

	fileSize, err := getFileSize(readFile)
	if err != nil {
		log.Fatal(err)
	}
	if fileSize < chunkCount {
		log.Fatalf("can't split into more than %d files", fileSize)
	}
	byteCount, rest := fileSize/chunkCount, fileSize%chunkCount

	reader := bufio.NewReader(readFile)
	filenameGenerator := NewFilenameGenerator()

	for i := 1; i < chunkCount; i++ {
		chunks, _, err := readChunksByChunkCount(reader, byteCount)
		if err != nil {
			log.Fatal(err)
		}

		err = createFileAndWrite("./tmp_dir/"+filenameGenerator.CurrentName, chunks)
		if err != nil {
			log.Fatal(err)
		}

		filenameGenerator.Increment()
	}

	chunks, _, err := readChunksByChunkCount(reader, byteCount+rest)
	if err != nil {
		log.Fatal(err)
	}

	err = createFileAndWrite("./tmp_dir/"+filenameGenerator.CurrentName, chunks)
	if err != nil {
		log.Fatal(err)
	}
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

func createFileAndWrite(filename string, bytes []byte) error {
	writeFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(writeFile)
	err = writeChunksByChunkCount(writer, bytes)
	if err != nil {
		return err
	}
	writeFile.Close()
	return nil
}
