package autors

type Autor struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
}
