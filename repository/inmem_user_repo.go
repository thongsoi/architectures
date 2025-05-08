package repository

import (
	"errors"

	"github.com/thongsoi/architectures/entity"
)

var users = make(map[string]entity.User)

// UserRepository defines the contract for user persistence
type UserRepository interface {
	Create(user entity.User) error
	FindByID(id string) (entity.User, error)
}

type InMemUserRepository struct{}

func (r InMemUserRepository) Create(user entity.User) error {
	users[user.ID] = user
	return nil
}

func (r InMemUserRepository) FindByID(id string) (entity.User, error) {
	user, exists := users[id]
	if !exists {
		return entity.User{}, ErrUserNotFound
	}
	return user, nil
}

var ErrUserNotFound = errors.New("user not found")
