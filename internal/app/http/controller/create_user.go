package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jbohme/crud/configs/logger"
	"github.com/jbohme/crud/configs/validation"
	"github.com/jbohme/crud/internal/app/http/request"
	"github.com/jbohme/crud/internal/app/model"
	"github.com/jbohme/crud/internal/app/view"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init createUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.NickName,
	)

	domainResult, err := uc.service.CreateUserServices(domain)
	if err != nil {
		logger.Error("Error trying to call CreateUser service", err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("CreateUser controller executed successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "createUser"),
	)
	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(
		domainResult,
	))
}
