package todos

type Todo struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"user_id"`
	Title  string `json:"title" `
	Task   string `json:"task"`
	Status string `json:"status"`
}
