package asciiweb

import (
	"errors"
	//"fmt"
	"strings"
)

func AsciiArt(input string, bannerName string) (string, error) {
	var Result string
	bannerName = strings.ToLower(bannerName)
	if strings.HasPrefix(bannerName, "standard") || strings.HasPrefix(bannerName, "shadow") || strings.HasPrefix(bannerName, "thinkertoy") {
		if !strings.HasSuffix(bannerName, "txt") {
			bannerName += ".txt"
		}
	} else {
		return "", errors.New("Invalid banner")
	}

	input = strings.ReplaceAll(input, "\\t", "   ")
	// input = strings.ReplaceAll(input, "\\n\\r", "\n")

	inputArr := strings.Split(input, "\r\n")
	letters := make([]string, 8)
	for i, word := range inputArr {
		if word == "" {
			check := ""
			if i != 0 {
				check = inputArr[i-1]
			}
			if i < len(inputArr)-1 || check != "" {
				Result += "\n"
			}
			continue
		}

		for _, ch := range word {
			if ch == 13 || ch == 10 {
				continue
			} else if ch < 32 || ch > 126 {
				return "", errors.New("Invalid char")
			} else {
				charLines := GetCharacter(string(ch), bannerName)
				for i, line := range charLines {
					letters[i] += line
				}
			}
		}

		for _, line := range letters {
			Result = Result + line + "\n"
		}
		letters = make([]string, 8)
	}
	return Result, nil
}
