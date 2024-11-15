package chaoxing

import "testing"

func TestLogin(t *testing.T) {
	t.Log("测试登录接口")
	_, err := ChaoxingUserLogin("18902833039", "Yu20030405")
	if err != nil {
		t.Fatal(err)
	}
}
