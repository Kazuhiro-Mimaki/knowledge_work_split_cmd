package cmd

import (
	"bufio"
	"fmt"
	"io"

	"split_cmd/file_io"
	"split_cmd/filename_generator"
)

func ExecuteByChunk(r io.Reader, fileSize, chunkCount int, filenameGenerator filename_generator.FilenameGenerator) error {
	if fileSize < chunkCount {
		return fmt.Errorf("ExecuteByChunk: can't split into more than %d files", fileSize)
	}
	byteCount, rest := fileSize/chunkCount, fileSize%chunkCount

	reader := bufio.NewReader(r)

	for i := 1; i < chunkCount; i++ {
		chunks, _, err := file_io.ReadChunksByByteCount(reader, byteCount)
		if err != nil {
			return fmt.Errorf("ExecuteByChunk: error when read chunks by byte count in loop : %s", err)
		}

		err = file_io.CreateFileAndWrite(filenameGenerator.GetCurrentWithPrefix(), chunks)
		if err != nil {
			return fmt.Errorf("ExecuteByChunk: error when create and write file in loop : %s", err)
		}

		filenameGenerator.Increment()
	}

	chunks, _, err := file_io.ReadChunksByByteCount(reader, byteCount+rest)
	if err != nil {
		return fmt.Errorf("ExecuteByChunk: error when read chunks by byte count : %s", err)
	}

	err = file_io.CreateFileAndWrite(filenameGenerator.GetCurrentWithPrefix(), chunks)
	if err != nil {
		return fmt.Errorf("ExecuteByChunk: error when create and write file : %s", err)
	}

	return nil
}
