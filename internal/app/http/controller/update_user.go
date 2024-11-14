package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jbohme/crud/configs/logger"
	"github.com/jbohme/crud/configs/rest_err"
	"github.com/jbohme/crud/configs/validation"
	"github.com/jbohme/crud/internal/app/http/request"
	"github.com/jbohme/crud/internal/app/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init updateUser http",
		zap.String("journey", "updateUser"),
	)
	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "updateUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	userId := c.Param("userId")
	if _, err := bson.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.NickName,
	)

	err := uc.service.UpdateUserServices(userId, domain)
	if err != nil {
		logger.Error("Error trying to call UpdateUser service", err,
			zap.String("journey", "updateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("UpdateUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)
	c.Status(http.StatusOK)
}
