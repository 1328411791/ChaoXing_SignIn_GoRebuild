package service

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func Test_Trim(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		Chaoxing().GetCourses(3)
	})
}
