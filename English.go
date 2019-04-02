package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func HttpGetDB(url string) (result string, err error) {
	resp, err1 := http.Get(url)

	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()

	buf := make([]byte, 4096)

	//循环爬取整页的数据

	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}

		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}

func Save2file(idx int, fileName, fileAdj, fileMean [][]string) {
	path := "./作业生成/dssd.txt"

	f, err := os.Create(path)

	if err != nil {
		fmt.Println("os.Create err :", err)
		return
	}

	defer f.Close()

	n := len(fileName)

	f.WriteString("单词" + "\t\t\t" + "单词词性" + "\t\t" + "单词意思" + "\n")

	for i := 0; i < n; i++ {
		f.WriteString(fileName[i][1] + "\t\t\t" + fileAdj[i][1] + "\t\t" + fileMean[i][1] + "\n")

	}

}

func SpiderPageDB(idx int, page chan int) {
	//获取 url

	url := "https://www.koolearn.com/dict/tag_1395_" + strconv.Itoa(idx) + ".html"
	//	fmt.Println(url)
	//爬取 url 对应的  页面

	result1, err := HttpGetDB(url)
	if err != nil {
		fmt.Println("HttpGet err = ", err)
		return

	}

	ret4 := regexp.MustCompile(`<a class="word" href="/dict/wd_([/d]).html`)

	NewUrl := "https://www.koolearn.com/dict/wd_" + strconv.Itoa(*(ret4)) + ".html"
	result, err := HttpGetDB(NewUrl)
	if err != nil {
		fmt.Println("HttpGetDB err", err)
		return
	}

	// 解析单词
	ret1 := regexp.MustCompile(`<h1 class="word-spell">(*)</h1>`)

	fileName := ret1.FindAllStringSubmatch(result, -1)

	// 解析单词词性

	ret2 := regexp.MustCompile(`<span class="prop">(*)</span>`)
	fileAdj := ret2.FindAllStringSubmatch(result, -1)

	//解析单词意思

	ret3 := regexp.MustCompile(`<span>(*)</span>`)
	fileMean := ret3.FindAllStringSubmatch(result, -1)

	Save2file(idx, fileName, fileAdj, fileMean)

	page <- idx
}

func toWork(start, end int) {
	fmt.Printf("正在爬取%d 页到 %d 页 ...", start, end)

	page := make(chan int) //防止线程提前结束
	for i := start; i <= end; i++ {
		go SpiderPageDB(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第 %d 页爬取完毕\n", <-page)
	}

}

func main() {
	//选择爬取的页数
	var start, end int

	fmt.Print("输入您想爬取的起始页(共7页):")
	fmt.Scan(&start)
	fmt.Print("输入爬取的终止页(共7页):")

	fmt.Scan(&end)

	toWork(start, end)
}
