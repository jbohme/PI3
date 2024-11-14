package converter

import (
	"github.com/jbohme/crud/internal/app/model"
	"github.com/jbohme/crud/internal/app/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		NickName: domain.GetNickName(),
		QtyWins:  domain.GetWins(),
	}
}
