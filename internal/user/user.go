package user

import (
	"context"

	db "github.com/Dev-Awaab/go-base-api/db/sqlc"
)


type User struct {
	ID    int64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	Create(ctx context.Context, params db.CreateUserParams) (*db.User, error)
}


type UserService interface {
	Create(c context.Context, req *CreateUserReq)(*CreateUserRes, error)
}

type CreateUserReq struct {
	Name     string `json:"name" binding:"required,min=2"` // Name is required with at least 2 characters
    Email    string `json:"email" binding:"required,email"` // Must be a valid email
    Password string `json:"password" binding:"required,min=6"`
}
type CreateUserRes struct {
	ID    int64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}