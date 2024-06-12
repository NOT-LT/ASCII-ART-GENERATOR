package asciiweb

import (
	"log"
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Start()

	case "darwin":
		err = exec.Command("open", url).Start()

	case "linux":
		err = exec.Command("xdg-open", url).Start()

	default:
		log.Printf("Unsupported platform: %s", runtime.GOOS)
		return
	}

	if err != nil {
		log.Printf("Failed to open browser: %v", err)
	}
}
