package service

import (
	"github.com/jbohme/crud/configs/rest_err"
	"github.com/jbohme/crud/internal/app/model"
	"github.com/jbohme/crud/internal/app/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserServices(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr
	FindUserByNickNameServices(userNickName string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(userEmail string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByIDServices(userId string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUserServices(userId string) *rest_err.RestErr
	LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr)

	//JoinRandomRoomServices(playerDomain model.PlayerDomainInterface) (model.RoomDomain, *rest_err.RestErr)
}
