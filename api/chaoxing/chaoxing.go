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
	GetActivity(ctx context.Context, req *v1.GetActivityReq) (res *v1.GetActivityRes, err error)
	ActivityAll(ctx context.Context, req *v1.ActivityAllReq) (res *v1.ActivityAllRes, err error)
	QRCodeSign(ctx context.Context, req *v1.QRCodeSignReq) (res *v1.QRCodeSignRes, err error)
	LocationSign(ctx context.Context, req *v1.LocationSignReq) (res *v1.LocationSignRes, err error)
	GeneralSign(ctx context.Context, req *v1.GeneralSignReq) (res *v1.GeneralSignRes, err error)
	GetSlideCaptcha(ctx context.Context, req *v1.GetSlideCaptchaReq) (res *v1.GetSlideCaptchaRes, err error)
	CheckSlideCaptcha(ctx context.Context, req *v1.CheckSlideCaptchaReq) (res *v1.CheckSlideCaptchaRes, err error)
	Uvtoken(ctx context.Context, req *v1.UvtokenReq) (res *v1.UvtokenRes, err error)
	Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error)
	PhotoSign(ctx context.Context, req *v1.PhotoSignReq) (res *v1.PhotoSignRes, err error)
}
