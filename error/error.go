package error

import "monkey/text"

type Error interface {
	Message() string
	Span() text.Span

	Info() string
}
