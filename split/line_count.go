package split

import (
	"bufio"
	"io"

	"split_cmd/file_io"
	"split_cmd/filename_generator"
)

func ByLine(r io.Reader, lineCount int, filenameGenerator filename_generator.FilenameGenerator) error {
	scanner := bufio.NewScanner(r)
	buffer := NewScannerBuffer()

	for scanner.Scan() {
		bytes := append(scanner.Bytes(), []byte("\n")...)
		buffer.Increment(bytes)

		// 指定した行数に達したらファイルを作成して書き込み → バッファをリセットして再度行数をカウント
		if buffer.lineCount == lineCount {
			err := file_io.CreateAndWrite(filenameGenerator.GetCurrentWithPrefix(), buffer.bytes)
			if err != nil {
				return err
			}

			buffer.Reset()
			filenameGenerator.Increment()
		}
	}

	// ファイルの最後の行が改行で終わっていない場合は、バッファに残っている行をファイルに書き込む
	if buffer.lineCount > 0 {
		err := file_io.CreateAndWrite(filenameGenerator.GetCurrentWithPrefix(), buffer.bytes)
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
