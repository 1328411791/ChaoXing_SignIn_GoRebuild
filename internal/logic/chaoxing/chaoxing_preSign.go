package chaoxing

import (
	"fmt"
	"strings"
)

func PreSign(course CourseType, active ActiveType, cookies []string) ([]CourseType, error) {
	// cookies参数只要_uid _d vc3 三个就可以了
	var _uid string

	// 遍历cookies
	for _, cookie := range cookies {
		// 如果cookie中包含_uid _d vc3 就添加到selectCookie中
		if strings.Contains(cookie, "_uid") {
			_uid = strings.Split(cookie, "=")[1]
		}
	}

	coursesByte, err := doPreSign(course, active, _uid, cookies)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 解析html 获取courseId,clazzId，personId id
	courses, err := extractCourses(coursesByte)

	return courses, nil
}

func doPreSign(course CourseType, active ActiveType, _uid string, cookies []string) ([]byte, error) {
	payload := fmt.Sprintf("courseId=%s&classId=%s&activePrimaryId=%d&general=1&sys=1&ls=1&appType=15&&tid=&uid=%s&ut=s", course.courseId, course.clazzId, active.ID, _uid)

	body, err := getGet(PRESIGN.METHOD, PRESIGN.URL, payload, cookies)

	if err != nil {
		return nil, err
	}

	return body, nil
}
