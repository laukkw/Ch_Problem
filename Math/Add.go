package Add

import (
	//	"bufio"
	"fmt"
	//	"io"
	"math/rand"
	"os"
	"time"
)

func WriteFileAdd(path string) {
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
		//	fmt.Printf("<%d题> ", i+1)
		//		var s string
		//		var t string
		//		s = fmt.Sprintf("%d", rand.Intn(100))
		//		t = fmt.Sprintf("%d", rand.Intn(100))
		var a, b int
		a = rand.Intn(100)
		b = rand.Intn(100)
		add := a + b

		//		fmt.Printf("%v + %v = %v \n", a, b, add)

		if i%2 == 0 {
			Start := fmt.Sprintf("<第%d题>  ", i+1)

			buf := fmt.Sprintf("%v + %v = ____ |||||| answer =  %v\n", a, b, add)
			//	Answer := fmt.Sprintf("%v  + %v = %v \n", a, b, add)

			_, err = f.WriteString(Start)
			_, err := f.WriteString(buf)
			//	_, err = f.WriteString(Answer)
			if err != nil {
				fmt.Println("err = ", err)
			}
			//	fmt.Printf("\n")
		}

		//fmt.Printf("%d + %v = \n", rand.Intn(100), rand.Intn(100),)
		//fmt.Println()
	}
	fmt.Println("作业已经生成 前往./作业生成/ 查看")

}

func WriteFileSub(path string) {
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
		//	fmt.Printf("<%d题> ", i+1)
		//		var s string
		//		var t string
		//		s = fmt.Sprintf("%d", rand.Intn(100))
		//		t = fmt.Sprintf("%d", rand.Intn(100))

		var a, b, Sub int
		a = rand.Intn(100)
		b = rand.Intn(100)
		if a > b {
			Sub = a - b
		} else {
			a, b = b, a
			Sub = a - b
		}

		Start := fmt.Sprintf("<第%d题>  ", i+1)
		_, err = f.WriteString(Start)

		buf := fmt.Sprintf("%v - %v = _____ ||||| answer = %v \n", a, b, Sub)
		_, err := f.WriteString(buf)

		if err != nil {
			fmt.Println("err = ", err)
		}

		//	fmt.Printf("%v - %v = %v \n", a, b, Sub)
		//fmt.Printf("%d + %v = \n", rand.Intn(100), rand.Intn(100),)
		//fmt.Println()

	}
	fmt.Println("作业已经生成 前往./作业生成/ 查看")

}
func WriteFileDiv(path string) {
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
		//fmt.Printf("<%d题> ", i+1)
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

			buf := fmt.Sprintf("%v /  %v = ____ ||||| answer = %v \n", add, a, b)
			_, err := f.WriteString(buf)

			if err != nil {
				fmt.Println("err = ", err)
			}
		}
		//	fmt.Printf("%v /  %v = %v \n", add, a, b)
		//fmt.Printf("%d + %v = \n", rand.Intn(100), rand.Intn(100),)
		//fmt.Println()

	}
	fmt.Println("作业已经生成 前往./作业生成/ 查看")
}
func WriteFileMul(path string) {
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
		//	fmt.Printf("<%d题> ", i+1)
		//		var s string
		//		var t string
		//		s = fmt.Sprintf("%d", rand.Intn(100))
		//		t = fmt.Sprintf("%d", rand.Intn(100))
		var a, b int
		a = rand.Intn(20)
		b = rand.Intn(10)
		add := a * b

		Start := fmt.Sprintf("<第%d题>  ", i+1)
		_, err = f.WriteString(Start)

		buf := fmt.Sprintf("%v *  %v = _____ |||| answer =  %v \n", a, b, add)
		_, err := f.WriteString(buf)

		if err != nil {
			fmt.Println("err = ", err)
		}

		//	fmt.Printf("%v *  %v = %v \n", a, b, add)
		//fmt.Printf("%d + %v = \n", rand.Intn(100), rand.Intn(100),)
		//fmt.Println()

	}
	fmt.Println("作业已经生成 前往./作业生成/ 查看")

}

//func main() {
//	path := "./作业生成/加法.txt"
//	WriteFile(path) // 写入函数
//}
