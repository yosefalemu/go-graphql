// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Actor struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	ReleaseDate string `json:"releaseDate"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type Mutation struct {
}

type NewActor struct {
	Name string `json:"name"`
}

type NewMovie struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	ReleaseDate string `json:"releaseDate"`
}

type Query struct {
}

type UpdateActor struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type UpdateMovie struct {
	ID          string  `json:"id"`
	Title       *string `json:"title,omitempty"`
	URL         *string `json:"url,omitempty"`
	ReleaseDate *string `json:"releaseDate,omitempty"`
}
