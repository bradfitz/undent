package main

import (
	"bytes"
	"log"
	"os/exec"
)

func getClip() ([]byte, error) {
	// TODO what if the linux does not have X (like WSL)
	return exec.Command("xsel").Output()
}

func pasteClip(out []byte) {
	// TODO what if the linux does not have X (like WSL)
	cmd := exec.Command("xsel")
	cmd.Stdin = bytes.NewReader(out)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
