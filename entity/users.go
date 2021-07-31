package entity

//Users object for REST(CRUD)
type Users struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}
