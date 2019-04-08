package main

import (
	"Ch_Problem/EgUrl"
	"fmt"
)

func main() {
	//选择爬取的页数
	var start, end int

	fmt.Print("输入您想爬取的起始页(共7页):")
	fmt.Scan(&start)
	fmt.Print("输入爬取的终止页(共7页):")

	fmt.Scan(&end)

	Url.ToWork(start, end)
}
