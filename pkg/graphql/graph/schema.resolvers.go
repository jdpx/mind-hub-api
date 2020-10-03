package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/generated"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/model"
)

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	return []*model.Course{&courseOne, &courseTwo}, nil
}

func (r *queryResolver) Sessions(ctx context.Context) ([]*model.Session, error) {
	return []*model.Session{&sessionOne, &sessionTwo, &sessionThree}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var (
	sessionOne   = model.Session{ID: "1", Title: "Session One"}
	sessionTwo   = model.Session{ID: "2", Title: "Session Two"}
	sessionThree = model.Session{ID: "3", Title: "Session Three"}

	courseOne = model.Course{ID: "1", Title: "Course One", Sessions: []*model.Session{&sessionOne, &sessionTwo, &sessionThree}}
	courseTwo = model.Course{ID: "2", Title: "Course Two"}
)

type mutationResolver struct{ *Resolver }
