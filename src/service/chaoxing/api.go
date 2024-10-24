package chaoxing

type Api struct {
	URL    string
	METHOD string
}

var (
	LOGIN_PAGE = Api{
		URL:    "https://passport2.chaoxing.com/mlogin?fid=&newversion=true&refer=http%3A%2F%2Fi.chaoxing.com",
		METHOD: "GET",
	}
	LOGIN = Api{
		URL:    "https://passport2.chaoxing.com/fanyalogin",
		METHOD: "POST",
	}
	PRESIGN = Api{
		URL:    "https://mobilelearn.chaoxing.com/newsign/preSign",
		METHOD: "GET",
	}
	ANALYSIS = Api{
		URL:    "https://mobilelearn.chaoxing.com/pptSign/analysis",
		METHOD: "GET",
	}
	ANALYSIS2 = Api{
		URL:    "https://mobilelearn.chaoxing.com/pptSign/analysis2",
		METHOD: "GET",
	}
	PPTSIGN = Api{
		URL:    "https://mobilelearn.chaoxing.com/pptSign/stuSignajax",
		METHOD: "GET",
	}
	PPTACTIVEINFO = Api{
		URL:    "https://mobilelearn.chaoxing.com/v2/apis/active/getPPTActiveInfo",
		METHOD: "GET",
	}
	COURSELIST = Api{
		URL:    "https://mooc1-1.chaoxing.com/visit/courselistdata",
		METHOD: "POST",
	}
	BACKCLAZZDATA = Api{
		URL:    "https://mooc1-api.chaoxing.com/mycourse/backclazzdata",
		METHOD: "GET",
	}
	ACTIVELIST = Api{
		URL:    "https://mobilelearn.chaoxing.com/v2/apis/active/student/activelist",
		METHOD: "GET",
	}
	ACCOUNTMANAGE = Api{
		URL:    "https://passport2.chaoxing.com/mooc/accountManage",
		METHOD: "GET",
	}
	PANCHAOXING = Api{
		URL:    "https://pan-yz.chaoxing.com",
		METHOD: "GET",
	}
	PANLIST = Api{
		URL:    "https://pan-yz.chaoxing.com/opt/listres",
		METHOD: "POST",
	}
	PANTOKEN = Api{
		URL:    "https://pan-yz.chaoxing.com/api/token/uservalid",
		METHOD: "GET",
	}
	PANUPLOAD = Api{
		URL:    "https://pan-yz.chaoxing.com/upload",
		METHOD: "POST",
	}
	WEBIM = Api{
		URL:    "https://im.chaoxing.com/webim/me",
		METHOD: "GET",
	}
)
