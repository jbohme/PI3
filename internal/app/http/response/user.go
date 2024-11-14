package response

type UserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	NickName string `json:"nick_name"`
	QtyWins  uint   `json:"qty_wins"`
}
