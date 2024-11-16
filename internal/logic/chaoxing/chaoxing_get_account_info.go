package chaoxing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetAccountInfo(cookies []string) error {
	// cookies参数只要_uid _d vc3 三个就可以了
	var selectCookie []string

	// 遍历cookies
	for _, cookie := range cookies {
		// 如果cookie中包含_uid _d vc3 就添加到selectCookie中
		if strings.Contains(cookie, "_uid") || strings.Contains(cookie, "_d") || strings.Contains(cookie, "vc3") {
			selectCookie = append(selectCookie, cookie)
		}
	}

	doGetAccountInfo(selectCookie)
	return nil
}

type AccountInfo struct {
	courseId string
	classId  string
}

func doGetAccountInfo(cookies []string) (*[]AccountInfo, error) {

	req, err := http.NewRequest(ACCOUNTMANAGE.METHOD, ACCOUNTMANAGE.URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8;")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Cookie", strings.Join(cookies, ";"))
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

		return nil, nil
	} else {
		fmt.Println("登录失败")
		return nil, fmt.Errorf("登录失败")
	}

	return nil, nil
}
