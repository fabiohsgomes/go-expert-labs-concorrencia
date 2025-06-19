package user_entity

import (
	"context"
	"fullcycle-auction_go/internal/internal_error"
)

type User struct {
	Id   string
	Name string
}

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *User) *internal_error.InternalError
	FindUsers(ctx context.Context) ([]User, *internal_error.InternalError)
	FindUserById(ctx context.Context, userId string) (*User, *internal_error.InternalError)
}
