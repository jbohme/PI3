package service

import (
	"github.com/jbohme/crud/configs/logger"
	"github.com/jbohme/crud/configs/rest_err"
	"github.com/jbohme/crud/internal/app/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID services.",
		zap.String("journey", "findUserByID"))

	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail services.",
		zap.String("journey", "findUserByEmail"))

	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) FindUserByNickNameServices(nickName string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByNickName services.",
		zap.String("journey", "findUserByNickName"))

	return ud.userRepository.FindUserByNickName(nickName)
}

func (ud *userDomainService) findUserByEmailAndPasswordServices(
	email string,
	password string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmailAndPassword services.",
		zap.String("journey", "findUserByEmailAndPassword"))

	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
