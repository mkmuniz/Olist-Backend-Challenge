package book

type Book struct {
	Name        string `json:"name"`
	Edition     string `json:"edition"`
	Publication string `json:"publicationYear"`
	Authors     string `json:"authors"`
}
