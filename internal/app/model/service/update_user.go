package service

import (
	"github.com/jbohme/crud/configs/logger"
	"github.com/jbohme/crud/configs/rest_err"
	"github.com/jbohme/crud/internal/app/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserServices(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUser model.",
		zap.String("journey", "updateUser"))

	if user, _ := ud.FindUserByNickNameServices(userDomain.GetNickName()); user != nil {
		if user.GetID() != userId {
			return rest_err.NewBadRequestError("Nick Name is already in use")
		}
	}

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateUser"))
		return err
	}

	logger.Info("UpdateUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))
	return nil
}
