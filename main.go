package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		file_text, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		tokenizer := NewTokenizer(string(file_text))
		parser := NewParser(tokenizer)
		interpreter := NewInterperter(parser)
		fmt.Printf("\nres: %v\n", interpreter.interpret())
	}
}
