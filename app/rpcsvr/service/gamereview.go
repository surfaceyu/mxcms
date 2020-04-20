package service

import (
	"mxcms/app"
	"mxcms/app/models"
)

func GetGameReview(gameid string) (*models.GameReview, error) {
	result := models.GameReview{
		GameId:gameid,
	}
	app.Db.First(&result)
	return &result, nil
}
