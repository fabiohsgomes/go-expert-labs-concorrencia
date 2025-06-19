package user_usecase

import (
	"context"
	"fullcycle-auction_go/internal/entity/user_entity"
	"fullcycle-auction_go/internal/internal_error"

	"github.com/google/uuid"
)

type UserUseCaseInterface interface {
	CreateUser(ctx context.Context, input UserInputDTO) (*UserOutputDTO, *internal_error.InternalError)
	FindUsers(ctx context.Context) ([]UserOutputDTO, *internal_error.InternalError)
	FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError)
}

type UserUseCase struct {
	UserRepository user_entity.UserRepositoryInterface
}

func NewUserUseCase(userRepository user_entity.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		userRepository,
	}
}

type UserInputDTO struct {
	Name string `json:"name"`
}

type UserOutputDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (u *UserUseCase) CreateUser(ctx context.Context, input UserInputDTO) (dto *UserOutputDTO, err *internal_error.InternalError) {
	user := &user_entity.User{
		Id:   uuid.New().String(),
		Name: input.Name,
	}

	err = u.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return dto, err
	}

	dto = &UserOutputDTO{
		Id:   user.Id,
		Name: user.Name,
	}

	return dto, err
}

func (u *UserUseCase) FindUsers(ctx context.Context) ([]UserOutputDTO, *internal_error.InternalError) {
	usersEntity, err := u.UserRepository.FindUsers(ctx)
	if err != nil {
		return nil, err
	}

	var usersOutput []UserOutputDTO
	for _, user := range usersEntity {
		usersOutput = append(usersOutput, UserOutputDTO{
			Id:   user.Id,
			Name: user.Name,
		})
	}

	return usersOutput, nil
}

func (u *UserUseCase) FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError) {
	userEntity, err := u.UserRepository.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		Id:   userEntity.Id,
		Name: userEntity.Name,
	}, nil
}
