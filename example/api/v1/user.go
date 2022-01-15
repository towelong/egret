package v1

import (
	"github.com/gin-gonic/gin"
)

type GetUserReq struct {
	UserId int64 `json:"user_id" uri:"user_id" validate:"min=1"`
}

type GetUserResp struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func (s *ShopService) GetUserById(ctx *gin.Context) {
	var in GetUserReq
	if err := ctx.ShouldBindUri(&in); err != nil {
		_ = ctx.Error(err)
		return
	}
	out, err := s.server.GetUserById(ctx, &in)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(200, out)
}
