package errors

import "fmt"

type NotFound struct {
	Code   int
	Entity string
}

func (n *NotFound) Error() string {
	return fmt.Sprintf("%s was not found", n.Entity)
}
