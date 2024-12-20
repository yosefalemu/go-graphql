package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/yosefalemu/go-graphql.git/graph/model"
	"gorm.io/gorm"
)

// CreateMovie is the resolver for the createMovie field.
func (r *mutationResolver) CreateMovie(ctx context.Context, input model.NewMovie) (*model.Movie, error) {
	movieID := uuid.New()
	newMovie := model.Movie{
		ID:    movieID.String(),
		Title: input.Title,
		URL:   input.URL,
		ReleaseDate: func() *string {
			if input.ReleaseDate != nil {
				return input.ReleaseDate
			}
			now := time.Now().Format(time.RFC3339)
			return &now
		}(),
		CreatedAt: func(s string) *string { return &s }(time.Now().Format(time.RFC3339)),
		UpdatedAt: func(s string) *string { return &s }(time.Now().Format(time.RFC3339)),
	}

	result := r.DB.Create(&newMovie)

	if result.Error != nil {
		return nil, result.Error
	}

	return &newMovie, nil
}

// CreateActor is the resolver for the createActor field.
func (r *mutationResolver) CreateActor(ctx context.Context, input model.NewActor) (*model.Actor, error) {
	actorID := uuid.New()
	newActor := model.Actor{
		ID:        actorID.String(),
		Name:      input.Name,
		CreatedAt: func(s string) *string { return &s }(time.Now().Format(time.RFC3339)),
		UpdatedAt: func(s string) *string { return &s }(time.Now().Format(time.RFC3339)),
	}

	result := r.DB.Create(&newActor)

	if result.Error != nil {
		return nil, result.Error
	}

	return &newActor, nil
}

// UpdateMovie is the resolver for the updateMovie field.
func (r *mutationResolver) UpdateMovie(ctx context.Context, input model.UpdateMovie) (*model.Movie, error) {
	if _, err := uuid.Parse(input.ID); err != nil {
		return nil, errors.New("invalid uuid")
	}
	var movie model.Movie
	result := r.DB.First(&movie, "id =?", input.ID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			errFound := fmt.Errorf("movie with id %s not found", input.ID)
			return nil, errFound
		}
		return nil, result.Error
	}
	if input.Title != nil {
		movie.Title = *input.Title
	}
	if input.URL != nil {
		movie.URL = *input.URL
	}
	if input.ReleaseDate != nil {
		movie.ReleaseDate = input.ReleaseDate
	}
	movie.UpdatedAt = func(s string) *string { return &s }(time.Now().Format(time.RFC3339))
	result = r.DB.Save(&movie)
	if result.Error != nil {
		return nil, result.Error
	}
	return &movie, nil
}

// UpdateActor is the resolver for the updateActor field.
func (r *mutationResolver) UpdateActor(ctx context.Context, input model.UpdateActor) (*model.Actor, error) {
	if _, err := uuid.Parse(input.ID); err != nil {
		return nil, errors.New("invalid uuid")
	}
	var actor model.Actor
	result := r.DB.First(&actor, "id = ?", input.ID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			errFound := fmt.Errorf("actor with id %s not found", input.ID)
			return nil, errFound
		}
		return nil, result.Error
	}
	if input.Name != nil {
		actor.Name = *input.Name
	}
	actor.UpdatedAt = func(s string) *string { return &s }(time.Now().Format(time.RFC3339))
	result = r.DB.Save(&actor)
	if result.Error != nil {
		return nil, result.Error
	}
	return &actor, nil
}

// DeleteMovie is the resolver for the deleteMovie field.
func (r *mutationResolver) DeleteMovie(ctx context.Context, id string) (*model.Movie, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.New("invalid uuid")
	}
	var movie model.Movie
	result := r.DB.First(&movie, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			errFound := fmt.Errorf("movie with id %s not found", id)
			return nil, errFound
		}
		return nil, result.Error
	}
	result = r.DB.Delete(&movie)
	if result.Error != nil {
		return nil, result.Error
	}
	return &movie, nil
}

// DeleteActor is the resolver for the deleteActor field.
func (r *mutationResolver) DeleteActor(ctx context.Context, id string) (*model.Actor, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.New("invalid uuid")
	}
	var actor model.Actor
	result := r.DB.First(&actor, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			errFound := fmt.Errorf("actor with id %s not found", id)
			return nil, errFound
		}
		return nil, result.Error
	}
	result = r.DB.Delete(&actor)
	if result.Error != nil {
		return nil, result.Error
	}
	return &actor, nil
}

// Movies is the resolver for the movies field.
func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	var movies []*model.Movie
	result := r.DB.Find(&movies)
	if result.Error != nil {
		return nil, result.Error
	}
	return movies, nil
}

// Movie is the resolver for the movie field.
func (r *queryResolver) Movie(ctx context.Context, id string) (*model.Movie, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.New("invalid uuid")
	}
	var movie model.Movie
	result := r.DB.First(&movie, "id =?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			errFound := fmt.Errorf("movie with id %s not found", id)
			return nil, errFound
		}
		return nil, result.Error
	}
	return &movie, nil
}

// Actors is the resolver for the actors field.
func (r *queryResolver) Actors(ctx context.Context) ([]*model.Actor, error) {
	var actors []*model.Actor

	result := r.DB.Find(&actors)

	if result.Error != nil {
		return nil, result.Error
	}

	return actors, nil
}

// Actor is the resolver for the actor field.
func (r *queryResolver) Actor(ctx context.Context, id string) (*model.Actor, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.New("inavlid uuid")
	}
	var actor model.Actor
	result := r.DB.First(&actor, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			errFound := fmt.Errorf("actor with id %s not found", id)
			return nil, errFound
		}
		return nil, result.Error
	}
	return &actor, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
