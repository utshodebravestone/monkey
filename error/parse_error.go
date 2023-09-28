package error

import (
	"fmt"

	"monkey/text"
)

type ParseError struct {
	message string
	span    text.Span
}

func (pe *ParseError) Message() string {
	return pe.message
}

func (pe *ParseError) Span() text.Span {
	return pe.span
}

func (pe *ParseError) Info() string {
	return fmt.Sprintf("parse error: %s in %s", pe.message, pe.span.ToString())
}
