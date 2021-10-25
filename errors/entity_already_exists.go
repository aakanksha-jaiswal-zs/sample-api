package errors

import "fmt"

type EntityAlreadyExists struct {
	Entity string
}

func (e EntityAlreadyExists) Error() string {
	return fmt.Sprintf("entity %s already exists", e.Entity)
}
