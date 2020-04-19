package autors

type Validator struct{}

func (v Validator) Validate(autor *Autor) error {
	if autor.Name == "" {
		return MissingNameError
	}

	return nil
}
