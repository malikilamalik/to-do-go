package users

type User struct {
	Id       int64  `json:"id"`
	Username string `xorm:"not null unique 'usr_name'" json:"username" `
	Hash     string `json:"-"`
}
