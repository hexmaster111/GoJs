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

	defer func() {
		ts, _ := tokenToString[n.tok.kind]
		fmt.Printf("Visited: %v\n", ts)
	}()

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

type BinOpNode struct {
	tok         Token      // + / - * Number
	left, right *BinOpNode // null for number/var
}

type Parser struct {
	tokenizer    *Tokenizer
	currentToken Token
}

func NewParser(tokenizer *Tokenizer) *Parser {
	var ret Parser
	ret.tokenizer = tokenizer
	ret.currentToken = ret.tokenizer.NextToken()
	return &ret
}

func ParseError(msg string) { panic("Parse Error: " + msg) }

func (p *Parser) consume(expectedKind TokenKind) {
	if p.currentToken.kind != expectedKind {
		ParseError(fmt.Sprintf("Expected %v, got %v", expectedKind, p.currentToken.kind))
	}

	p.currentToken = p.tokenizer.NextToken()
}

func (p *Parser) factor() *BinOpNode {
	token := p.currentToken

	switch token.kind {

	case TOKEN_IDENT:
		p.consume(TOKEN_IDENT)
		return &BinOpNode{left: nil, right: nil, tok: token}
	case TOKEN_NUMBER:
		p.consume(TOKEN_NUMBER)
		return &BinOpNode{left: nil, right: nil, tok: token}
	case TOKEN_OPENPARN:
		p.consume(TOKEN_OPENPARN)
		node := p.expr()
		p.consume(TOKEN_CLOSEPARN)
		return node
	}

	str, had := tokenToString[token.kind]

	if had {
		panic(fmt.Sprintf("Unexpected Token: %v", str))
	} else {
		panic(fmt.Sprintf("Unknown Token: '%v'", token.lit))
	}
}

func (p *Parser) term() *BinOpNode {
	node := p.factor()
	for p.currentToken.kind == TOKEN_TIEMS || p.currentToken.kind == TOKEN_DIV {
		token := p.currentToken
		switch token.kind {
		case TOKEN_TIEMS:
			p.consume(TOKEN_TIEMS)
		case TOKEN_DIV:
			p.consume(TOKEN_DIV)
		}
		node = &BinOpNode{left: node, tok: token, right: p.factor()}
	}
	return node
}

func (p *Parser) expr() *BinOpNode {
	node := p.term()
	for p.currentToken.kind == TOKEN_PLUS || p.currentToken.kind == TOKEN_MINUS {
		token := p.currentToken
		switch token.kind {
		case TOKEN_PLUS:
			p.consume(TOKEN_PLUS)
		case TOKEN_MINUS:
			p.consume(TOKEN_MINUS)
		}
		node = &BinOpNode{left: node, tok: token, right: p.term()}
	}
	return node
}

func (p *Parser) parse() *BinOpNode { return p.expr() }

/*
TOKEN_NUMBER 1
TOKEN_PLUS +
TOKEN_NUMBER 2
TOKEN_NEWLINE
TOKEN_EOF
*/
