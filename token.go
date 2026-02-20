package main

import "fmt"

type TokenKind int

const (
	TOKEN_UNKNOWN TokenKind = iota
	TOKEN_EOF
	TOKEN_LET
	TOKEN_ASSIGN
	TOKEN_EQULES
	TOKEN_NUMBER
	TOKEN_IDENT
	TOKEN_NEWLINE
	TOKEN_PLUS
	TOKEN_INC
	TOKEN_MINUS
	TOKEN_DEC
	TOKEN_TIEMS
	TOKEN_DIV
	TOKEN_COMMENT
	TOKEN_OPENPARN
	TOKEN_CLOSEPARN
)

var tokenToString = map[TokenKind]string{
	TOKEN_UNKNOWN:   "TOKEN_UNKNOWN",
	TOKEN_EOF:       "TOKEN_EOF",
	TOKEN_LET:       "TOKEN_LET",
	TOKEN_ASSIGN:    "TOKEN_ASSIGN",
	TOKEN_EQULES:    "TOKEN_EQULES",
	TOKEN_NUMBER:    "TOKEN_NUMBER",
	TOKEN_IDENT:     "TOKEN_IDENT",
	TOKEN_NEWLINE:   "TOKEN_NEWLINE",
	TOKEN_PLUS:      "TOKEN_PLUS",
	TOKEN_INC:       "TOKEN_INC",
	TOKEN_MINUS:     "TOKEN_MINUS",
	TOKEN_DEC:       "TOKEN_DEC",
	TOKEN_TIEMS:     "TOKEN_TIEMS",
	TOKEN_DIV:       "TOKEN_DIV",
	TOKEN_COMMENT:   "TOKEN_COMMENT",
	TOKEN_OPENPARN:  "TOKEN_OPENPARN",
	TOKEN_CLOSEPARN: "TOKEN_CLOSEPARN",
}

type Token struct {
	kind TokenKind
	lit  string
}

var keywords = map[string]TokenKind{
	"let": TOKEN_LET,
}

func (t Token) toString() string {
	tokStr, exist := tokenToString[t.kind]

	if !exist {
		panic("No string repr in table")
	}

	if t.kind == TOKEN_NEWLINE {
		return fmt.Sprintf("%v", tokStr)
	} else {
		return fmt.Sprintf("%v %v", tokStr, t.lit)
	}

}
