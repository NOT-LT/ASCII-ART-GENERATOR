package asciiweb

func GetCharacter(word string, banner string) []string {
	lines := make([]string, 8)
	for i := 0; i < 8; i++ {
		for _, ch := range word {
			charLines := GetLine(1 + int(ch-32)*9 + i, banner)
			lines[i] += charLines
		}
	}
	return lines
}
