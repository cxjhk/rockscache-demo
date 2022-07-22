package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"awesome/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type HandlerGetFriendProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var (
	CacheFriendProfileKey = "cache:friend:%d:profile:%d"
)

func NewHandlerGetFriendProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandlerGetFriendProfileLogic {
	return &HandlerGetFriendProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandlerGetFriendProfileLogic) HandlerGetFriendProfile() error {
	_, err := l.svcCtx.RocksCache.Fetch(fmt.Sprintf(CacheFriendProfileKey, 417093819998142711, 417091069155475700), 300*time.Second, func() (string, error) {
		return "", errors.New("错误")
	})
	if err != nil {
		return err
	}
	return nil
}
