package chaoxing

import "testing"

// TestLogin 测试登录接口
func TestLogin(t *testing.T) {
	t.Log("测试登录接口")

	_, err := ChaoxingUserLogin("18902833039", "Yu20030405")
	if err != nil {
		t.Fatal(err)
	}
}

// TestGetCourses 测试获取课程列表接口
func TestGetCourses(t *testing.T) {
	t.Log("测试获取课程列表接口")
	cookies, err := ChaoxingUserLogin("18902833039", "Yu20030405")
	if err != nil {
		t.Fatal(err)
	}
	courses, err := GetCourses(cookies)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(courses)
}

// TestAccountInfo 测试获取账号信息接口
func TestGetCoursesByCourses(t *testing.T) {
	t.Log("测试获取账号信息接口")
	cookies, err := ChaoxingUserLogin("18902833039", "Yu20030405")
	if err != nil {
		t.Fatal(err)
	}
	GetAccountInfo(cookies)
}
