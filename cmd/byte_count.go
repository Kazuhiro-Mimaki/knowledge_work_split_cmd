package cmd

import (
	"bufio"
	"fmt"
	"os"

	"split_cmd/file_io"
	"split_cmd/filename_generator"
)

func ExecuteByteCount(readFilePath string, byteCount int, filenameGenerator *filename_generator.FilenameGenerator) error {
	readFile, err := os.Open(readFilePath)
	if err != nil {
		return fmt.Errorf("ExecuteByteCount: error when opening file: %s", err)
	}
	defer readFile.Close()

	reader := bufio.NewReader(readFile)

	for {
		chunks, cursor, err := file_io.ReadChunksByByteCount(reader, byteCount)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("ExecuteByteCount: error when read chunks by byte count in loop: %s", err)
		}

		err = file_io.CreateFileAndWrite(filenameGenerator.GetCurrentWithPrefix(), chunks)
		if err != nil {
			return fmt.Errorf("ExecuteByteCount: error when create and write file by byte count: %s", err)
		}

		filenameGenerator.Increment()

		// Check for EOF
		if err == nil && cursor < byteCount {
			break
		}
	}

	return nil
}
