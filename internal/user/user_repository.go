package user

import (
	"context"
	"database/sql"

	db "github.com/Dev-Awaab/go-base-api/db/sqlc" // Import sqlc package
)


type userRepository struct {
	queries *db.Queries
}


func NewUserRepository(dbConn *sql.DB) UserRepository {
	return &userRepository{
		queries: db.New(dbConn),
	}
}


func (u *userRepository) Create(ctx context.Context, params db.CreateUserParams) (*db.User, error) {
	user, err := u.queries.CreateUser(ctx, params) 
	if err != nil {
		return nil, err
	}
	return &user, nil
}
