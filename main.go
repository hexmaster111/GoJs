package main

import (
	"fmt"
	"os"
)

type LocalDeclStmt struct{}
type ExprStmt struct{}
type Stmt interface{}
type Block struct {
	statements []Stmt
}

func main() {

	file_text, err := os.ReadFile("test.js")

	if err != nil {
		panic(err)
	}

	scanner := NewScanner(string(file_text))

	for {
		tok := scanner.NextToken()
		fmt.Printf("%v %v\n", tok.kind, tok.lit)
		if tok.kind == TOKEN_EOF {
			break
		}
	}

	fmt.Printf("%v\n", file_text)
}
