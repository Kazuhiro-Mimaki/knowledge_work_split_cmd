package main

import (
	"bufio"
	"log"
	"os"
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
			err := createAndWrite("./tmp_dir/"+filenameGenerator.CurrentName, buffer.bytes)
			if err != nil {
				log.Fatal(err)
			}
			buffer.reset()
			filenameGenerator.Increment()
		}
	}

	if buffer.lineCount > 0 {
		err := createAndWrite("./tmp_dir/"+filenameGenerator.CurrentName, buffer.bytes)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while scan file: %s", err)
	}
}

func writeChunksByLineCount(writer *bufio.Writer, chunks []byte) error {
	_, err := writer.Write(chunks)
	err = writer.Flush()
	return err
}

func createAndWrite(filename string, bytes []byte) error {
	writeFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(writeFile)
	err = writeChunksByLineCount(writer, bytes)
	if err != nil {
		return err
	}
	writeFile.Close()
	return nil
}
