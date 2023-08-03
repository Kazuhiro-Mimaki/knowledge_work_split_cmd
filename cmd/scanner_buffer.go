package cmd

type ScannerBuffer struct {
	lineCount int
	bytes     []byte
}

func NewScannerBuffer() *ScannerBuffer {
	return &ScannerBuffer{0, []byte{}}
}

func (b *ScannerBuffer) Reset() {
	b.lineCount = 0
	b.bytes = []byte{}
}

func (b *ScannerBuffer) AppendBytes(bytes []byte) {
	b.bytes = append(b.bytes, bytes...)
}

func (b *ScannerBuffer) IncrementLineCount() {
	b.lineCount += 1
}
