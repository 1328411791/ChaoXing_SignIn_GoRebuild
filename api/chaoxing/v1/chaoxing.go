package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/chaoxing/login" method:"post" tags:"chaoxing" summary:"登录"`
	Phone    string `json:"phone" dc:"手机号"`
	Password string `json:"password" dc:"密码"`
}

type LoginRes struct {
	Content string `json:"content" dc:"返回结果"`
}

// 查询活动接口

type GetActivityReq struct {
	g.Meta `path:"/chaoxing/activity" method:"post" tags:"chaoxing" summary:"查询活动"`
	Uid    string `json:"uid" dc:"uid"`
	D      string `json:"_d" dc:"_d"`
	Vc3    string `json:"vc3" dc:"vc3"`
	Uf     string `json:"uf" dc:"uf"`
}

type GetActivityRes struct {
	Content string `json:"content" dc:"返回结果"`
}

type ActivityAllReq struct {
	g.Meta `path:"/chaoxing/activityAll" method:"post" tags:"chaoxing" summary:"获取所有活动"`
	Uid    string `json:"uid" dc:"uid"`
	D      string `json:"_d" dc:"_d"`
	Vc3    string `json:"vc3" dc:"vc3"`
	Uf     string `json:"uf" dc:"uf"`
}

type ActivityAllRes struct {
	Content string `json:"content" dc:"返回结果"`
}

type QRCodeSignReq struct {
	g.Meta   `path:"/chaoxing/qrcode" method:"post" tags:"chaoxing" summary:"扫码签到"`
	Name     string `json:"name" dc:"name"`
	Fid      string `json:"fid" dc:"fid"`
	Uid      string `json:"uid" dc:"uid"`
	ActiveId string `json:"activeId" dc:"activeId"`
	Uf       string `json:"uf" dc:"uf"`
	D        string `json:"_d" dc:"_d"`
	Vc3      string `json:"vc3" dc:"vc3"`
	Enc      string `json:"enc" dc:"enc"`
	Lat      string `json:"lat" dc:"lat"`
	Lon      string `json:"lon" dc:"lon"`
	Address  string `json:"address" dc:"address"`
	Altitude string `json:"altitude" dc:"altitude"`
	Enc2     string `json:"enc2" dc:"enc2"`
	Validate string `json:"validate" dc:"validate"`
}

type QRCodeSignRes struct {
	Content string `json:"content" dc:"返回结果"`
}

type LocationSignReq struct {
	g.Meta   `path:"/chaoxing/location" method:"post" tags:"chaoxing" summary:"位置签到"`
	Uf       string `json:"uf" dc:"uf"`
	D        string `json:"_d" dc:"_d"`
	Vc3      string `json:"vc3" dc:"vc3"`
	Name     string `json:"name" dc:"name"`
	Uid      string `json:"uid" dc:"uid"`
	Lat      string `json:"lat" dc:"lat"`
	Lon      string `json:"lon" dc:"lon"`
	Fid      string `json:"fid" dc:"fid"`
	Address  string `json:"address" dc:"address"`
	ActiveId string `json:"activeId" dc:"activeId"`
	Validate string `json:"validate" dc:"validate"`
}

type LocationSignRes struct {
	Content string `json:"content" dc:"返回结果"`
}

type GeneralSignReq struct {
	g.Meta   `path:"/chaoxing/general" method:"post" tags:"chaoxing" summary:"通用签到"`
	Uf       string `json:"uf" dc:"uf"`
	D        string `json:"_d" dc:"_d"`
	Vc3      string `json:"vc3" dc:"vc3"`
	Name     string `json:"name" dc:"name"`
	ActiveId string `json:"activeId" dc:"activeId"`
	Uid      string `json:"uid" dc:"uid"`
	Fid      string `json:"fid" dc:"fid"`
	SignCode string `json:"signCode" dc:"signCode"`
	Validate string `json:"validate" dc:"validate"`
}

type GeneralSignRes struct {
	Content string `json:"content" dc:"返回结果"`
}

type GetSlideCaptchaReq struct {
	g.Meta `path:"/chaoxing/getSlideCaptcha" method:"post" tags:"chaoxing" summary:"获取滑块验证码"`
	Uf     string `json:"uf" dc:"uf"`
	D      string `json:"_d" dc:"_d"`
	Vc3    string `json:"vc3" dc:"vc3"`
	Uid    string `json:"uid" dc:"uid"`
}

type GetSlideCaptchaRes struct {
	Content string `json:"content" dc:"返回结果"`
}

type CheckSlideCaptchaReq struct {
	g.Meta       `path:"/chaoxing/checkSlideCaptcha" method:"post" tags:"chaoxing" summary:"检查滑块验证码"`
	Uf           string `json:"uf" dc:"uf"`
	D            string `json:"_d" dc:"_d"`
	Vc3          string `json:"vc3" dc:"vc3"`
	Uid          string `json:"uid" dc:"uid"`
	Token        string `json:"token" dc:"token"`
	TextClickArr string `json:"textClickArr" dc:"textClickArr"`
}

type CheckSlideCaptchaRes struct {
	Content string `json:"content" dc:"返回结果"`
}

type UvtokenReq struct {
	g.Meta `path:"/chaoxing/uvtoken" method:"post" tags:"chaoxing" summary:"获取UVToken"`
	Uf     string `json:"uf" dc:"uf"`
	D      string `json:"_d" dc:"_d"`
	Uid    string `json:"uid" dc:"uid"`
	Vc3    string `json:"vc3" dc:"vc3"`
}

type UvtokenRes struct {
	Content string `json:"content" dc:"返回结果"`
}

type UploadReq struct {
	g.Meta `path:"/chaoxing/upload" method:"post" tags:"chaoxing" summary:"上传文件"`
	Uf     string `json:"uf" dc:"uf"`
	D      string `json:"_d" dc:"_d"`
	Uid    string `json:"uid" dc:"uid"`
	Vc3    string `json:"vc3" dc:"vc3"`
	Token  string `json:"token" dc:"token"`
	Buffer []byte `json:"buffer" dc:"buffer"`
}

type UploadRes struct {
	Content string `json:"content" dc:"返回结果"`
}

type PhotoSignReq struct {
	g.Meta   `path:"/chaoxing/photo" method:"post" tags:"chaoxing" summary:"拍照签到"`
	Uf       string `json:"uf" dc:"uf"`
	D        string `json:"_d" dc:"_d"`
	Vc3      string `json:"vc3" dc:"vc3"`
	Name     string `json:"name" dc:"name"`
	ActiveId string `json:"activeId" dc:"activeId"`
	Uid      string `json:"uid" dc:"uid"`
	Fid      string `json:"fid" dc:"fid"`
	ObjectId string `json:"objectId" dc:"objectId"`
	Validate string `json:"validate" dc:"validate"`
}

type PhotoSignRes struct {
	Content string `json:"content" dc:"返回结果"`
}
