package movies

type Validator struct{}

func (v Validator) Validate(movie *Movie) error {
	if movie.Title == "" {
		return MissingTitleError
	}

	return nil
}
