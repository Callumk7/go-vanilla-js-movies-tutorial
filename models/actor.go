package models

// A nice note: a pointer in go * can be thought of as something that can make a value nullable

type Actor struct {
	ID        int
	FirstName string
	LastName  string
	ImageURL  *string
}
