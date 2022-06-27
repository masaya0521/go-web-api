package interactors

import (
	"backend/entity"
	ports "backend/usecase"
	"context"
)

type UserInteractor struct {
	OutputPort ports.UserOutputPort
	Repository ports.UserRepository
}

func NewUserInputPort(outputPort ports.UserOutputPort, repository ports.UserRepository) ports.UserInputPort {
	return &UserInteractor{
		OutputPort: outputPort,
		Repository: repository,
	}
}

func (u *UserInteractor) AddUser(ctx context.Context, user *entity.User) error {
	users, err := u.Repository.AddUser(ctx, user)
	if err != nil {
		return u.OutputPort.OutputError(err)
	}

	return u.OutputPort.OutputUsers(users)
}

func (u *UserInteractor) GetUsers(ctx context.Context) error {
	users, err := u.Repository.GetUsers(ctx)
	if err != nil {
		return u.OutputPort.OutputError(err)
	}

	return u.OutputPort.OutputUsers(users)
}