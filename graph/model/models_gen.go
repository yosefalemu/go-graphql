// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Movie struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

type Mutation struct {
}

type NewMovie struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type Query struct {
}
