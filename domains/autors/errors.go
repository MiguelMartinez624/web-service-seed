package autors

import "errors"

var (
	MissingNameError error = simpleErr("Missing Autor's Name")
)

// simpleErr create simple error with flat msg
func simpleErr(msg string) error {
	return errors.New(msg)
}
