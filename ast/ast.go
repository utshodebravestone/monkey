package ast

import (
	"fmt"

	"monkey/diagnostics"
)

type Node interface {
	Span() diagnostics.Span
	ToString() string
}

type Program struct {
	Statements []Statement
}

func (p *Program) Span() diagnostics.Span {
	stmtLen := len(p.Statements)

	if stmtLen == 0 {
		return diagnostics.NewSpan(0, 0)
	} else {
		start := 0
		end := stmtLen - 1
		return diagnostics.ExtentSpan(p.Statements[start].Span(), p.Statements[end].Span())

	}
}

func (p *Program) ToString() string {
	s := ""
	for _, stmt := range p.Statements {
		s += fmt.Sprintf("%s\n", stmt.ToString())
	}
	return s
}
