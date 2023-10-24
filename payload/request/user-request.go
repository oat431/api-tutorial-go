package request

type UserRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Birthday  string `json:"birth_date"`
}
