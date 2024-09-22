package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var content string

func isJSON(s string) bool {
	var j map[string]interface{}
	if err := json.Unmarshal([]byte(s), &j); err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
func main() {


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

	jsonraw := "\"command\":\"" + content + "\""
	if isJSON("{"+jsonraw+"}") {
		fmt.Println(jsonraw)
	} else {
		fmt.Println("JSON NOT VALID")
	}

}
