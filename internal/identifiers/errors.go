package identifiers

import "fmt"

var (
	ErrInvalidIdentifier   = &InvalidIdentifierError{}
	ErrIdentifierGenerator = &IdentifierGeneratorError{}
)

type InvalidIdentifierError struct {
	wrapped error
}

func (r *InvalidIdentifierError) Error() string {
	if r.wrapped == nil {
		return "invalid identifier"
	}

	return fmt.Sprintf("invalid identifier: %v", r.wrapped.Error())
}

func (r *InvalidIdentifierError) Unwrap() error {
	return r.wrapped
}

type IdentifierGeneratorError struct {
	wrapped error
}

func (e *IdentifierGeneratorError) Error() string {
	if e.wrapped == nil {
		return "failed to generate identifier"
	}

	return fmt.Sprintf("failed to generate identifier: %v", e.wrapped.Error())
}

func (e *IdentifierGeneratorError) Unwrap() error {
	return e.wrapped
}
