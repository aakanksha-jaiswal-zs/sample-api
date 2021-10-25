package errors

import "fmt"

type DB struct {
	Err error
}

func (e DB) Error() string {
	return fmt.Sprintf("database error: %s", e.Err.Error())
}
