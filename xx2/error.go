package xx2

import (
	"fmt"
)

func IfError(condition bool, i interface{}) error {
	if condition {
		return fmt.Errorf("%s", i)
	}
	return nil
}
