package China

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

func Save2file(idx int, name []string) {
	path := "./答案生成/" + "古诗文第" + strconv.Itoa(idx) + "页"

	f, err := os.Create(path)

	if err != nil {
		fmt.Println("os err = ", err)
		return
	}
	defer f.Close()
	n := len(name)

	for i := 1; i < n; i++ {
		//	f.WriteString(fileName[i][1] + "\n")

		f.WriteString(name[i] + "\n")
	}
}

func SpiderPageDB(idx int, page chan int) {
	url := "http://www.docx88.com/wkid-8c7e93fcfab069dc50220113-" + strconv.Itoa(idx) + ".html"
	fmt.Println(url)
	result, err := HttpGetDB(url)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	ret1 := regexp.MustCompile(`<p class="txt">(.*?)</p>`)
	fileName := ret1.FindAllStringSubmatch(result, -1)
	// 爬取古诗文
	n := len(fileName)
	name := make([]string, 0)
	for i := 1; i < n; i++ {
		name = append(name, fileName[i][1])
	}

	Save2file(idx, name)
	page <- idx
}

func ToWork(start, end int) {
	fmt.Printf("正在爬去...")
	page := make(chan int)
	for i := start; i <= end; i++ {
		go SpiderPageDB(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("%d", <-page)
	}
}
