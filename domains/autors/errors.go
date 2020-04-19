package autors

import "errors"

var (
	MissingNameError  error = simpleErr("Missing Autor's Name")
	AutorNoExistError       = simpleErr("Autor doesn't exist")
	InvalidIDError          = simpleErr("Invalid Autor's ID")
)

// simpleErr create simple error with flat msg
func simpleErr(msg string) error {
	return errors.New(msg)
}
