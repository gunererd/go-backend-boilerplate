package utils

import "fmt"

// FError is a generic error type.
type FError struct {
	Code      int
	ErrorCode string
	Message   string
}

// Error returns a string representation of FError instance.
func (b *FError) Error() string {
	return fmt.Sprintf(
		"{\n\t\"code\": \"%d\",\n\t\"message\": \"%s\",\n\t\"message\": \"%s\",\"\n}",
		b.Code, b.Message, b.ErrorCode,
	)
}
