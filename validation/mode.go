package validation

/**
 * l or n or b は1種類のみしか指定できない
 */
func Mode(l, n, b int) string {
	if l != 0 && n == 0 && b == 0 {
		return "l"
	} else if l == 0 && n != 0 && b == 0 {
		return "n"
	} else if l == 0 && n == 0 && b != 0 {
		return "b"
	} else if l == 0 && n == 0 && b == 0 {
		return "noArgs"
	} else {
		return ""
	}
}
