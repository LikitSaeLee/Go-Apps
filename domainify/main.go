package main

import (
	"fmt"
	"os"
	"bufio"
	"math/rand"
	"strings"
	"unicode"
)

var tlds = []string{"com", "net"}
const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func main() {
	stdin := os.Stdin
	s := bufio.NewScanner(stdin)
	for s.Scan() {
		domain := s.Text()
		var newText []rune
		for _, r := range domain {
			if unicode.IsSpace(r) {
				r = '-'
			} else if !strings.ContainsRune(allowedChars, r) {
				continue
			}
			newText = append(newText, r)
		}
		fmt.Println(string(newText) + "." + tlds[rand.Intn(len(tlds))])
	}
}
