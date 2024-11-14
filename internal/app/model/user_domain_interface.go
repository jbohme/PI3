package model

import "github.com/jbohme/crud/configs/rest_err"

type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetWins() uint
	GetNickName() string
	SetID(string)

	newWin()

	EncryptPassword()
	GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(
	email, password, name, nickName string,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		nickName: nickName,
		qtyWins:  0,
	}
}

func NewUserUpdateDomain(
	name, nickName string,
) UserDomainInterface {
	return &userDomain{
		name:     name,
		nickName: nickName,
	}
}

func NewUserLoginDomain(
	email, password string,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}
