package split

import (
	"bufio"
	"fmt"
	"io"

	"split_cmd/file_io"
	"split_cmd/filename_generator"
)

func ByByteCount(r io.Reader, byteCount int, filenameGenerator filename_generator.FilenameGenerator) error {
	reader := bufio.NewReader(r)

	for {
		bytes, cursor, err := file_io.ReadByByteCount(reader, byteCount)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("SplitByteCount: error when read bytes by byte count in loop: %s", err)
		}

		err = file_io.CreateFileAndWrite(filenameGenerator.GetCurrentWithPrefix(), bytes)
		if err != nil {
			return fmt.Errorf("SplitByteCount: error when create and write file by byte count: %s", err)
		}

		filenameGenerator.Increment()

		// Check for EOF
		if err == nil && cursor < byteCount {
			break
		}
	}

	return nil
}
