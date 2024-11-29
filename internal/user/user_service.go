package user

import (
	"context"
	"time"

	db "github.com/Dev-Awaab/go-base-api/db/sqlc"
)

type service struct {
	UserRepository
	timeout time.Duration
}



func NewUserService(ur UserRepository) UserService {
	return &service{
		ur,
		time.Duration(5) * time.Second,
	}
}


// Create implements UserService.
// Subtle: this method shadows the method (UserRepository).Create of service.UserRepository.
func (s *service) Create(c context.Context, req *CreateUserReq)(*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	

	params := db.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password, 
	}

	user, err := s.UserRepository.Create(ctx,params)
	if err != nil {
		return nil, err
	}
	return &CreateUserRes{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}