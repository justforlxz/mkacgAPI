package modules

import (
	"fmt"
	"net/http"

	"github.com/martini-contrib/render"
)

type NewsSource struct {
	Title  string
	Source string
	Img    string
	Next   string
	Url    string
}

func News(r render.Render, w http.ResponseWriter, request *http.Request) {
	//规定参数顺序
	/*
	   1.type 只能是json
	   2.page
	*/
	if request.FormValue("type") == "json" {
		if len(request.FormValue("p")) == 0 {
			T1, T2 := News_index("http://acg.17173.com")
			r.JSON(200, map[string]interface{}{"news": T1, "page": T2})
		} else if request.FormValue("p") == "begin" || request.FormValue("p") == "end" {
			T1, T2 := News_index("http://acg.17173.com")
			r.JSON(200, map[string]interface{}{"news": T1, "page": T2})
		} else {
			var Page = "http://acg.17173.com/index_" + request.FormValue("p") + ".shtml"
			T1, T2 := News_index(Page)
			r.JSON(200, map[string]interface{}{"news": T1, "page": T2})
		}
	} else {
		r.Text(200, "请输入正确的参数")
	}

}
func Bangumi(cs ...string) {
	for _, num := range cs {
		fmt.Println(num)
	}
}
