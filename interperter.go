package main

import (
	"fmt"
	"strconv"
)

type Interpreter struct {
	parser *Parser
}

var chvals = map[string]float64{
	"CH1": 5,
	"CH2": 10,
	"a":   1,
	"b":   2,
	"c":   3,
}

func NewInterperter(parser *Parser) *Interpreter {
	return &Interpreter{parser: parser}
}

func (i *Interpreter) visit_BinOp(n *BinOpNode) float64 {

	var r float64

	switch n.tok.kind {
	case TOKEN_PLUS:
		r = i.visit(n.left) + i.visit(n.right)
		break
	case TOKEN_MINUS:
		r = i.visit(n.left) - i.visit(n.right)
		break
	case TOKEN_DIV:
		r = i.visit(n.left) / i.visit(n.right)
		break
	case TOKEN_TIEMS:
		r = i.visit(n.left) * i.visit(n.right)
		break
	default:
		panic("Unrechable")
	}

	return r
}

func (i *Interpreter) visit_Num(n *BinOpNode) float64 {

	val, err := strconv.ParseFloat(n.tok.lit, 64)

	if err != nil {
		panic(err)
	}

	return val
}

func (i *Interpreter) visit_Var(n *BinOpNode) float64 {

	val, have := chvals[n.tok.lit]

	if !have {
		panic(fmt.Sprintf("No channel %s", n.tok.lit))
	}

	return val
}

func (i *Interpreter) visit(n *BinOpNode) float64 {

	switch n.tok.kind {
	case TOKEN_NUMBER:
		return i.visit_Num(n)
	case TOKEN_IDENT:
		return i.visit_Var(n)
	case TOKEN_DIV, TOKEN_TIEMS, TOKEN_PLUS, TOKEN_MINUS:
		return i.visit_BinOp(n)
	}

	panic("UNRECHABLE")
}

func (i *Interpreter) interpret() float64 {
	tree := i.parser.parse()
	return i.visit(tree)
}
