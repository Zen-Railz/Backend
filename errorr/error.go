package errorr

import (
	"fmt"
)

func (e *entity) Error() string {
	return fmt.Sprintf("(%s) %s", e.code, e.message)
}
