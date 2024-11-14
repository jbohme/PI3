package converter

import (
	"github.com/jbohme/crud/internal/app/model"
	"github.com/jbohme/crud/internal/app/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.NickName,
	)
	domain.SetID(entity.ID.Hex())
	return domain
}
