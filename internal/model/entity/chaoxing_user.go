// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ChaoxingUser is the golang structure for table chaoxing_user.
type ChaoxingUser struct {
	Id         int         `json:"id"         orm:"id"          description:""`
	Phone      string      `json:"phone"      orm:"phone"       description:"超星登录手机号"`
	Password   string      `json:"password"   orm:"password"    description:"登录密码"`
	Cookies    string      `json:"cookies"    orm:"cookies"     description:"登录cookies"`
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:""`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""`
}
