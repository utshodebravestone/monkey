package token

type TokenKind string

const (
	// SINGLE CHARACTER OPERATORS
	PLUS     TokenKind = "PLUS"
	MINUS    TokenKind = "MINUS"
	ASTERISK TokenKind = "ASTERISK"
	SLASH    TokenKind = "SLASH"
	MODULO   TokenKind = "MODULO"
	HAT      TokenKind = "HAT"
	EQUAL    TokenKind = "EQUAL"
	GREATER  TokenKind = "GREATER"
	LESSER   TokenKind = "LESSER"

	// DOUBLE CHARACTER OPERATORS
	EQUAL_EQUAL   TokenKind = "EQUAL_EQUAL"
	BANG_EQUAL    TokenKind = "BANG_EQUAL"
	GREATER_EQUAL TokenKind = "GREATER_EQUAL"
	LESSER_EQUAL  TokenKind = "LESSER_EQUAL"

	// SEPARATORS
	OPEN_PAREN    TokenKind = "OPEN_PAREN"
	CLOSE_PAREN   TokenKind = "CLOSE_PAREN"
	OPEN_BRACE    TokenKind = "OPEN_BRACE"
	CLOSE_BRACE   TokenKind = "CLOSE_BRACE"
	OPEN_BRACKET  TokenKind = "OPEN_BRACKET"
	CLOSE_BRACKET TokenKind = "CLOSE_BRACKET"
	COLON         TokenKind = "COLON"
	COMMA         TokenKind = "COMMA"
	SEMICOLON     TokenKind = "SEMICOLON"

	// LITERALS
	NUMBER     TokenKind = "NUMBER"
	STRING     TokenKind = "STRING"
	IDENTIFIER TokenKind = "IDENTIFIER"

	// KEYWORDS
	TRUE  TokenKind = "TRUE"
	FALSE TokenKind = "FALSE"
	LET   TokenKind = "LET"
	FUN   TokenKind = "FUN"
	RET   TokenKind = "RET"
	IF    TokenKind = "IF"
	ELSE  TokenKind = "ELSE"
	NOT   TokenKind = "NOT"
	AND   TokenKind = "AND"
	OR    TokenKind = "OR"

	// SPECIAL TOKENS
	ILLEGAL TokenKind = "ILLEGAL"
	EOF     TokenKind = "EOF"
)

var keywords = map[string]TokenKind{
	"true":  TRUE,
	"false": FALSE,
	"let":   LET,
	"fun":   FUN,
	"ret":   RET,
	"if":    IF,
	"else":  ELSE,
	"not":   NOT,
	"and":   AND,
	"or":    OR,
}

func LookupKeyword(lexeme string) TokenKind {
	if kind, ok := keywords[lexeme]; ok {
		return kind
	} else {
		return IDENTIFIER
	}
}
