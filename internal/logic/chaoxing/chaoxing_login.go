package chaoxing

import (
	"ChaoXing_SignIn_GoRebuild/utility"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func ChaoxingUserLogin(username string, password string) ([]string, error) {
	// 将key转换为byte数组
	key := []byte("u2oh6Vu^")
	// 使用DES算法加密密码
	encryptedPassword, err := utility.EncryptDES(password, key)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(encryptedPassword)

	cookies, err := doLogin(username, password)

	return cookies, nil
}

func doLogin(uname, password string) ([]string, error) {
	formData := fmt.Sprintf("uname=%s&password=%s&fid=-1&t=true&refer=https%%3A%%2F%%2Fi.chaoxing.com&forbidotherlogin=0&validate=", uname, password)
	req, err := http.NewRequest(LOGIN.METHOD, LOGIN.URL, strings.NewReader(formData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	// 获取cookie

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if status, ok := result["status"].(bool); ok && status {
		fmt.Println("登录成功")
		cookies := resp.Header["Set-Cookie"]
		if cookies == nil {
			fmt.Println("网络异常，换个环境重试")
			// 返回异常信息
			return nil, fmt.Errorf("网络异常，换个环境重试")
		}

		// 处理cookies，只保留参数名和参数值
		for i, cookie := range cookies {
			cookie = strings.Split(cookie, ";")[0]
			cookies[i] = cookie
		}

		return cookies, nil
	} else {
		fmt.Println("登录失败")
		return nil, fmt.Errorf("登录失败")
	}
}
