package service

import (
	"context"
	"github.com/towelong/egret/example/api/v1"
)

func (u *ShopInterface) GetUserById(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserResp, error) {
	resp, err := u.userUsecase.GetUserById(ctx, req.UserId)
	return &v1.GetUserResp{
		Id:   resp.ID,
		Name: resp.Name,
		Age:  resp.Age,
	}, err
}
