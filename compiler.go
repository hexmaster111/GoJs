package main

import (
	"strings"
)

type CodeGen struct {
	ast *BinOpNode
}

func (cg *CodeGen) visit(n *BinOpNode) []string {
	switch n.tok.kind {
	case TOKEN_NUMBER, TOKEN_IDENT:
		return []string{"= " + n.tok.lit}
	case TOKEN_TIEMS:
		return cg.binop(n, "*")
	case TOKEN_DIV:
		return cg.binop(n, "/")
	case TOKEN_MINUS:
		return cg.binop(n, "-")
	case TOKEN_PLUS:
		return cg.binop(n, "+")
	}
	panic("unknown token in visit")
}

func (cg *CodeGen) binop(n *BinOpNode, op string) []string {
	left := cg.visit(n.left)
	right := cg.visit(n.right)

	if len(right) > 1 {
		left[0] = op + left[0]
		return append(right, left...)
	}
	right[0] = op + right[0]
	return append(left, right...)
}

func (cg *CodeGen) generate() string {
	return strings.Join(cg.visit(cg.ast), "\n")
}

func NewCodeGen(p *Parser) *CodeGen {
	return &CodeGen{ast: p.parse()}
}
