package logic

import (
	"context"

	"github.com/atadzan/shorturl/internal/svc"
	"github.com/atadzan/shorturl/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req *types.ShortenReq) (resp *types.ShortenResp, err error) {
	resp = new(types.ShortenResp)
	resp.Shorten = "your url will be shortened soon"
	return
}
