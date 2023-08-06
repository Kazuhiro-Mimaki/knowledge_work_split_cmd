package output_filename

type NumericFilenameGenerator struct {
	currentRunes []rune
	suffix       string
}

func NewNumericFilenameGenerator(defaultRuneCount int, suffix string) FilenameGenerator {
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
	return f.suffix + string(f.GetCurrentRunes())
}

// 数字を逆順で走査し、インクリメントする
func (f *NumericFilenameGenerator) Increment() []rune {
	nameLength := len(f.currentRunes)
	for i := 1; i <= nameLength; i++ {
		currentIndex := nameLength - i
		currentChar := f.currentRunes[currentIndex]
		if currentChar == '9' {
			f.currentRunes[currentIndex] = '0'
			if i == nameLength {
				f.currentRunes = append(f.currentRunes, '0')
			}
		} else {
			f.currentRunes[currentIndex] = currentChar + 1
			break
		}
	}
	return f.currentRunes
}
