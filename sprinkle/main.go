package main

import (
	"fmt"
	"os"
	"math/rand"
	"bufio"
	"strings"
)

var words = []string{
	"*app",
	"*web",
	"alpha-*",
	"fun*",
	"guru*",
}

func main() {
	stdin := os.Stdin
	scaner := bufio.NewScanner( stdin )
	fmt.Println( "Enter your app name:" )
	for scaner.Scan() {
		fmt.Println( "Enter your app name:" )
		text := scaner.Text()
		rand_word := words[rand.Intn(len(words))]
		fmt.Println( strings.Replace( rand_word, "*", text, -1 ) )
	}
}


