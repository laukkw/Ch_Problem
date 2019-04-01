package main

import (
	//	"bufio"
	"fmt"
	//	"io"
	"math/rand"
	"os"
	"time"
)

func WriteFile(path string) {
	//新建文件

	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer f.Close()

	rand.Seed(time.Now().UnixNano()) //以当前系统时间当做种子参数

	//	我想把它存到数组里 那样方便我进行运算

	for i := 0; i < 100; i++ {
		fmt.Printf("<%d题> ", i+1)
		//		var s string
		//		var t string
		//		s = fmt.Sprintf("%d", rand.Intn(100))
		//		t = fmt.Sprintf("%d", rand.Intn(100))
		var a, b int
		a = rand.Intn(20)
		b = rand.Intn(10)
		add := a * b
		if a == 0 || b == 0 {
			i--
		} else {

			Start := fmt.Sprintf("<第%d题>  ", i+1)
			_, err = f.WriteString(Start)

			buf := fmt.Sprintf("%v /  %v = %v \n", add, a, b)
			_, err := f.WriteString(buf)

			if err != nil {
				fmt.Println("err = ", err)
			}
		}
		fmt.Printf("%v /  %v = %v \n", add, a, b)
		//fmt.Printf("%d + %v = \n", rand.Intn(100), rand.Intn(100),)
		//fmt.Println()
	}
}

//func main() {
//	path := "./作业生成/除法.txt"
//	WriteFile(path) // 写入函数
//}
