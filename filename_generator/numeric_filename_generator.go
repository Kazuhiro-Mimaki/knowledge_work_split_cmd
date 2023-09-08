package filename_generator

type NumericFilenameGenerator struct {
	current []rune
	prefix  string
}

func NewNumericFilenameGenerator(suffixLength int, prefix string) FilenameGenerator {
	var current []rune
	if suffixLength == 0 {
		// 指定がない場合はデフォルト (aa) から開始
		current = []rune{'0', '0'}
	} else {
		// 指定がある場合は suffixLength * (a) から開始
		for i := 0; i < suffixLength; i++ {
			current = append(current, '0')
		}
	}

	return &NumericFilenameGenerator{current, prefix}
}

func (f *NumericFilenameGenerator) GetCurrentWithPrefix() string {
	return f.prefix + string(f.current)
}

// ファイル名をインクリメントする (aa -> ab -> ac -> ... -> az -> ba -> bb -> ...)
func (f *NumericFilenameGenerator) Increment() []rune {
	nameLength := len(f.current)
	for i := 1; i <= nameLength; i++ {
		currentIndex := nameLength - i
		currentChar := f.current[currentIndex]
		if currentChar == '9' {
			f.current[currentIndex] = '0'
			if i == nameLength {
				f.current = append(f.current, '0')
			}
		} else {
			f.current[currentIndex] = currentChar + 1
			break
		}
	}
	return f.current
}
