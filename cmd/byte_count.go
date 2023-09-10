package cmd

import (
	"bufio"
	"fmt"
	"io"

	"split_cmd/file_io"
	"split_cmd/filename_generator"
)

func ExecuteByteCount(r io.Reader, byteCount int, filenameGenerator filename_generator.FilenameGenerator) error {
	reader := bufio.NewReader(r)

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
