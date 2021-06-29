package handle

import (
	"fmt"
)

type Error struct {
	Code   int   `json:"code"`
	RawErr error `json:"error"`
}

func (err Error) Error() string {
	return fmt.Sprintf("(%d) %s", err.Code, err.RawErr)
}

var (
	ErrExists = Error{40301, fmt.Errorf("duplicate names")}
)
