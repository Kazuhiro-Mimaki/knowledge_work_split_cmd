package filename_generator

type AlphabetFilenameGenerator struct {
	current []rune
	prefix  string
}

func NewAlphabetFilenameGenerator(suffixLength int, prefix string) FilenameGenerator {
	var current []rune
	if suffixLength == 0 {
		// 指定がない場合はデフォルト (aa) から開始
		current = []rune{'a', 'a'}
	} else {
		// 指定がある場合は suffixLength * (a) から開始
		for i := 0; i < suffixLength; i++ {
			current = append(current, 'a')
		}
	}

	return &AlphabetFilenameGenerator{current, prefix}
}

func (f *AlphabetFilenameGenerator) GetCurrentWithPrefix() string {
	return f.prefix + string(f.current)
}

// ファイル名をインクリメントする (aa -> ab -> ac -> ... -> az -> ba -> bb -> ...)
func (f *AlphabetFilenameGenerator) Increment() []rune {
	nameLength := len(f.current)
	for i := 1; i <= nameLength; i++ {
		currentIndex := nameLength - i
		currentChar := f.current[currentIndex]
		if currentChar == 'z' {
			f.current[currentIndex] = 'a'
			if i == nameLength {
				f.current = append(f.current, 'a')
			}
		} else {
			f.current[currentIndex] = currentChar + 1
			break
		}
	}
	return f.current
}
