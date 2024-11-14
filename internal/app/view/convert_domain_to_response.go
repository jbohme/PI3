package view

import (
	"github.com/jbohme/crud/internal/app/http/response"
	"github.com/jbohme/crud/internal/app/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:       userDomain.GetID(),
		Email:    userDomain.GetEmail(),
		Name:     userDomain.GetName(),
		NickName: userDomain.GetNickName(),
		QtyWins:  userDomain.GetWins(),
	}
}
