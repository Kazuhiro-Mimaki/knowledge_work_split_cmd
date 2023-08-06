package cmd

import (
	"bufio"
	"fmt"
	"os"

	"split_cmd/file_io"
	"split_cmd/output_filename"
)

func ExecuteByChunk(readFilePath string, chunkCount int, filenameGenerator output_filename.FilenameGenerator) error {
	readFile, err := os.Open(readFilePath)
	if err != nil {
		return fmt.Errorf("ExecuteByChunk: error when opening file: %s", err)
	}
	defer readFile.Close()

	fileSize, err := file_io.GetFileSize(readFile)
	if err != nil {
		return fmt.Errorf("ExecuteByChunk: error when get file size %d files", err)
	}
	if fileSize < chunkCount {
		return fmt.Errorf("ExecuteByChunk: can't split into more than %d files", fileSize)
	}
	byteCount, rest := fileSize/chunkCount, fileSize%chunkCount

	reader := bufio.NewReader(readFile)

	for i := 1; i < chunkCount; i++ {
		chunks, _, err := file_io.ReadChunksByByteCount(reader, byteCount)
		if err != nil {
			return fmt.Errorf("ExecuteByChunk: error when read chunks by byte count in loop : %s", err)
		}

		err = file_io.CreateFileAndWrite(filenameGenerator.GetOutputFilePath(), chunks)
		if err != nil {
			return fmt.Errorf("ExecuteByChunk: error when create and write file in loop : %s", err)
		}

		filenameGenerator.Increment()
	}

	chunks, _, err := file_io.ReadChunksByByteCount(reader, byteCount+rest)
	if err != nil {
		return fmt.Errorf("ExecuteByChunk: error when read chunks by byte count : %s", err)
	}

	err = file_io.CreateFileAndWrite(filenameGenerator.GetOutputFilePath(), chunks)
	if err != nil {
		return fmt.Errorf("ExecuteByChunk: error when create and write file : %s", err)
	}

	return nil
}
