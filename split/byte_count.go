package split

import (
	"bufio"
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
			return err
		}

		err = file_io.CreateFileAndWrite(filenameGenerator.GetCurrentWithPrefix(), bytes)
		if err != nil {
			return err
		}

		filenameGenerator.Increment()

		// Check for EOF
		if err == nil && cursor < byteCount {
			break
		}
	}

	return nil
}
