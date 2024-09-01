package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "data.txt"

func writeToFile(data string) error {
	//write to file
	err := os.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		return errors.New("error writing to file")
	}
	return nil
}

func readFromFile() (string, error) {
	res, err := os.ReadFile(fileName)
	if err != nil {
		return "", errors.New("error reading from file")
	}
	return string(res), nil
}

//Task: Write a program that reads the account balance from a file and prints it to the console
//The account balance should be stored in a file called data.txt
//The account balance should be a float64

func main() {
	var accountBalance float64 = 1000.0
	fmt.Println("Welcome to the bank")
	err := writeToFile(fmt.Sprint(accountBalance))

	//read from file
	data, err := readFromFile()

	dataFloat, err := strconv.ParseFloat(data, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(dataFloat)

}
