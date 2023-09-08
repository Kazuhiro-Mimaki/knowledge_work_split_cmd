package filename_generator

import "errors"

type Mode int

const (
	Alphabet Mode = iota + 1
	Numeric
)

type FilenameGenerator struct {
	currentRunes []rune
	prefix       string
	mode         Mode
}

func New(initialRuneCount int, prefix string, mode Mode) (*FilenameGenerator, error) {
	initialRune, _, err := runeByMode(mode)
	if err != nil {
		return nil, err
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

	return &FilenameGenerator{currentRunes: runes, prefix: prefix, mode: mode}, nil
}

func (f FilenameGenerator) GetCurrentWithPrefix() string {
	return f.prefix + string(f.currentRunes)
}

// ファイル名をインクリメントする (aa -> ab -> ac -> ... -> az -> ba -> bb -> ...)
func (f *FilenameGenerator) Increment() ([]rune, error) {
	initialRune, lastRune, err := runeByMode(f.mode)
	if err != nil {
		return []rune{}, err
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
	return f.currentRunes, nil
}

// 指定されたモードに従い、アルファベットか数字を判断し、閾値を返す
func runeByMode(mode Mode) (rune, rune, error) {
	switch mode {
	case Alphabet:
		return 'a', 'z', nil
	case Numeric:
		return '0', '9', nil
	default:
		return 0, 0, errors.New("invalid mode")
	}
}
