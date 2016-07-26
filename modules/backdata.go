package modules

import (
	"fmt"

	
"github.com/kirigayakazushin/mkacgAPI/modules"
)

func News(cs ...string) {
	//规定参数顺序
	/*
	   1.type 只能是json
	   2.page
	*/
	if cs[0] == "json" {
		modules.News_index("s")
	}
}
func Bangumi(cs ...string) {
	for _, num := range cs {
		fmt.Println(num)
	}
}
