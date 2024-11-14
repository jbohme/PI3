package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required,min=4,max=100" `
	NickName string `json:"nick_name" binding:"required,min=3,max=50"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" binding:"omitempty,min=4,max=100"`
	NickName string `json:"nick_name" binding:"required,min=3,max=50"`
}
