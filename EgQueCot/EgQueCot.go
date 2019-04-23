package EgQueCot

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
func SpiderPageDB2(url string) (connect string, err error) {
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
	tmpConnect := ret1.FindAllStringSubmatch(result, -1)
	for _, data := range tmpConnect {
		connect = data[1]

		connect = strings.Replace(connect, "\n", "", -1)
		connect = strings.Replace(connect, "\r", "", -1)
		connect = strings.Replace(connect, " ", "", -1)
		connect = strings.Replace(connect, "\t", "", -1)
		break
	}

	return

}

func storeWorldsTOFile(i int, fileConnect []string) {
	path := "./作业生成/" + "英语3500知英议汉第" + strconv.Itoa(i) + "页"
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err == ", err)
		return
	}
	defer f.Close()
	n := len(fileConnect)
	f.WriteString("意思" + "\t\t\t" + "单词" + "\t\t\t" + "词性" + "\n")
	for i := 0; i < n; i++ {
		f.WriteString(fileConnect[i] + "\t\t\t")
		f.WriteString("_______________ ______________________________" + "\n")
		//	f.WriteString(fileContent[i] + "\t\t\t")
		//	f.WriteString(fileAdj[i] + "\n")
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

	fileConnect := make([]string, 0)

	for _, data := range NewUrl {
		connect, err := SpiderPageDB2("https://www.koolearn.com/" + data[1])

		if err != nil {
			fmt.Println("err = ", err)
			continue
		}
		fileConnect = append(fileConnect, connect)

	}
	storeWorldsTOFile(i, fileConnect)
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
