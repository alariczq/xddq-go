package utils

func SubString(s string, start, end int) string {

	runes := make([]rune, 0, end-start)
	for i, s := range []rune(s) {
		if i < start {
			continue
		}
		runes = append(runes, s)
		if i == end-1 {
			break
		}
	}

	return string(runes)
}

func SubLastString(s string, n int) string {
	runes := make([]rune, 0, n)
	for _, s := range []rune(s) {
		runes = append(runes, s)
	}
	return string(runes[len(runes)-n:])
}
