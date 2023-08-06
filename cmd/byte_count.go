package cmd

import (
	"bufio"
	"fmt"
	"os"

	"split_cmd/output_filename"
	"split_cmd/utils"
)

func ExecuteByteCount(readFilePath string, byteCount int, filenameGenerator output_filename.FilenameGenerator) error {
	readFile, err := os.Open(readFilePath)
	if err != nil {
		return fmt.Errorf("ExecuteByteCount: error when opening file: %s", err)
	}
	defer readFile.Close()

	reader := bufio.NewReader(readFile)

	for {
		chunks, cursor, err := utils.ReadChunksByByteCount(reader, byteCount)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("ExecuteByteCount: error when read chunks by byte count in loop: %s", err)
		}

		err = utils.CreateFileAndWrite(filenameGenerator.GetOutputFilePath(), chunks)
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
