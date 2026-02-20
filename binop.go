package main

import (
	"fmt"
	"strings"
)

type BinOpNode struct {
	tok         Token      // + / - * Number
	left, right *BinOpNode // null for number/var
}

func (n *BinOpNode) _printTree(depth int) {

	b := strings.Builder{}

	for i := 0; i < depth*2; i++ {
		b.WriteByte(' ')
	}

	fmt.Printf("%v %v\n", b.String(), n.tok.toString())

	switch n.tok.kind {
	case TOKEN_MINUS:
		n.left._printTree(depth + 1)
		n.right._printTree(depth + 1)
	case TOKEN_DIV:
		n.left._printTree(depth + 1)
		n.right._printTree(depth + 1)
	case TOKEN_PLUS:
		n.left._printTree(depth + 1)
		n.right._printTree(depth + 1)
	case TOKEN_TIEMS:
		n.left._printTree(depth + 1)
		n.right._printTree(depth + 1)
	case TOKEN_IDENT:
	case TOKEN_NUMBER:

	}

}

func (n *BinOpNode) printTree() {
	n._printTree(0)
}
