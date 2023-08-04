package utils

type Filename struct {
	currentRunes []rune
}

func NewFilename(currentRunes []rune) *Filename {
	return &Filename{currentRunes}
}

// アルファベットを逆順で走査し、アルファベットをインクリメントする
func (f *Filename) Increment() []rune {
	tmp := f.currentRunes
	for i := 1; i <= len(tmp); i++ {
		currentIndex := len(tmp) - i
		currentChar := tmp[currentIndex]
		if currentChar == 'z' {
			f.currentRunes[currentIndex] = 'a'
			if i == len(tmp) {
				f.currentRunes = append(tmp, 'a')
			}
		} else {
			f.currentRunes[currentIndex] = currentChar + 1
			break
		}
	}
	return f.currentRunes
}
