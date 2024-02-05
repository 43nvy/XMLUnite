package ui

import "fmt"

type ConsoleUI struct {
}

func (c *ConsoleUI) InputData(data *string) {
	fmt.Scan(data)
}

func (c *ConsoleUI) OutputData(data ...string) {
	for _, str := range data {
		fmt.Println(str)
	}
}

func New() ConsoleUI {
	return ConsoleUI{}
}
