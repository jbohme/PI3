package repository

import (
	"github.com/jbohme/crud/configs/rest_err"
	"github.com/jbohme/crud/internal/app/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr
	FindUserByID(userId string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(userEmail string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByNickName(userNickName string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailAndPassword(userEmail string, password string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(userId string) *rest_err.RestErr
}
