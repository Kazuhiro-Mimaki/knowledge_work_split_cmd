package utils

type NumericFilenameGenerator struct {
	currentRunes []rune
	suffix       string
}

func NewNumericFilenameGenerator(defaultRuneCount int, suffix string) IFilenameGenerator {
	var runes []rune
	if defaultRuneCount == 0 {
		// 指定がない場合はデフォルトで 00 から開始
		runes = []rune{'0', '0'}
	} else {
		// 指定がある場合は defaultRuneCount * 0 から開始
		for i := 0; i < defaultRuneCount; i++ {
			runes = append(runes, '0')
		}
	}
	return &NumericFilenameGenerator{currentRunes: runes, suffix: suffix}
}

func (f *NumericFilenameGenerator) GetCurrentRunes() []rune {
	return f.currentRunes
}

func (f *NumericFilenameGenerator) GetOutputFilePath() string {
	return "./tmp_dir/" + f.suffix + string(f.GetCurrentRunes())
}

// 数字を逆順で走査し、インクリメントする
func (f *NumericFilenameGenerator) Increment() []rune {
	tmp := f.currentRunes
	for i := 1; i <= len(tmp); i++ {
		currentIndex := len(tmp) - i
		currentChar := tmp[currentIndex]
		if currentChar == '9' {
			f.currentRunes[currentIndex] = '0'
			if i == len(tmp) {
				f.currentRunes = append(tmp, '0')
			}
		} else {
			f.currentRunes[currentIndex] = currentChar + 1
			break
		}
	}
	return f.currentRunes
}
