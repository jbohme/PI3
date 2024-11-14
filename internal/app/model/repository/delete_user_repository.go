package repository

import (
	"context"
	"github.com/jbohme/crud/configs/logger"
	"github.com/jbohme/crud/configs/rest_err"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
	"os"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Init deleteUser repository",
		zap.String("journey", "deleteUser"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, _ := bson.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to delete user",
			err,
			zap.String("journey", "deleteUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("User delete successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))

	return nil
}
