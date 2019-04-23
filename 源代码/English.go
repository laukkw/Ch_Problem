package main

import (
	"fmt"
	//	"io"
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
func SpiderPageDB2(url string) (title, connect, adj string, err error) {
	fmt.Println(url)
	result, err1 := HttpGet(url)
	if err1 != nil {
		err = err1
		return
	}
	ret1 := regexp.MustCompile(`<h1 class="word-spell">(.*?)</h1>`)
	if ret1 == nil {
		fmt.Println("re err = ", err)
		return
	}
	tmpContent := ret1.FindAllStringSubmatch(result, -1)
	for _, data := range tmpContent {
		connect = data[1]

		title = strings.Replace(title, "\n", "", -1)
		title = strings.Replace(title, "\r", "", -1)
		title = strings.Replace(title, " ", "", -1)
		title = strings.Replace(title, "\t", "", -1)
		break
	}

	// 解析单词词性

	ret2 := regexp.MustCompile(`<span class="prop">(.*?)</span>`)
	tmpAdj := ret2.FindAllStringSubmatch(result, -1)
	for _, data := range tmpAdj {
		adj = data[1]

		title = strings.Replace(title, "\n", "", -1)
		title = strings.Replace(title, "\r", "", -1)
		title = strings.Replace(title, " ", "", -1)
		title = strings.Replace(title, "\t", "", -1)
		break
	}

	//解析单词意思

	ret3 := regexp.MustCompile(`<span>(.*?)</span>`)
	tmpTitle := ret3.FindAllStringSubmatch(result, -1)
	for _, data := range tmpTitle {
		title = data[1]

		title = strings.Replace(title, "\n", "", -1)
		title = strings.Replace(title, "\r", "", -1)
		title = strings.Replace(title, " ", "", -1)
		title = strings.Replace(title, "\t", "", -1)
		break
	}
	return

}

func storeWorldsTOFile(i int, fileContent []string, fileTitle []string, fileAdj []string) {
	f, err := os.Create(strconv.Itoa(i) + ".txt")
	if err != nil {
		fmt.Println("err == ", err)
		return
	}
	defer f.Close()
	n := len(fileTitle)
	for i := 0; i < n; i++ {
		f.WriteString(fileTitle[i] + "\t\t\t")
		f.WriteString(fileContent[i] + "\t\t\t")
		f.WriteString(fileAdj[i] + "\n")
	}
}

func SpiderPageDB(i int, page chan int) {

	url := "https://www.koolearn.com/dict/tag_1395_" + strconv.Itoa(i) + ".html"

	fmt.Println(url)
	result, err := HttpGet(url)

	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	re := regexp.MustCompile(`<a class="word" href="(.*?)">`)
	if re == nil {
		fmt.Println("re err == ", err)
		return
	}
	NewUrl := re.FindAllStringSubmatch(result, -1)
	//	NewUrl := "https://www.koolearn.com/" + Url2

	fileTitle := make([]string, 0)

	fileContent := make([]string, 0)

	fileAdj := make([]string, 0)

	for _, data := range NewUrl {
		title, content, adj, err := SpiderPageDB2("https://www.koolearn.com/" + data[1])

		if err != nil {
			fmt.Println("err = ", err)
			continue
		}
		fileContent = append(fileContent, content)
		fileTitle = append(fileTitle, title)
		fileAdj = append(fileAdj, adj)

	}
	storeWorldsTOFile(i, fileTitle, fileContent, fileAdj)
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

	var start, end int
	fmt.Println("请输入爬去的起始页 >= 1 ")
	fmt.Scan(&start)
	fmt.Println("请输入爬去的末尾页 >= 1 ")
	fmt.Scan(&end)

	//实际的工作

	ToWork(start, end)

}
