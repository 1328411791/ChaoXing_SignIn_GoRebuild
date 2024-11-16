package chaoxing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetCourses(cookies []string) error {
	doGetCourses(cookies)
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
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
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
		/*
			data := result["data"]
			arr := make([]CourseType, 0)
			var end_of_courseid int
			for i := 1; ; i++ {
				i = strings.Index(data, "course_", i)
				if i == -1 {
					break
				}
				end_of_courseid = strings.Index(data, "_", i+7)
				arr = append(arr, CourseType{
					courseId: data[i+7 : end_of_courseid],
					classId:  data[end_of_courseid+1 : strings.Index(data, `"`, i+1)],
				})
			}

		*/

		return nil, nil
	} else {
		fmt.Println("登录失败")
		return nil, fmt.Errorf("登录失败")
	}
}
