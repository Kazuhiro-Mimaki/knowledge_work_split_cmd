package utils

type FilenameManager struct {
	CurrentRunes []rune
}

func NewFilenameManager(defaultRuneCount int) *FilenameManager {
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
	return &FilenameManager{runes}
}

// アルファベットを逆順で走査し、アルファベットをインクリメントする
func (f *FilenameManager) Increment() []rune {
	tmp := f.CurrentRunes
	for i := 1; i <= len(tmp); i++ {
		currentIndex := len(tmp) - i
		currentChar := tmp[currentIndex]
		if currentChar == 'z' {
			f.CurrentRunes[currentIndex] = 'a'
			if i == len(tmp) {
				f.CurrentRunes = append(tmp, 'a')
			}
		} else {
			f.CurrentRunes[currentIndex] = currentChar + 1
			break
		}
	}
	return f.CurrentRunes
}
