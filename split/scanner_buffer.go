package split

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

func (b *ScannerBuffer) Increment(bytes []byte) {
	b.appendBytes(bytes)
	b.incrementLineCount()
}

func (b *ScannerBuffer) appendBytes(bytes []byte) {
	b.bytes = append(b.bytes, bytes...)
}

func (b *ScannerBuffer) incrementLineCount() {
	b.lineCount += 1
}
