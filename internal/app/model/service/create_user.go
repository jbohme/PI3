package service

import (
	"github.com/jbohme/crud/configs/logger"
	"github.com/jbohme/crud/configs/rest_err"
	"github.com/jbohme/crud/internal/app/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser model.",
		zap.String("journey", "createUser"))

	if user, _ := ud.FindUserByEmailServices(userDomain.GetEmail()); user != nil {
		return nil, rest_err.NewBadRequestError("Email is already in use")
	}

	if user, _ := ud.FindUserByNickNameServices(userDomain.GetNickName()); user != nil {
		return nil, rest_err.NewBadRequestError("Nick Name is already in use")
	}

	userDomain.EncryptPassword()
	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("CreateUser service executed successfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"))
	return userDomainRepository, nil
}
