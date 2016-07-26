package modules

import "fmt"

func News(cs ...string) {
	//规定参数顺序
	/*
	   1.type 只能是json
	   2.page
	*/
	if cs[0] == "json" {
		News_index("s")
	}
}
func Bangumi(cs ...string) {
	for _, num := range cs {
		fmt.Println(num)
	}
}
