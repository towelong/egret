package domain

import "context"

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type UserRepo interface {
	GetUserById(ctx context.Context, id int64) (*User, error)
}

type UserUsecase interface {
	GetUserById(ctx context.Context, id int64) (*User, error)
}
