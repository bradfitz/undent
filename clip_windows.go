package main

import (
	"bytes"
	"log"
	"os/exec"
)

func getClip() ([]byte, error) {
	// this is hacky, but the way simpler than in 
	// https://github.com/atotto/clipboard/blob/master/clipboard_windows.go
	return exec.Command("powershell.exe", "Get-ClipBoard").Output()
}

func pasteClip(out []byte) {
	cmd := exec.Command("clip")
	cmd.Stdin = bytes.NewReader(out)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
