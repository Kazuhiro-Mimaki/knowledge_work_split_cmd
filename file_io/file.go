package file_io

import (
	"bufio"
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
