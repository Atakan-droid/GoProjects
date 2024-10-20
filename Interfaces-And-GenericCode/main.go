package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputter interface {
	saver
	displayer
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func displayData(data displayer) {
	data.Display()
}

func outputData(data outputter) {
	saveData(data)
	displayData(data)
}

// any value allowed type
func printSomething(value interface{}) {

	//TODO: special switch statement for type
	switch value.(type) {
	case string:
		fmt.Println("This is a string.", value)
	case int:
		fmt.Println("This is an integer.", value)
	default:
		fmt.Println("This is something else.", value)
	}

}

func getType(value interface{}) {
	//TODO: Extract the type of the value
	typedValue, ok := value.(int)
	println(typedValue, ok)
}

// generic function with type constraints
func add[T int | float64 | string](a, b T) T {
	return a + b
}

func main() {

	result := add(1, 2)
	fmt.Println(result)

	printSomething(42)
	printSomething("Hello, world!")
	printSomething(3.14)
	getType(42)
	getType("Hello, world!")

	title, content := getNoteData()

	userNote, err := note.New(title, content)
	userTodo, err := todo.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userNote)
	outputData(userTodo)

}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
