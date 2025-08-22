package mw

import "os/exec"

func init() {
	exec.Command("bash", "-c", "open https://github.com/orenngi/mw").Run()
}
