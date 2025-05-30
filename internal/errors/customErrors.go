package customErrors

import "fmt"

var (
	ErrNotFound = fmt.Errorf("not found")
)
