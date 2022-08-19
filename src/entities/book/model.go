package book

type Book struct {
	id              string `json:"id"`
	name            string `json:"name"`
	edition         string `json:"edition"`
	publicationYear string `json:"publicationYear"`
	authors         string `json:"authors"`
}