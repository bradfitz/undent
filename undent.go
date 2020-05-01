package main

import (
	"bytes"
	"flag"
	"log"
	"os"

	"github.com/atotto/clipboard"
)

var clip = flag.Bool("clip", false, "fix clipboard; else use stdin/stdout or named file")

func main() {
	flag.Parse()

	var in, out []byte
	if *clip {
		str, err := clipboard.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		in = []byte(str)
	}
	lines := bytes.Split(in, []byte("\n"))
	for len(lines) > 0 && linesStartWith(firstNonEmptyByte(lines), lines) {
		for i := range lines {
			if len(lines[i]) != 0 {
				lines[i] = lines[i][1:]
			}
		}
	}
	out = bytes.Join(lines, []byte("\n"))
	if *clip {
		clipboard.WriteAll(string(out))
	} else {
		os.Stdout.Write(out)
	}
}

// returning 0 means stop
func firstNonEmptyByte(lines [][]byte) byte {
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		switch line[0] {
		case ' ', '\t':
			return line[0]
		default:
			return 0
		}
	}
	return 0
}

func linesStartWith(b byte, lines [][]byte) bool {
	if b == 0 {
		return false
	}
	for _, line := range lines {
		if len(line) == 0 || line[0] == b {
			continue
		}
		return false
	}
	return true
}
