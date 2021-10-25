package errors

import (
	"fmt"
	"strings"
)

type InvalidParam struct {
	Param []string
}

func (e InvalidParam) Error() string {
	if len(e.Param) > 1 {
		return fmt.Sprintf("invalid value for parameters %s", strings.Join(e.Param, ", "))
	}

	return fmt.Sprintf("invalid value for parameter %s", e.Param[0])
}
