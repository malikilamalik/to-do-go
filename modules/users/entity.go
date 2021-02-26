package users

type User struct {
	Id       int64  `xorm:"id" json:"id"`
	Username string `xorm:"username" json:"username" `
	Hash     string `xorm:"hash" json:"-"`
}
