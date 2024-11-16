package chaoxing

import (
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
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

	req.Header.Add("Cookie", strings.Join(cookies, ";"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
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

	if resp.StatusCode == 302 {
		return nil, gerror.NewCode(gcode.CodeNotImplemented)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
