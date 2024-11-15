package chaoxing

import (
	v1 "ChaoXing_SignIn_GoRebuild/api/chaoxing/v1"
	"ChaoXing_SignIn_GoRebuild/internal/service"
	"context"
)

func (c ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	//TODO implement me

	err = service.Chaoxing().UserLogin(req.Phone, req.Password)
	if err != nil {
		return &v1.LoginRes{Content: "登录失败"}, err
	} else {
		return &v1.LoginRes{Content: "登录成功"}, err
	}
}
