package main

import (
	"bufio"
	"fmt"
	"strings"
)

func GetInputFromUser(prompt string, reader *bufio.Reader) string {
	fmt.Printf(prompt + " -> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error", err)
		return GetInputFromUser(prompt, reader)
	}
	return strings.Trim(input, " \r\n")
}
