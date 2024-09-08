package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"struct_practice/note"
)

func main() {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display()
	note.Save()
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}
