package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	debug := true
	reader := bufio.NewReader(os.Stdin)

	for {
		var file_text string
		if debug {
			file_text = "CH[0] + DynoSpeed * 2.5 - (3 / 4) 4"
		} else {
			var err error
			fmt.Print(">> ")
			file_text, err = reader.ReadString('\n')
			if err != nil {
				break
			}
		}

		fmt.Printf("Input: %v", file_text)
		fmt.Printf("\n--- TOKENS ---\n")
		NewTokenizer(string(file_text)).dumpTokens()

		fmt.Printf("\n--- AST ---\n")
		NewParser(NewTokenizer(string(file_text))).parse().printTree()
		fmt.Printf("\n--- OUTPUT ---\n")
		interpreter := NewInterperter(NewParser(NewTokenizer(string(file_text))))
		fmt.Printf("%v\n", interpreter.interpret())
		fmt.Printf("\n--- COMPILED ---\n")
		codeGen := NewCodeGen(NewParser(NewTokenizer(string(file_text))))
		fmt.Printf("%v\n", codeGen.generate())

		if debug {
			break
		}
	}
}
