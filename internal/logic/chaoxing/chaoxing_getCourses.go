package chaoxing

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func GetCourses(cookies []string) ([]CourseType, error) {
	// cookies参数只要_uid _d vc3 三个就可以了
	var selectCookie []string

	// 遍历cookies
	for _, cookie := range cookies {
		// 如果cookie中包含_uid _d vc3 就添加到selectCookie中
		if strings.Contains(cookie, "_uid") || strings.Contains(cookie, "_d") || strings.Contains(cookie, "vc3") {
			selectCookie = append(selectCookie, cookie)
		}
	}

	coursesByte, err := doGetCourses(selectCookie)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 解析html 获取courseId,clazzId，personId id
	courses, err := extractCourses(coursesByte)

	return courses, nil
}

type CourseType struct {
	courseId string
	clazzId  string
	personId string
	id       string
}

func doGetCourses(cookies []string) ([]byte, error) {
	payload := strings.NewReader("courseType=1&courseFolderId=0&courseFolderSize=0")

	body, err := getPost(COURSELIST.METHOD, COURSELIST.URL, payload, cookies)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func extractCourses(htmlByte []byte) ([]CourseType, error) {

	// 解析HTML
	doc, err := html.Parse(strings.NewReader(string(htmlByte)))

	// 定义一个切片来存储提取的数据
	if err != nil {
		fmt.Println(err)
	}

	var courses []CourseType
	var parseNode func(*html.Node)
	parseNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "li" {
			// 输出节点的属性
			var course CourseType
			var flag = false
			// 查找属性
			for _, attr := range n.Attr {
				if attr.Key == "class" && attr.Val == "course clearfix" {
					flag = true
				}
				if attr.Key == "courseid" {
					course.courseId = attr.Val
				}
				if attr.Key == "clazzid" {
					course.clazzId = attr.Val
				}
				if attr.Key == "personid" {
					course.personId = attr.Val
				}
				if attr.Key == "id" {
					course.id = attr.Val
				}
			}
			if course.courseId != "" && course.clazzId != "" && flag == true {
				courses = append(courses, course)
			}
		}

		// 递归遍历子节点
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parseNode(c)
		}
	}

	// 从文档根节点开始解析
	parseNode(doc)
	return courses, nil
}
