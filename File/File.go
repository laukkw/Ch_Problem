package File

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//	"strings"
	//"strconv"
)

/*
func writefile(path2 string, buf string) {
	path2 = "./作业生成/第4页"
	f2, err2 := os.Open(path2)
	if err2 != nil {
		fmt.Println("os err = ", err2)
		return
	}
	defer f2.Close()
	//	n := len(buf)
	//	for i := 0; i < n; i++ {
	//	f2.WriteString(buf)

	n, _ := f2.Seek(0, os.SEEK_END)

	_, err2 = f2.WriteAt([]byte(buf), n)

	//	}

	//	fmt.Printf("buf  = %s", string(buf))

}
*/
func saveFile(idx int, name []string) {
	path2 := "./作业生成/古诗文作业"
	f2, err2 := os.Create(path2)
	if err2 != nil {
		fmt.Println("os err = ", err2)
		return
	}
	defer f2.Close()
	for i := 0; i < 200; i++ {
		f3, err3 := f2.WriteString(name[i] + "\n\n\n")
		f2.WriteString("___________________________" + "\n")
		f2.WriteString("\n" + "\n")
		if err3 != nil {
			fmt.Println(f3)
			fmt.Println("err3 ==", err3)
			return
		}
	}

}
func Readfile(idx int, path string) {
	//打开文件

	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return

	}
	defer f.Close()

	r := bufio.NewReader(f)
	name := make([]string, 0)
	//	var name string
	for i := 0; i < 200; i++ {
		buf, err := r.ReadString(',')
		if err != nil {
			if err == io.EOF {
				//	fmt.Println("err ==", err)
				break
			}
			fmt.Println("err = ", err)

		}
		name = append(name, buf)
		fmt.Println(buf)

	}
	saveFile(idx, name)

}
