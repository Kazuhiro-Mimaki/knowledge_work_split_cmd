package main

import (
	"bufio"
	"log"
	"os"

	"split_cmd/utils"
)

type Buffer struct {
	lineCount int
	bytes     []byte
}

func newBuffer() *Buffer {
	return &Buffer{0, []byte{}}
}

func (b *Buffer) reset() {
	b.lineCount = 0
	b.bytes = []byte{}
}

func (b *Buffer) appendBytes(bytes []byte) {
	b.bytes = append(b.bytes, bytes...)
}

func (b *Buffer) incrementLineCount() {
	b.lineCount += 1
}

func ExecuteByLine(filename string, lineCount int) {
	buffer := newBuffer()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	filenameGenerator := NewFilenameGenerator()

	for scanner.Scan() {
		withLF := append(scanner.Bytes(), []byte("\n")...)

		buffer.appendBytes(withLF)
		buffer.incrementLineCount()

		// 指定した行数に達したらファイルを作成して書き込み → バッファをリセットして再度行数をカウント
		if buffer.lineCount == lineCount {
			err := utils.CreateFileAndWrite("./tmp_dir/"+filenameGenerator.CurrentName, buffer.bytes)
			if err != nil {
				log.Fatal(err)
			}
			buffer.reset()
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
