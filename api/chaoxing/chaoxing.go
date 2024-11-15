// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package chaoxing

import (
	"context"

	"ChaoXing_SignIn_GoRebuild/api/chaoxing/v1"
)

type IChaoxingV1 interface {
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
}
