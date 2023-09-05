package filename_generator

type Mode int

const (
	ALPHABET Mode = iota + 1
	NUMERIC
)

type FilenameGenerator struct {
	currentRunes []rune
	suffix       string
	mode         Mode
}

func New(initialRuneCount int, suffix string, mode Mode) FilenameGenerator {
	var initialRune rune

	switch mode {
	case ALPHABET:
		initialRune = 'a'
	case NUMERIC:
		initialRune = '0'
	default:
		panic("invalid mode")
	}

	var runes []rune
	if initialRuneCount == 0 {
		// 指定がない場合はデフォルト (aa | 00) から開始
		runes = []rune{initialRune, initialRune}
	} else {
		// 指定がある場合は initialRuneCount * (a | 0) から開始
		for i := 0; i < initialRuneCount; i++ {
			runes = append(runes, initialRune)
		}
	}

	return FilenameGenerator{currentRunes: runes, suffix: suffix, mode: mode}
}

func (f FilenameGenerator) GetCurrent() []rune {
	return f.currentRunes
}

func (f FilenameGenerator) GetCurrentWithSuffix() string {
	return f.suffix + string(f.GetCurrent())
}

func (f FilenameGenerator) GetMode() Mode {
	return f.mode
}

// アルファベットまたは数字 を逆順で走査し、インクリメントする
func (f *FilenameGenerator) Increment() []rune {
	var initialRune, lastRune rune

	switch f.GetMode() {
	case ALPHABET:
		initialRune, lastRune = 'a', 'z'
	case NUMERIC:
		initialRune, lastRune = '0', '9'
	default:
		panic("invalid mode")
	}

	nameLength := len(f.currentRunes)
	for i := 1; i <= nameLength; i++ {
		currentIndex := nameLength - i
		currentChar := f.currentRunes[currentIndex]
		if currentChar == lastRune {
			f.currentRunes[currentIndex] = initialRune
			if i == nameLength {
				f.currentRunes = append(f.currentRunes, initialRune)
			}
		} else {
			f.currentRunes[currentIndex] = currentChar + 1
			break
		}
	}
	return f.currentRunes
}
