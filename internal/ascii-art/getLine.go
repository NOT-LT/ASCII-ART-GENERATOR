package asciiweb

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetLine(lineNum int, banner string) string {

	file, err := os.Open("banners/" + banner)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return ""
	}
	defer file.Close()
	content := bufio.NewReader(file)
	x := ""
	for i := 0; i <= lineNum; i++ {
		x, err = content.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading the file:", err)
			break
		}
	}
	return strings.TrimRight(x, "\r\n")
}
