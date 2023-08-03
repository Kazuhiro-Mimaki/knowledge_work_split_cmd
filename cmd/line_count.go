package cmd

import (
	"bufio"
	"fmt"
	"os"

	"split_cmd/utils"
)

func ExecuteByLine(filename, suffix string, lineCount int) error {
	buffer := NewScannerBuffer()

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("ExecuteByLine: error when opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	filenameGenerator := utils.NewFilenameGenerator()

	for scanner.Scan() {
		withLF := append(scanner.Bytes(), []byte("\n")...)

		buffer.AppendBytes(withLF)
		buffer.IncrementLineCount()

		// 指定した行数に達したらファイルを作成して書き込み → バッファをリセットして再度行数をカウント
		if buffer.lineCount == lineCount {
			err := utils.CreateFileAndWrite("./tmp_dir/"+suffix+filenameGenerator.CurrentName, buffer.bytes)
			if err != nil {
				return fmt.Errorf("ExecuteByLine: error when create and write file in loop: %s", err)
			}
			buffer.Reset()
			filenameGenerator.Increment()
		}
	}

	if buffer.lineCount > 0 {
		err := utils.CreateFileAndWrite("./tmp_dir/"+filenameGenerator.CurrentName, buffer.bytes)
		if err != nil {
			return fmt.Errorf("ExecuteByLine: error when create and write file: %s", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ExecuteByLine: error when scan file: %s", err)
	}

	return nil
}
