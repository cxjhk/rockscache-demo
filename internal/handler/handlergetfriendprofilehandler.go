package handler

import (
	"net/http"

	"awesome/internal/logic"
	"awesome/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func handlerGetFriendProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHandlerGetFriendProfileLogic(r.Context(), svcCtx)
		err := l.HandlerGetFriendProfile()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
