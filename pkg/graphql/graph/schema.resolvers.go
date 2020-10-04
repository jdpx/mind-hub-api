package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/generated"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/model"
)

func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetRequestContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

func GetNestedPreloads(ctx *graphql.RequestContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.SelectionSet, nil), prefixColumn)...)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}
func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	preloads := graphql.GetOperationContext(ctx)
	fmt.Println(preloads.RawQuery)

	c, err := r.resolveCourses(ctx, preloads.RawQuery)

	return c, err
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
