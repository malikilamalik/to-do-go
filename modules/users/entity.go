package users

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username" `
	Hash     string `json:"-"`
}
