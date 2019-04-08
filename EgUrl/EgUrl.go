package Url

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

func Save2file(idx int, englishname [][]string) {
	path := "/home/rzry/桌面/" + "第 " + strconv.Itoa(idx) + " 页.txt"
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("os err = ", err)
		return
	}
	defer f.Close()
	n := len(englishname)
	for i := 0; i < n; i++ {
		NewUrl := "https://www.koolearn.com/" + englishname[i][1] + ".html"

		f.WriteString(NewUrl)
		f.WriteString("\n")

	}

}
func SpiderPageDB(idx int, page chan int) {
	url := "https://www.koolearn.com/dict/tag_1395_" + strconv.Itoa(idx) + ".html"

	fmt.Println(url)
	result, err := HttpGetDB(url)

	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	ret1 := regexp.MustCompile(`<a class="word" href="(.*?).html`)

	englishname := ret1.FindAllStringSubmatch(result, -1)

	Save2file(idx, englishname)

	page <- idx
}

func ToWork(start, end int) {
	fmt.Printf("正在爬取...")
	page := make(chan int)

	for i := start; i <= end; i++ {
		go SpiderPageDB(i, page)

	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d 爬取完毕\n", <-page)
	}

}
