package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Question make question to use and retrieve answer
func Question(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", question)
	answer, _ := reader.ReadString('\n')

	return strings.TrimSuffix(answer, "\n")
}
