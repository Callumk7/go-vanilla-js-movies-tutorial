package models

// A nice note: a pointer in go * can be thought of as something that can make a value nullable

type Actor struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	ImageURL  *string `json:"image_url,omitempty"`
}
