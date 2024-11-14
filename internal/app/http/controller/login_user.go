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

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init loginUser controller",
		zap.String("journey", "loginUser"),
	)
	var loginRequest request.UserLogin
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "loginUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(
		loginRequest.Email,
		loginRequest.Password,
	)

	domainResult, token, err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Error("Error trying to call CreateUser service", err,
			zap.String("journey", "loginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("CreateUser controller executed successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "loginUser"),
	)

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}
