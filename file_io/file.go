package file_io

import (
	"bufio"
	"io"
	"os"
)

func ReadByByteCount(r io.Reader, byteCount int) ([]byte, int, error) {
	bytes := make([]byte, byteCount)
	cursor, err := r.Read(bytes)
	if err != nil {
		return nil, 0, err
	}
	return bytes[:cursor], cursor, nil
}

func writeBytes(bw *bufio.Writer, bytes []byte) error {
	_, err := bw.Write(bytes)
	err = bw.Flush()
	return err
}

func CreateAndWrite(path string, bytes []byte) error {
	w, err := os.Create(path)
	if err != nil {
		return err
	}
	defer w.Close()

	bw := bufio.NewWriter(w)

	err = writeBytes(bw, bytes)
	if err != nil {
		return err
	}

	return nil
}
