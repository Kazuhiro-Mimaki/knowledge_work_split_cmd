package utils

import (
	"bufio"
	"errors"
	"os"
)

func ReadChunksByByteCount(reader *bufio.Reader, byteCount int) ([]byte, int, error) {
	chunks := make([]byte, byteCount)
	cursor, err := reader.Read(chunks)
	return chunks[:cursor], cursor, err
}

func writeChunks(writer *bufio.Writer, chunks []byte) error {
	_, err := writer.Write(chunks)
	err = writer.Flush()
	return err
}

func GetFileSize(file *os.File) (int, error) {
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

func CreateFileAndWrite(filename string, bytes []byte) error {
	writeFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(writeFile)
	err = writeChunks(writer, bytes)
	if err != nil {
		return err
	}
	writeFile.Close()
	return nil
}
