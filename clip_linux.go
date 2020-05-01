package main

import (
	"bytes"
	"log"
	"os/exec"
)

func getClip() ([]byte, error) {
	// TODO what if the linux does not have X (like WSL)
	return exec.Command("xclip", "-selection", "clipboard", "-o").Output()
}

func pasteClip(out []byte) {
	// TODO what if the linux does not have X or xclip(like WSL)
	cmd := exec.Command("xclip", "-selection", "c")
	cmd.Stdin = bytes.NewReader(out)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
