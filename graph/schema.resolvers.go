package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/hanafiadhi/MyGrams/graph/generated"
	"github.com/hanafiadhi/MyGrams/graph/model"
	"github.com/hanafiadhi/MyGrams/internal/config"
	"github.com/hanafiadhi/MyGrams/internal/pkg/jwt"
)

var (
	db = config.ConnectDB()
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.NewUser) (string, error) {
	user, err := db.Register(&input)
	if err != nil {
		return "", err
	}
	token, err := jwt.GenerateToken(user)
	return token, err
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginUser) (string, error) {
	correct, err := db.Login(&input)
	if err != nil {
		return "", err
	}
	token, err := jwt.GenerateToken(correct)
	return token, err
}

// Createuser is the resolver for the createuser field.
func (r *mutationResolver) Createuser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user, err := db.CreateUser(&input)
	return user, err
}

// Updateuser is the resolver for the updateuser field.
func (r *mutationResolver) Updateuser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Updateuser - updateuser"))
}

// Createphoto is the resolver for the createphoto field.
func (r *mutationResolver) Createphoto(ctx context.Context, input *model.NewPhoto) (*model.Photo, error) {
	panic(fmt.Errorf("not implemented: Createphoto - createphoto"))
}

// Updatephoto is the resolver for the updatephoto field.
func (r *mutationResolver) Updatephoto(ctx context.Context, input *model.UpdatePhoto) (*model.Photo, error) {
	panic(fmt.Errorf("not implemented: Updatephoto - updatephoto"))
}

// Deletephoto is the resolver for the deletephoto field.
func (r *mutationResolver) Deletephoto(ctx context.Context, input *model.DeletePhoto) (*model.Photo, error) {
	panic(fmt.Errorf("not implemented: Deletephoto - deletephoto"))
}

// Cratecomment is the resolver for the cratecomment field.
func (r *mutationResolver) Cratecomment(ctx context.Context, input *model.NewComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: Cratecomment - cratecomment"))
}

// Updatecomment is the resolver for the updatecomment field.
func (r *mutationResolver) Updatecomment(ctx context.Context, input *model.UpdateComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: Updatecomment - updatecomment"))
}

// Deletecomment is the resolver for the deletecomment field.
func (r *mutationResolver) Deletecomment(ctx context.Context, input *model.DeleteComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: Deletecomment - deletecomment"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Photo is the resolver for the photo field.
func (r *queryResolver) Photo(ctx context.Context) ([]*model.Photo, error) {
	panic(fmt.Errorf("not implemented: Photo - photo"))
}

// Comment is the resolver for the comment field.
func (r *queryResolver) Comment(ctx context.Context) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented: Comment - comment"))
}

// Photos is the resolver for the photos field.
func (r *userResolver) Photos(ctx context.Context, obj *model.User) ([]*model.Photo, error) {
	panic(fmt.Errorf("not implemented: Photos - photos"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
type todoResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	return input.Username, nil
}
