package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//	"strconv"
)

func readfile(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer f.Close()
	r := bufio.NewReader(f)

	for {
		buf, err := r.ReadString(',')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err = ", err)
		}
		fmt.Println(buf)

		path2 := "./作业生成/第2页"
		f2, err2 := os.Create(path2)
		if err2 != nil {
			fmt.Println("os err = ", err2)
			return
		}
		defer f2.Close()
		f2.WriteString(buf)

		//	fmt.Printf("buf  = %s", string(buf))
	}

}
func main() {
	path := "./作业生成/第2页"
	readfile(path)
}
