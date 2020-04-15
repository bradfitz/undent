package main

import (
	"os/exec"
)

func getClip() ([]byte, error) {
	return exec.Command("clip").Output()
}

func pasteClip(out []byte) {
	// TODO fallback to powershell?
	panic("unimplemented on windows")
}
