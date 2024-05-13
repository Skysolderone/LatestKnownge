package logic

import (
	"context"

	"v1/rpc/check/check"
	"v1/rpc/check/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *check.CheckReq) (*check.CheckResp, error) {
	r, err := l.svcCtx.Model.FindOne(l.ctx, in.Book)
	if err != nil {
		return nil, err
	}

	return &check.CheckResp{
		Found: true,
		Price: r.Price,
	}, nil
}
