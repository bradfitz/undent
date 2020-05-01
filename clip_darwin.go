package main

import (
	"bytes"
	"log"
	"os/exec"
)

func getClip() ([]byte, error) {
	return exec.Command("pbpaste").Output()
}

func pasteClip(out []byte) {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = bytes.NewReader(out)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
