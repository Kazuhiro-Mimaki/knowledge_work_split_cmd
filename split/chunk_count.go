package split

import (
	"bufio"
	"errors"
	"io"

	"split_cmd/file_io"
	"split_cmd/filename_generator"
)

func ByChunk(r io.Reader, fileSize, chunkCount int, filenameGenerator filename_generator.FilenameGenerator) error {
	if fileSize < chunkCount {
		return errors.New("SplitByChunk: can't split file by chunk because file size is smaller than chunk count")
	}
	byteCount, rest := fileSize/chunkCount, fileSize%chunkCount

	reader := bufio.NewReader(r)

	for i := 1; i < chunkCount; i++ {
		bytes, _, err := file_io.ReadByByteCount(reader, byteCount)
		if err != nil {
			return err
		}

		err = file_io.CreateFileAndWrite(filenameGenerator.GetCurrentWithPrefix(), bytes)
		if err != nil {
			return err
		}

		filenameGenerator.Increment()
	}

	bytes, _, err := file_io.ReadByByteCount(reader, byteCount+rest)
	if err != nil {
		return err
	}

	err = file_io.CreateFileAndWrite(filenameGenerator.GetCurrentWithPrefix(), bytes)
	if err != nil {
		return err
	}

	return nil
}
