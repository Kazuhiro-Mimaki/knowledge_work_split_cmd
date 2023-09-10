package file_io

import (
	"bufio"
	"bytes"
	"reflect"
	"testing"
)

func TestReadByByteCount(t *testing.T) {
	t.Run("正常系 (バイト列の最初から2つを読み取る)", func(t *testing.T) {
		mockBuff := bytes.NewBuffer([]byte{'A', 'B', 'C'})
		reader := bufio.NewReader(mockBuff)

		bytes, cursor, _ := ReadByByteCount(reader, 2)

		wantBytes := []byte{'A', 'B'}
		wantCursor := 2

		if string(bytes) != string(wantBytes) {
			t.Errorf("bytes == %v, want %s", bytes, wantBytes)
		}
		if cursor != wantCursor {
			t.Errorf("cursor == %d, want %d", cursor, wantCursor)
		}
	})

	t.Run("正常系 (バイト列の最初から5つを読み取る)", func(t *testing.T) {
		mockBuff := bytes.NewBuffer([]byte{'A', 'B', 'C'})
		reader := bufio.NewReader(mockBuff)

		bytes, cursor, _ := ReadByByteCount(reader, 5)

		wantBytes := []byte{'A', 'B', 'C'}
		wantCursor := 3

		if string(bytes) != string(wantBytes) {
			t.Errorf("bytes == %v, want %s", bytes, wantBytes)
		}
		if cursor != wantCursor {
			t.Errorf("cursor == %d, want %d", cursor, wantCursor)
		}
	})

	t.Run("正常系 (バイト列の終端まで繰り返し読み取る)", func(t *testing.T) {
		mockBuff := bytes.NewBuffer([]byte{'A', 'B', 'C'})
		reader := bufio.NewReader(mockBuff)

		byteBuf := [][]byte{}

		for {
			bytes, cursor, _ := ReadByByteCount(reader, 1)
			if cursor < 1 {
				break
			}
			byteBuf = append(byteBuf, bytes)
		}

		want := [][]byte{{'A'}, {'B'}, {'C'}}

		if !reflect.DeepEqual(byteBuf, want) {
			t.Errorf("byteBuf == %v, want %s", byteBuf, want)
		}
	})
}

func TestWriteBytes(t *testing.T) {
	t.Run("正常系 (バイト列を書き込む)", func(t *testing.T) {
		outputBuff := bytes.NewBuffer([]byte{})
		writer := bufio.NewWriter(outputBuff)

		_ = writeBytes(writer, []byte{'A', 'B', 'C'})

		want := []byte{'A', 'B', 'C'}

		if string(outputBuff.Bytes()) != string(want) {
			t.Errorf("output == %s, want %s", outputBuff.Bytes(), want)
		}
	})

	t.Run("正常系 (バイト列が空の場合)", func(t *testing.T) {
		outputBuff := bytes.NewBuffer([]byte{})
		writer := bufio.NewWriter(outputBuff)

		_ = writeBytes(writer, []byte{})

		want := []byte{}

		if string(outputBuff.Bytes()) != string(want) {
			t.Errorf("output == %s, want %s", outputBuff.Bytes(), want)
		}
	})
}
