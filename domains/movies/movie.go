package movies

type Movie struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Title       string `json:"title" bson:"title,omitempty"`
	Description string `json:"desciption" bson:"Description,omitempty"`
	AutorID     string `json:"autor_id" bson:"autor_id,omitempty"`
}
