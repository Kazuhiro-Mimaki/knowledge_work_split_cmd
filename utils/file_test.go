package utils

import (
	"bufio"
	"bytes"
	"reflect"
	"testing"
)

func TestReadChunksByByteCount(t *testing.T) {
	t.Run("バイト列の最初から2つを読み取る", func(t *testing.T) {
		mockBuff := bytes.NewBuffer([]byte{'A', 'B', 'C'})
		reader := bufio.NewReader(mockBuff)

		chunks, cursor, _ := ReadChunksByByteCount(reader, 2)

		wantChunks := []byte{'A', 'B'}
		wantCursor := 2

		if string(chunks) != string(wantChunks) {
			t.Errorf("chunks == %v, want %s", chunks, wantChunks)
		}
		if cursor != wantCursor {
			t.Errorf("cursor == %d, want %d", cursor, wantCursor)
		}
	})
	t.Run("バイト列の最初から5つを読み取る", func(t *testing.T) {
		mockBuff := bytes.NewBuffer([]byte{'A', 'B', 'C'})
		reader := bufio.NewReader(mockBuff)

		chunks, cursor, _ := ReadChunksByByteCount(reader, 5)

		wantChunks := []byte{'A', 'B', 'C'}
		wantCursor := 3

		if string(chunks) != string(wantChunks) {
			t.Errorf("chunks == %v, want %s", chunks, wantChunks)
		}
		if cursor != wantCursor {
			t.Errorf("cursor == %d, want %d", cursor, wantCursor)
		}
	})
	t.Run("バイト列の終端まで繰り返し読み取る", func(t *testing.T) {
		mockBuff := bytes.NewBuffer([]byte{'A', 'B', 'C'})
		reader := bufio.NewReader(mockBuff)

		bytes := [][]byte{}

		for {
			chunks, cursor, _ := ReadChunksByByteCount(reader, 1)
			if cursor < 1 {
				break
			}
			bytes = append(bytes, chunks)
		}

		want := [][]byte{{'A'}, {'B'}, {'C'}}

		if !reflect.DeepEqual(bytes, want) {
			t.Errorf("bytes == %v, want %s", bytes, want)
		}
	})
}

func TestWriteChunks(t *testing.T) {
	t.Run("バイト列を書き込む", func(t *testing.T) {
		outputBuff := bytes.NewBuffer([]byte{})
		writer := bufio.NewWriter(outputBuff)

		_ = writeChunks(writer, []byte{'A', 'B', 'C'})

		want := []byte{'A', 'B', 'C'}

		if string(outputBuff.Bytes()) != string(want) {
			t.Errorf("output == %s, want %s", outputBuff.Bytes(), want)
		}
	})
	t.Run("バイト列が空の場合", func(t *testing.T) {
		outputBuff := bytes.NewBuffer([]byte{})
		writer := bufio.NewWriter(outputBuff)

		_ = writeChunks(writer, []byte{})

		want := []byte{}

		if string(outputBuff.Bytes()) != string(want) {
			t.Errorf("output == %s, want %s", outputBuff.Bytes(), want)
		}
	})
}
