package chaoxing

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"strings"
)

func GetActive(course CourseType, cookies []string) ([]CourseType, error) {
	// cookies参数只要_uid _d vc3 三个就可以了
	var selectCookie []string

	// 遍历cookies
	for _, cookie := range cookies {
		// 如果cookie中包含_uid _d vc3 就添加到selectCookie中
		if strings.Contains(cookie, "_uid") || strings.Contains(cookie, "uf") || strings.Contains(cookie, "_d") || strings.Contains(cookie, "vc3") {
			selectCookie = append(selectCookie, cookie)
		}
	}

	coursesByte, err := doGetActive(course, selectCookie)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 解析html 获取courseId,clazzId，personId id
	courses, err := extractCourses(coursesByte)

	return courses, nil
}

func doGetActive(course CourseType, cookies []string) ([]byte, error) {
	payload := fmt.Sprintf("fid=0&courseId=%s&classId=%s&_=%d", course.courseId, course.clazzId, gtime.Now().Unix())

	body, err := getGet(ACTIVELIST.METHOD, ACTIVELIST.URL, strings.NewReader(payload), cookies)

	if err != nil {
		return nil, err
	}

	return body, nil
}
