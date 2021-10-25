package errors

import (
	"fmt"
	"strings"
)

type MissingParam struct {
	Param []string
}

func (e MissingParam) Error() string {
	if len(e.Param) > 1 {
		return fmt.Sprintf("missing parameters: %s", strings.Join(e.Param, ", "))
	}

	return fmt.Sprintf("missing parameter: %s", e.Param[0])
}
