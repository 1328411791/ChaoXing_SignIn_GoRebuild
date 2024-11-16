package chaoxing

import (
	"testing"
)

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
	// 检查返回的课程列表是否为空
	if len(courses) == 0 {
		t.Fatal("返回的课程列表为空")
	}
}

// TestAccountInfo 测试获取账号信息接口
func TestGetCoursesByCourses(t *testing.T) {
	t.Log("测试获取账号信息接口")
	_, err := ChaoxingUserLogin("18902833039", "Yu20030405")
	if err != nil {
		t.Fatal(err)
	}

}

// TestGetActiveByCourses 测试获取活动信息接口
func TestGetActiveByCourses(t *testing.T) {
	t.Log("测试获取活动信息接口")
	cookies, err := ChaoxingUserLogin("18902833039", "Yu20030405")
	if err != nil {
		t.Fatal(err)
	}
	courses, err := GetCourses(cookies)
	if err != nil {
		t.Fatal(err)
	}

	allActive, err := GetAllActive(courses, cookies)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(allActive)

}
