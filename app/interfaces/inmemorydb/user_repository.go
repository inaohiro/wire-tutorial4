package inmemorydb

import (
	"context"
	"errors"
	"github.com/google/wire"
	"github.com/inaohiro-sandbox/wire-tutorial4/app/domain"
	"log"
)

var Set = wire.NewSet(
	NewUserRepository,
)

var store = map[int]*domain.User{}

type InMemoryUserRepository struct{}

//func NewUserRepository() domain.UserRepository {
//	return &InMemoryUserRepository{}
//}

func NewUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{}
}

func (repo *InMemoryUserRepository) Get(ctx context.Context, id int) (user *domain.User, err error) {
	if found, ok := store[id]; ok {
		user = found
	} else {
		return nil, errors.New("no such entity")
	}
	return user, nil
}

func (repo *InMemoryUserRepository) Save(ctx context.Context, user *domain.User) (err error) {
	if user.ID == 0 {
		user.ID = int(len(store) + 1)
	}
	store[user.ID] = user

	return nil
}

func Dump() {
	for _, v := range store {
		log.Println(v)
	}
}
