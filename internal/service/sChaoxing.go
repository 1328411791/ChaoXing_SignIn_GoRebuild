package service

import (
	"ChaoXing_SignIn_GoRebuild/internal/dao"
	"ChaoXing_SignIn_GoRebuild/internal/logic/chaoxing"
	"encoding/json"
)

type SChaoxing struct{}

func init() {
	RegisterChaoxing(New())
}

func New() *SChaoxing {
	return &SChaoxing{}
}

func (s *SChaoxing) UserLogin(phone string, password string) (err error) {
	// 登录获取cookie

	login, err := chaoxing.ChaoxingUserLogin(phone, password)
	if err != nil {
		return err
	}

	// 检查数据库是否存在该用户，查询手机号是否存在
	record, err := dao.ChaoxingUser.DB().Model("chaoxing_user").Where("phone =?", phone).One()
	if err != nil {
		return err
	}

	// 将cookies转化为json字符串
	marshal, err := json.Marshal(login)
	cookiesJson := string(marshal)

	if err != nil {
		return err
	}

	if record.IsEmpty() {
		// 不存在则插入
		_, err = dao.ChaoxingUser.DB().Model("chaoxing_user").Data(map[string]interface{}{
			"phone":    phone,
			"password": password,
			"cookies":  cookiesJson,
		}).Insert()
		if err != nil {
			return err
		}
	} else {
		// 存在则更新

		_, err = dao.ChaoxingUser.DB().Model("chaoxing_user").Data(map[string]interface{}{
			"phone":    phone,
			"password": password,
			"cookies":  cookiesJson,
		}).Where("phone =?", phone).Update()
	}

	return nil
}

func (s *SChaoxing) GetCourses(userId int) (*chaoxing.CourseType, error) {
	// 从数据库读取cookies
	cookies, err := getCookies4Id(userId)
	if err != nil {
		return nil, err
	}

	chaoxing.GetCourses(cookies)

	return nil, nil
}

func getCookies4Id(id int) ([]string, error) {
	one, err := dao.ChaoxingUser.DB().Model("chaoxing_user").Where("id=?", id).One()
	if err != nil {
		return nil, err
	}
	if one == nil {
		return nil, nil
	}

	cookiesJson := one["cookies"].String()

	// 将str转化为string[]
	var cookies []string
	err = json.Unmarshal([]byte(cookiesJson), &cookies)
	if err != nil {
		return nil, err
	}
	return cookies, nil
}
