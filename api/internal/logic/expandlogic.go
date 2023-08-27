package logic

import (
	"context"
	"github.com/atadzan/shorturl/api/internal/svc"
	"github.com/atadzan/shorturl/api/internal/types"
	"github.com/atadzan/shorturl/rpc/transform/transformer"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandLogic) Expand(req *types.ExpandReq) (*types.ExpandResp, error) {
	// manual code
	rpcResp, err := l.svcCtx.Transformer.Expand(l.ctx, &transformer.ExpandReq{
		Shorten: req.Shorten,
	})
	if err != nil {
		return nil, err
	}
	return &types.ExpandResp{
		Url: rpcResp.Url,
	}, nil
}
