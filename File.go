package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//	"strings"
	//	"strconv"
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

func appendToFile(path2 string, newbuf string) (err error) {
	f, err := os.Create(path2)
	if err != nil {
		fmt.Println("err ==", err)
	} else {
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(newbuf), n)

	}
	defer f.Close()
	return err
}

func readfile(path string) {
	//打开文件

	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return

	}
	defer f.Close()

	r := bufio.NewReader(f)
	for i := 0; i < 200; i++ {
		buf, err := r.ReadString(',')
		if err != nil {
			if err == io.EOF {
				//	fmt.Println("err ==", err)
				break
			}
			fmt.Println("err = ", err)
		}
		newbuf := buf
		fmt.Println(newbuf)
		path2 := "./作业生成/第4页"
		appendToFile(path2, newbuf)

	}

}
func main() {
	path := "./作业生成/第2页"

	readfile(path)

}
