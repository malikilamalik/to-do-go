package users

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username" `
	Hash     string `json:"-"`
}

//Request For Login
type Request struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

//Returen token after login
type LoginResponse struct {
	Token string `json:"token"`
}
