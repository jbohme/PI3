package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jbohme/crud/configs/logger"
	"github.com/jbohme/crud/configs/rest_err"
	"github.com/jbohme/crud/internal/app/view"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
	"net/http"
	"net/mail"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init findUserByID",
		zap.String("journey", "findUserByID"))

	userId := c.Param("userId")
	if _, err := bson.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userId",
			err,
			zap.String("journey", "findUserByID"))

		errorMessage := rest_err.NewBadRequestError("UserID is not a valid id")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error trying to call findUserByID services",
			err,
			zap.String("journey", "findUserByID"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByID controller executed successfully",
		zap.String("journey", "findUserByID"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail",
		zap.String("journey", "findUserByEmail"))

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate userEmail",
			err,
			zap.String("journey", "findUserByEmail"))

		errorMessage := rest_err.NewBadRequestError("UserEmail is not a valid email")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call findUserByEmail services",
			err,
			zap.String("journey", "findUserByEmail"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller executed successfully",
		zap.String("journey", "findUserByEmail"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
