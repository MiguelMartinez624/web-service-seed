package movies

import "errors"

var (
	MissingTitleError error = simpleErr("Password cant be empty")
)

// simpleErr create simple error with flat msg
func simpleErr(msg string) error {
	return errors.New(msg)
}
