package main

import (
	"fmt"
	"os"
	"strings"
)

var content string

func main() {
	// read file
	raw, err := os.ReadFile("../script.ps1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, line := range strings.Split(string(raw), "\n") {
		b := strings.TrimSpace(line)
		c := strings.ReplaceAll(b, "\"", "'")
		d := strings.ReplaceAll(c, "\\", "\\\\")
		content = content + d
	}

	fmt.Println("{\"command\":\"" + content + "\"}")
}
