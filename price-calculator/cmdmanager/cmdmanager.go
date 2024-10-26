package cmdmanager

import "fmt"

type CMDManager struct {
}

func New() *CMDManager {
	return &CMDManager{}
}

func (cmd *CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm with 'q' and press Enter")

	var lines []string
	for {
		var line string
		fmt.Print("Price: ")
		fmt.Scanln(&line)

		if line == "q" {
			break
		}
		lines = append(lines, line)
	}

	return lines, nil
}

func (cmd *CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}
