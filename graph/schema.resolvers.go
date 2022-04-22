package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	generated1 "iss-superfl-response-writeheader/graph/generated"
	model1 "iss-superfl-response-writeheader/graph/model"
	"time"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model1.NewTodo) (*model1.Todo, error) {
	return &model1.Todo{ID: 1, Text: "1", Done: false}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model1.Todo, error) {
	return []*model1.Todo{
		{ID: 0, Text: "0", Done: true},
	}, nil
}

func (r *subscriptionResolver) Todo(ctx context.Context) (<-chan *model1.Todo, error) {
	ticker := time.NewTicker(1 * time.Second)

	ch := make(chan *model1.Todo)

	go func() {
		for {
			<-ticker.C
			ch <- &model1.Todo{
				ID:   uint64(time.Now().Unix()),
				Text: time.Now().String(),
			}
		}
	}()

	return ch, nil
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

// Subscription returns generated1.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated1.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
