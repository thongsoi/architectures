package usecase

import (
	"github.com/thongsoi/architectures/entity"
	"github.com/thongsoi/architectures/repository"

	"github.com/google/uuid"
)

// RegisterUserInput represents input to RegisterUser use case
type RegisterUserInput struct {
	Name string
}

// RegisterUserOutput represents output from RegisterUser use case
type RegisterUserOutput struct {
	ID   string
	Name string
}

// RegisterUserInteractor contains dependencies
type RegisterUserInteractor struct {
	UserRepo repository.UserRepository
}

// Execute runs the use case
func (uc RegisterUserInteractor) Execute(input RegisterUserInput) (RegisterUserOutput, error) {
	user := entity.User{
		ID:   uuid.New().String(),
		Name: input.Name,
	}
	err := uc.UserRepo.Create(user)
	if err != nil {
		return RegisterUserOutput{}, err
	}
	return RegisterUserOutput{ID: user.ID, Name: user.Name}, nil
}
