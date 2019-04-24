package main

import (
	//	"io"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url) //发送  get请求
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	//读网页内容
	buf := make([]byte, 4*1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf[:n]) //累加读取数据
	}
	return
}
func SpiderPageDB2(url string) (name string, err error) {
	fmt.Println(url)
	result, err1 := HttpGet(url)
	if err1 != nil {
		err = err1
		return
	}
	ret1 := regexp.MustCompile(`<p class="txt">(.*?)</p>`)
	if ret1 == nil {
		fmt.Println("ret1 err = ", err)
		return
	}

	FileName := ret1.FindAllStringSubmatch(result, -1)
	for _, data := range FileName {
		name = data[1]
		name = strings.Replace(name, "\n", "", -1)
		break
	}
	return
}

func storeChinaToFile(i int, fileName []string) {
	path := strconv.Itoa(i) + "页.txt"
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close()
	n := len(fileName)
	for i := 1; i < n; i++ {
		f.WriteString(fileName[i][1] + "\n")
	}
}
func SpiderPageDB(i int, page chan int) {
	Url := strconv.Itoa(i) + ".html"
	fileName := make([]string, 0)
	for _, data := range Url {
		name, err := SpiderPageDB2(" http://www.docx88.com/wkid-8c7e93fcfab069dc50220113-" + data[i])
		if err != nil {
			fmt.Println("err =", err)

			fileName = append(fileName, name)
		}
	}
	storeChinaToFile(i, fileName)
	page <- i
}

func ToWork(start, end int) {

	fmt.Printf("正在爬取...")
	page := make(chan int)

	for i := start; i <= end; i++ {
		go SpiderPageDB(i, page)

	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d爬取完毕\n", <-page)

	}
}

func main() {
	//选择爬取的页数
	var start, end int

	fmt.Print("输入您想爬取的起始页(共7页):")
	fmt.Scan(&start)
	fmt.Print("输入爬取的终止页(共7页):")

	fmt.Scan(&end)

	ToWork(start, end)
}
