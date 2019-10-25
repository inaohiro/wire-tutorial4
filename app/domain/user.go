package domain

import (
	"context"
	"log"
)

type User struct {
	ID   int
	Name string
}

type service struct {
	repo UserRepository
}

// usecase
type UserService interface {
	Create(ctx context.Context, name string) (created *User, err error)
	Show(ctx context.Context, id int) (user *User, err error)
}

type UserRepository interface {
	Save(ctx context.Context, user *User) (err error)
	Get(ctx context.Context, id int) (user *User, err error)
}

// interactor
func NewUserService(repo UserRepository) UserService {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, name string) (created *User, err error) {
	user := &User{Name: name}
	err = s.repo.Save(ctx, user)
	if err != nil {
		log.Printf(err.Error())
	}
	return
}

func (s *service) Show(ctx context.Context, id int) (user *User, err error) {
	user, err = s.repo.Get(ctx, id)
	if err != nil {
		log.Printf(err.Error())
	}
	return
}
