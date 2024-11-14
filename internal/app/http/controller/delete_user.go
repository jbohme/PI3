package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jbohme/crud/configs/logger"
	"github.com/jbohme/crud/configs/rest_err"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init deleteUser http",
		zap.String("journey", "deleteUser"),
	)

	userId := c.Param("userId")
	if _, err := bson.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	err := uc.service.DeleteUserServices(userId)
	if err != nil {
		logger.Error("Error trying to call deleteUser service", err,
			zap.String("journey", "deleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("UpdateUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)
	c.Status(http.StatusOK)
}
