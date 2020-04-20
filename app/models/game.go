package models

type GameReview struct{
	GameId string `json:"gameid" gorm:"primary_key"`
	Name string  `json:"name" gorm:"type:varchar(20);not null;"`
	Avatar string `json:"avatar" gorm:"type:text;not null;"`
	Platform string `json:"platform" gorm:"type:varchar(128);not null;"`
	Good string `json:"good" gorm:"type:text;not null;"`
	Bad string `json:"Bad" gorm:"type:text;not null;"`
	Score float32 `json:"score" gorm:"type:int(8);not null;"`
	Scoreword string `json:"scoreword" gorm:"type:varchar(20);not null;"`
}


