package chaoxing

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetCourses(cookies []string) error {
	// cookies参数只要_uid _d vc3 三个就可以了
	var selectCookie []string

	// 遍历cookies
	for _, cookie := range cookies {
		// 如果cookie中包含_uid _d vc3 就添加到selectCookie中
		if strings.Contains(cookie, "_uid") || strings.Contains(cookie, "_d") || strings.Contains(cookie, "vc3") {
			selectCookie = append(selectCookie, cookie)
		}
	}

	doGetCourses(selectCookie)
	return nil
}

type CourseType struct {
	courseId string
	classId  string
}

func doGetCourses(cookies []string) (*[]CourseType, error) {
	formData := fmt.Sprintf("courseType=1&courseFolderId=0&courseFolderSize=0")
	req, err := http.NewRequest(COURSELIST.METHOD, COURSELIST.URL, strings.NewReader(formData))
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

	var result string = string(body)

	fmt.Println(result)

	return nil, nil
}
