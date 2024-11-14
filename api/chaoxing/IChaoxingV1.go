package chaoxing

import (
	"ChaoXing_SignIn_GoRebuild/api"
	v1 "ChaoXing_SignIn_GoRebuild/api/chaoxing/v1"
	"context"
)

type IChaoxingV1 interface {
	ChaoxingUserLogin(ctx context.Context, req *v1.LoginReq) (res *api.Response, err error)
}
