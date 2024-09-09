package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "World"

	words := []string{"Andre", "Luiz", "Rodrigues", "Alves"}

	fullWord := strings.Join(words, " ")

	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who)
	fmt.Println("Nome completo", fullWord)
}
