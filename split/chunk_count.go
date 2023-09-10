package split

import (
	"bufio"
	"fmt"
	"io"

	"split_cmd/file_io"
	"split_cmd/filename_generator"
)

func ByChunk(r io.Reader, fileSize, chunkCount int, filenameGenerator filename_generator.FilenameGenerator) error {
	if fileSize < chunkCount {
		return fmt.Errorf("SplitByChunk: can't split into more than %d files", fileSize)
	}
	byteCount, rest := fileSize/chunkCount, fileSize%chunkCount

	reader := bufio.NewReader(r)

	for i := 1; i < chunkCount; i++ {
		bytes, _, err := file_io.ReadByByteCount(reader, byteCount)
		if err != nil {
			return fmt.Errorf("SplitByChunk: error when read bytes by byte count in loop : %s", err)
		}

		err = file_io.CreateFileAndWrite(filenameGenerator.GetCurrentWithPrefix(), bytes)
		if err != nil {
			return fmt.Errorf("SplitByChunk: error when create and write file in loop : %s", err)
		}

		filenameGenerator.Increment()
	}

	bytes, _, err := file_io.ReadByByteCount(reader, byteCount+rest)
	if err != nil {
		return fmt.Errorf("SplitByChunk: error when read bytes by byte count : %s", err)
	}

	err = file_io.CreateFileAndWrite(filenameGenerator.GetCurrentWithPrefix(), bytes)
	if err != nil {
		return fmt.Errorf("SplitByChunk: error when create and write file : %s", err)
	}

	return nil
}
