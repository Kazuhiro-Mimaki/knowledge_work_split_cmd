package split

import (
	"bufio"
	"fmt"
	"io"

	"split_cmd/file_io"
	"split_cmd/filename_generator"
)

func ByLine(r io.Reader, lineCount int, filenameGenerator filename_generator.FilenameGenerator) error {
	scanner := bufio.NewScanner(r)
	buffer := NewScannerBuffer()

	for scanner.Scan() {
		withLF := append(scanner.Bytes(), []byte("\n")...)

		buffer.AppendBytes(withLF)
		buffer.IncrementLineCount()

		// 指定した行数に達したらファイルを作成して書き込み → バッファをリセットして再度行数をカウント
		if buffer.lineCount == lineCount {
			err := file_io.CreateFileAndWrite(filenameGenerator.GetCurrentWithPrefix(), buffer.bytes)
			if err != nil {
				return fmt.Errorf("SplitByLine: error when create and write file in loop: %s", err)
			}
			buffer.Reset()
			filenameGenerator.Increment()
		}
	}

	if buffer.lineCount > 0 {
		err := file_io.CreateFileAndWrite(filenameGenerator.GetCurrentWithPrefix(), buffer.bytes)
		if err != nil {
			return fmt.Errorf("SplitByLine: error when create and write file: %s", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("SplitByLine: error when scan file: %s", err)
	}

	return nil
}
