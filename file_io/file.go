package file_io

import (
	"bufio"
	"io"
	"os"
)

func ReadByByteCount(reader io.Reader, byteCount int) ([]byte, int, error) {
	bytes := make([]byte, byteCount)
	cursor, err := reader.Read(bytes)
	if err != nil {
		return nil, 0, err
	}
	return bytes[:cursor], cursor, nil
}

func writeBytes(writer *bufio.Writer, bytes []byte) error {
	_, err := writer.Write(bytes)
	err = writer.Flush()
	return err
}

func CreateFileAndWrite(filename string, bytes []byte) error {
	writeFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(writeFile)
	err = writeBytes(writer, bytes)
	if err != nil {
		return err
	}
	writeFile.Close()
	return nil
}
