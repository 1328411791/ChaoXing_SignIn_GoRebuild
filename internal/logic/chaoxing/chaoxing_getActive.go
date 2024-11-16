package chaoxing

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"sync"
)

func GetAllActive(courses []CourseType, cookies []string) (map[string]interface{}, error) {
	var activeList map[string]interface{}
	var mu sync.Mutex
	var wg sync.WaitGroup

	type result struct {
		courseId string
		actives  []ActiveType
		err      error
	}

	resultChannel := make(chan result, len(courses))

	for _, course := range courses {
		wg.Add(1)
		go func(course CourseType) {
			defer wg.Done()
			actives, err := GetActive(course, cookies)
			// 将结果发送到 channel
			resultChannel <- result{courseId: course.courseId, actives: actives, err: err}
		}(course)
	}

	go func() {
		wg.Wait()
		close(resultChannel) // 所有 goroutine 完成后关闭 channel
	}()

	// 处理所有查询结果
	for res := range resultChannel {
		if res.err == nil && len(res.actives) > 0 {
			mu.Lock() // 锁住共享资源
			activeList[res.courseId] = res.actives
			mu.Unlock() // 解锁
		}
	}

	return activeList, nil
}

func GetActive(course CourseType, cookies []string) ([]ActiveType, error) {
	// cookies参数只要_uid _d vc3 三个就可以了

	actives, err := doGetActive(course, cookies)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return actives, nil
}

/*
{
                "userStatus": 1,
                "nameTwo": "",
                "otherId": "2",
                "groupId": 2,
                "source": 15,
                "isLook": 1,
                "type": 2,
                "releaseNum": 0,
                "attendNum": 4,
                "activeType": 2,
                "logo": "https://mobilelearn.chaoxing.com/front/mobile/common/images/newActiveIcon80/active_type_2_gray.png?v=7",
                "nameOne": "签到",
                "startTime": 1731503204000,
                "id": 4000108634038,
                "endTime": 1731503504000,
                "status": 2,
                "nameFour": "11-13 21:11"
            },
*/

type ActiveType struct {
	UserStatus int    `json:"userStatus"`
	NameTwo    string `json:"nameTwo"`
	OtherId    string `json:"otherId"`
	GroupId    int    `json:"groupId"`
	Source     int    `json:"source"`
	IsLook     int    `json:"isLook"`
	Type       int    `json:"type"`
	ReleaseNum int    `json:"releaseNum"`
	AttendNum  int    `json:"attendNum"`
	ActiveType int    `json:"activeType"`
	Logo       string `json:"logo"`
	NameOne    string `json:"nameOne"`
	StartTime  int64  `json:"startTime"`
	ID         int64  `json:"id"`
	EndTime    int64  `json:"endTime"`
	Status     int    `json:"status"`
	NameFour   string `json:"nameFour"`
}

func doGetActive(course CourseType, cookies []string) ([]ActiveType, error) {
	payload := fmt.Sprintf("fid=0&courseId=%s&classId=%s&_=%d", course.courseId, course.clazzId, gtime.Now().Unix())

	body, err := getGet(ACTIVELIST.METHOD, ACTIVELIST.URL, payload, cookies)

	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	// 提取活动列表
	var actives []ActiveType

	if int(result["result"].(float64)) == 1 {
		data := result["data"].(map[string]interface{})
		activelist := data["activeList"].([]interface{})
		for _, v := range activelist {
			// 将v解析为ActiveType
			jsonStr, _ := json.Marshal(v)
			var active ActiveType
			_ = json.Unmarshal(jsonStr, &active)

			// 只添加在starttime和endtime之间的活动
			nowTime := gtime.Now().Unix()
			if active.StartTime <= nowTime && active.EndTime >= nowTime {
				actives = append(actives, active)
			}
		}
	}

	return actives, nil
}
