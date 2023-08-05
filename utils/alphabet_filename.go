package utils

type AlphabetFilenameGenerator struct {
	currentRunes []rune
}

func NewAlphabetFilenameGenerator(defaultRuneCount int) IFilenameGenerator {
	var runes []rune
	if defaultRuneCount == 0 {
		// 指定がない場合はデフォルトで aa から開始
		runes = []rune{'a', 'a'}
	} else {
		// 指定がある場合は defaultRuneCount * a から開始
		for i := 0; i < defaultRuneCount; i++ {
			runes = append(runes, 'a')
		}
	}
	return &AlphabetFilenameGenerator{runes}
}

func (f *AlphabetFilenameGenerator) GetCurrentRunes() []rune {
	return f.currentRunes
}

// アルファベットを逆順で走査し、インクリメントする
func (f *AlphabetFilenameGenerator) Increment() []rune {
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
