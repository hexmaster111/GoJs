package main

import (
	"fmt"
)

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
