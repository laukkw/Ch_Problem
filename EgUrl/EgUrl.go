package Url

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

var ch = make(chan int) //定义一个全局的 channel

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
func SpiderPageDB2(idx int, NewUrl string, page2 chan int) {

	url := NewUrl
	result, err := HttpGetDB(url)
	if err != nil {
		fmt.Println("err ", err)
		return
	}
	// 解析单词
	ret1 := regexp.MustCompile(`<h1 class="word-spell">(.*?)</h1>`)

	fileName := ret1.FindAllStringSubmatch(result, -1)

	// 解析单词词性

	ret2 := regexp.MustCompile(`<span class="prop">(.*?)</span>`)
	fileAdj := ret2.FindAllStringSubmatch(result, -1)

	//解析单词意思

	ret3 := regexp.MustCompile(`<span>(.*?)</span>`)
	fileMean := ret3.FindAllStringSubmatch(result, -1)

	Save2file(idx, fileName, fileAdj, fileMean)

	page2 <- idx

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

/*
func Save2file(idx int, englishname [][]string) {
	path := "/home/rzry/桌面/" + "第 " + strconv.Itoa(idx) + " 页.txt"
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("os err = ", err)
		return
	}
	defer f.Close()
	n := len(englishname)
	page2 := make(chan int)
	for i := 0; i < n; i++ {
		NewUrl := "https://www.koolearn.com/" + englishname[i][1] + ".html"
		go SpiderPageDB2(NewUrl, page2)
		f.WriteString(NewUrl)
		f.WriteString("\n")
		fmt.Print("第 %d 页爬取完毕\n", <-page2)

	}

}
*/
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

	n := len(englishname)
	page2 := make(chan int)
	for i := 0; i < n; i++ {
		NewUrl := "https://www.koolearn.com/" + englishname[i][1] + ".html"
		go SpiderPageDB2(i, NewUrl, page2)
		//	fmt.Print(" %d ", <-page2)
	}

	//	Save2file(idx, englishname)
	page <- idx
	fmt.Println("第一次完了")
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
