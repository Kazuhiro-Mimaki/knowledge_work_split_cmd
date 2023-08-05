package cmd

import (
	"bufio"
	"fmt"
	"os"

	"split_cmd/utils"
)

func ExecuteByteCount(filename, suffix string, byteCount int, filenameManager utils.IFilenameGenerator) error {
	readFile, err := os.Open(filename)
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

		err = utils.CreateFileAndWrite("./tmp_dir/"+suffix+string(filenameManager.GetCurrentRunes()), chunks)
		if err != nil {
			return fmt.Errorf("ExecuteByteCount: error when create and write file by byte count: %s", err)
		}

		filenameManager.Increment()

		// Check for EOF
		if err == nil && cursor < byteCount {
			break
		}
	}

	return nil
}
