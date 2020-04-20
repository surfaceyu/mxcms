package handler

import (
	"context"
	"mxcms/app/models"
	"mxcms/app/rpcsvr/service"
)

func (this *Handlers) GetGameReview(ctx context.Context, req string, rsp *models.GameReview) error {
	result, err := service.GetGameReview(req)
	*rsp = *result
	return err
}
