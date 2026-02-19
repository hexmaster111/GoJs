package main

import (
	"fmt"
	"unicode"
)

type Tokenizer struct {
	input        string
	position     int
	readPosition int
	char         byte
	peek         byte
}

func NewTokenizer(txt string) *Tokenizer {
	s := &Tokenizer{input: txt}
	s.readChar()
	return s
}

func (s *Tokenizer) readChar() {
	if s.readPosition >= len(s.input) {
		s.char = 0
	} else {
		s.char = s.input[s.readPosition]
	}

	if s.readPosition+1 >= len(s.input) {
		s.peek = 0
	} else {
		s.peek = s.input[s.readPosition+1]
	}

	s.position = s.readPosition
	s.readPosition++
}

func (s *Tokenizer) skipWhiteSpace() {
	for unicode.IsSpace(rune(s.char)) && s.char != '\n' {
		s.readChar()
	}
}

func isLetter(c byte) bool { return unicode.IsLetter(rune(c)) }
func isNumber(c byte) bool { return unicode.IsNumber(rune(c)) }

var symbols = map[byte]bool{
	'+': true,
	'-': true,
	'/': true,
	'*': true,
	')': true,
	'(': true,
}

func isSpecial(c byte) bool {
	_, exist := symbols[c]
	return exist
}

func (s *Tokenizer) readIdent() string {
	var buf []byte

	for !unicode.IsSpace(rune(s.char)) && !isSpecial(s.char) {
		buf = append(buf, s.char)
		s.readChar()
	}
	return string(buf)
}

func (s *Tokenizer) readNumber() string {
	var buf []byte
	gotOneDot := false

	for isNumber(s.char) || (!gotOneDot && s.char == '.') {
		buf = append(buf, s.char)

		if s.char == '.' {
			gotOneDot = true
		}

		s.readChar()
	}
	return string(buf)
}

func LookupIdent(st string) TokenKind {
	kw, exist := keywords[st]

	if exist {
		return kw
	} else {
		return TOKEN_IDENT
	}
}

func (s *Tokenizer) NextToken() Token {
	var token Token
	s.skipWhiteSpace()
	if s.char == '\n' {
		token = Token{kind: TOKEN_NEWLINE, lit: string(s.char)}
	} else if s.char == '=' && s.peek == '=' {
		token = Token{kind: TOKEN_EQULES, lit: "=="}
		s.readChar()
	} else if s.char == '+' && s.peek == '+' {
		token = Token{kind: TOKEN_INC, lit: "++"}
		s.readChar()
	} else if s.char == '-' && s.peek == '-' {
		token = Token{kind: TOKEN_DEC, lit: "--"}
		s.readChar()
	} else if s.char == '/' && s.peek == '/' {
		token = Token{kind: TOKEN_COMMENT, lit: "//"}
		s.readChar()
	} else if s.char == '*' {
		token = Token{kind: TOKEN_TIEMS, lit: string(s.char)}
	} else if s.char == '+' {
		token = Token{kind: TOKEN_PLUS, lit: string(s.char)}
	} else if s.char == '-' {
		token = Token{kind: TOKEN_MINUS, lit: string(s.char)}
	} else if s.char == '/' {
		token = Token{kind: TOKEN_DIV, lit: string(s.char)}
	} else if s.char == '(' {
		token = Token{kind: TOKEN_OPENPARN, lit: string(s.char)}
	} else if s.char == ')' {
		token = Token{kind: TOKEN_CLOSEPARN, lit: string(s.char)}
	} else if s.char == '=' {
		token = Token{kind: TOKEN_ASSIGN, lit: string(s.char)}
	} else if s.char == 0 {
		token = Token{kind: TOKEN_EOF, lit: ""}
	} else if isLetter(s.char) {
		lit := s.readIdent()
		tokKind := LookupIdent(lit)
		token = Token{kind: tokKind, lit: lit}
		return token
	} else if isNumber(s.char) {
		lit := s.readNumber()
		token = Token{kind: TOKEN_NUMBER, lit: lit}
		return token
	} else {
		token = Token{kind: TOKEN_UNKNOWN, lit: string(s.char)}
		msg := fmt.Sprintf("UNKNOWN TOKEN! %s", token.lit)
		panic(msg)
	}

	s.readChar()
	return token
}

func (tokenizer *Tokenizer) dumpTokens() {
	for {
		tok := tokenizer.NextToken()

		tokStr, exist := tokenToString[tok.kind]

		if !exist {
			panic("")
		}

		if tok.kind == TOKEN_NEWLINE {
			fmt.Printf("%v\n", tokStr)
		} else {
			fmt.Printf("%v %v\n", tokStr, tok.lit)
		}

		if tok.kind == TOKEN_EOF {
			break
		}
	}
}
