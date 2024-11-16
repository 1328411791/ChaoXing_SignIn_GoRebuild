package chaoxing

import "testing"

func TestLogin(t *testing.T) {
	t.Log("测试登录接口")

	cookies, err := ChaoxingUserLogin("18902833039", "Yu20030405")
	err = GetAccountInfo(cookies)
	if err != nil {
		t.Fatal(err)
	}
}
