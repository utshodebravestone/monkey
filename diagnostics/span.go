package diagnostics

import "fmt"

type Span struct {
	Start, End uint
}

func NewSpan(start, end uint) Span {
	return Span{
		Start: start,
		End:   end,
	}
}

func (s *Span) ToString() string {
	return fmt.Sprintf("%d:%d", s.Start, s.End)
}

func ExtentSpan(a, b Span) Span {
	return NewSpan(a.Start, b.End)
}
