// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ChaoxingUser is the golang structure of table chaoxing_user for DAO operations like Where/Data.
type ChaoxingUser struct {
	g.Meta     `orm:"table:chaoxing_user, do:true"`
	Id         interface{} //
	Phone      interface{} // 超星登录手机号
	Password   interface{} // 登录密码
	Cookies    interface{} // 登录cookies
	UpdateTime *gtime.Time //
	CreateTime *gtime.Time //
}
