package cmd

import (
	"bufio"
	"fmt"
	"os"

	"split_cmd/utils"
)

func ExecuteByLine(filename, suffix string, suffixLength, lineCount int) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("ExecuteByLine: error when opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	filenameManager := utils.NewFilenameManager(suffixLength)
	buffer := NewScannerBuffer()

	for scanner.Scan() {
		withLF := append(scanner.Bytes(), []byte("\n")...)

		buffer.AppendBytes(withLF)
		buffer.IncrementLineCount()

		// 指定した行数に達したらファイルを作成して書き込み → バッファをリセットして再度行数をカウント
		if buffer.lineCount == lineCount {
			err := utils.CreateFileAndWrite("./tmp_dir/"+suffix+string(filenameManager.CurrentRunes), buffer.bytes)
			if err != nil {
				return fmt.Errorf("ExecuteByLine: error when create and write file in loop: %s", err)
			}
			buffer.Reset()
			filenameManager.Increment()
		}
	}

	if buffer.lineCount > 0 {
		err := utils.CreateFileAndWrite("./tmp_dir/"+suffix+string(filenameManager.CurrentRunes), buffer.bytes)
		if err != nil {
			return fmt.Errorf("ExecuteByLine: error when create and write file: %s", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ExecuteByLine: error when scan file: %s", err)
	}

	return nil
}
