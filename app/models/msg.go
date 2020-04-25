package models

type MsgLogin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Msgmenu struct {
	Menu
	Child []Menu
}
