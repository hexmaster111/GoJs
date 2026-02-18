package main

import (
	"unicode"
)

type Scanner struct {
	input        string
	position     int
	readPosition int
	char         byte
}

func NewScanner(txt string) *Scanner {
	s := &Scanner{input: txt}
	s.readChar()
	return s
}

func (s *Scanner) readChar() {
	if s.readPosition >= len(s.input) {
		s.char = 0
	} else {
		s.char = s.input[s.readPosition]
	}
	s.position = s.readPosition
	s.readPosition++
}

func (s *Scanner) skipWhiteSpace() {
	for unicode.IsSpace(rune(s.char)) && s.char != '\n' {
		s.readChar()
	}
}

func isLetter(c byte) bool { return unicode.IsLetter(rune(c)) }
func isNumber(c byte) bool { return unicode.IsNumber(rune(c)) }

func (s *Scanner) readIdent() string {
	var buf []byte

	for isLetter(s.char) {
		buf = append(buf, s.char)
		s.readChar()
	}
	return string(buf)
}

func (s *Scanner) readNumber() string {
	var buf []byte

	for isNumber(s.char) {
		buf = append(buf, s.char)
		s.readChar()
	}
	return string(buf)
}

func LookupIdent(st string) TokenKind {
	if st == "let" {
		return TOKEN_LET
	} else {
		return TOKEN_IDENT
	}
}

func (s *Scanner) NextToken() Token {
	var token Token
	s.skipWhiteSpace()

	switch s.char {
	case '\n':
		token = Token{kind: TOKEN_NEWLINE, lit: string(s.char)}
	case '=':
		token = Token{kind: TOKEN_EQ, lit: string(s.char)}
	case 0:
		token = Token{kind: TOKEN_EOF, lit: ""}

	default:
		if isLetter(s.char) {
			lit := s.readIdent()
			tokKind := LookupIdent(lit)
			token = Token{kind: tokKind, lit: lit}
			return token
		} else if isNumber(s.char) {
			lit := s.readNumber()
			token = Token{kind: TOKEN_NUMBER, lit: lit}
			return token
		} else {
			token = Token{kind: TOKEN_ILLGLE, lit: string(s.char)}
		}
	}

	s.readChar()
	return token
}
