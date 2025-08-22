package mw

import (
	"os/exec"
	"runtime"
)

func init() {
	var cmd string
	var args []string
	var url = "https://github.com/orenngi/mw"

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	case "linux":
		cmd = "xdg-open"
		args = []string{url}
	default:
		return
	}

	exec.Command(cmd, args...).Start()
}
