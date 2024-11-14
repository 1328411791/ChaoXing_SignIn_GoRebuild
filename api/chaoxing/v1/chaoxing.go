package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"chaoxing" summary:"登录"`
	Phone    string `json:"phone" dc:"手机号"`
	Password string `json:"password" dc:"密码"`
}

type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
