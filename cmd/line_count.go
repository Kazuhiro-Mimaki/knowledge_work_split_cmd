package cmd

import (
	"bufio"
	"log"
	"os"

	"split_cmd/utils"
)

func ExecuteByLine(filename string, lineCount int) {
	buffer := NewScannerBuffer()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
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
			err := utils.CreateFileAndWrite("./tmp_dir/"+filenameGenerator.CurrentName, buffer.bytes)
			if err != nil {
				log.Fatal(err)
			}
			buffer.Reset()
			filenameGenerator.Increment()
		}
	}

	if buffer.lineCount > 0 {
		err := utils.CreateFileAndWrite("./tmp_dir/"+filenameGenerator.CurrentName, buffer.bytes)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while scan file: %s", err)
	}
}
