package token

type TokenKind string

const (
	// SPECIAL TOKENS
	ILLEGAL TokenKind = "ILLEGAL"
	EOF     TokenKind = "EOF"

	// SINGLE CHARACTER TOKENS

	// OPERATORS
	PLUS     TokenKind = "PLUS"
	MINUS    TokenKind = "MINUS"
	ASTERISK TokenKind = "ASTERISK"
	SLASH    TokenKind = "SLASH"
	MODULO   TokenKind = "MODULO"
	HAT      TokenKind = "HAT"
	BANG     TokenKind = "BANG"
	EQUAL    TokenKind = "EQUAL"

	// SEPARATORS
	OPEN_PAREN    TokenKind = "OPEN_PAREN"
	CLOSE_PAREN   TokenKind = "CLOSE_PAREN"
	OPEN_BRACE    TokenKind = "OPEN_BRACE"
	CLOSE_BRACE   TokenKind = "CLOSE_BRACE"
	OPEN_BRACKET  TokenKind = "OPEN_BRACKET"
	CLOSE_BRACKET TokenKind = "CLOSE_BRACKET"
	COMMA         TokenKind = "COMMA"
	SEMICOLON     TokenKind = "SEMICOLON"
)
