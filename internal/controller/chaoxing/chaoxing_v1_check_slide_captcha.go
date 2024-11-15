package chaoxing

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"ChaoXing_SignIn_GoRebuild/api/chaoxing/v1"
)

func (c *ControllerV1) CheckSlideCaptcha(ctx context.Context, req *v1.CheckSlideCaptchaReq) (res *v1.CheckSlideCaptchaRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
