package model

type userDomain struct {
	id       string
	email    string
	password string
	name     string
	nickName string
	qtyWins  uint
}

func (ud *userDomain) SetID(id string) {
	ud.id = id
}

func (ud *userDomain) newWin() {
	ud.qtyWins += 1
}

func (ud *userDomain) GetID() string {
	return ud.id
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetNickName() string {
	return ud.nickName
}

func (ud *userDomain) GetWins() uint {
	return ud.qtyWins
}
