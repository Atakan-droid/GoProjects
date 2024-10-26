package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

// Constructor
func New(inputFilePath, outputFilePath string) FileManager {
	return FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("error opening file " + fm.InputFilePath)
	}
	defer file.Close()

	// Reader
	reader := bufio.NewScanner(file)

	// Read line by line
	var lines []string
	for reader.Scan() {
		lines = append(lines, reader.Text())
	}

	err = reader.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("error reading file")
	}

	file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("error creating file")
	}
	defer file.Close()

	// Write JSON
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("error to convert data to JSON")
	}

	file.Close()
	return nil
}
