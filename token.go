package main

type TokenKind string

const (
	TOKEN_ILLGLE  = "ILLEGLE"
	TOKEN_EOF     = "EOF"
	TOKEN_LET     = "let"
	TOKEN_EQ      = "="
	TOKEN_NUMBER  = "NUMBER"
	TOKEN_IDENT   = "IDENT"
	TOKEN_NEWLINE = "NEWLINE"
)

type Token struct {
	kind TokenKind
	lit  string
}

var keywords = map[string]TokenKind{
	"let": TOKEN_LET,
	"=":   TOKEN_EQ,
}
