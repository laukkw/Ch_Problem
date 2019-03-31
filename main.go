package main

import (
	"fmt"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func HandleButton(ctx *glib.CallbackContext) {
	arg := ctx.Data()   //获取用户传递的参数，是空接口类型
	p, ok := arg.(*int) //类型断言
	if ok {             //如果ok为true，说明类型断言正确
		//fmt.Println("*p = ", *p) //用户传递传递的参数为&tmp，是一个变量的地址
		*p = 250 //操作指针所指向的内存
	}

	fmt.Println("按钮gen被按下")

	//gtk.MainQuit() //关闭gtk程序
}

func main() {
	// 初始化
	gtk.Init(&os.Args)
	//创建主窗口
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	// 创建 窗x口属性
	win.SetTitle("给我的弟弟妹妹们")
	win.SetSizeRequest(480, 320)

	// 创建容器控件
	layout := gtk.NewFixed()

	// 布局添加到窗口上

	win.Add(layout)
	//创建按钮运行按钮

	Generate := gtk.NewButtonWithLabel("Generate test")
	Answer := gtk.NewButtonWithLabel("Show Answer")

	Generate.SetSizeRequest(140, 80) //设置按钮的大小
	Answer.SetSizeRequest(140, 80)

	//创建选项按钮 添加四个加减乘除

	//MathButtom := gtk.NewButtonWithLabel("100以内")

	AddButtom := gtk.NewButtonWithLabel("加法")

	SubButtom := gtk.NewButtonWithLabel("减法")

	MulButtom := gtk.NewButtonWithLabel("乘法")

	DivButtom := gtk.NewButtonWithLabel("除法")

	//设置四个按钮的大小

	MulButtom.SetSizeRequest(60, 40)

	AddButtom.SetSizeRequest(60, 40)

	SubButtom.SetSizeRequest(60, 40)

	DivButtom.SetSizeRequest(60, 40)
	// 按钮添加到布局中
	layout.Put(Generate, 170, 220)
	layout.Put(Answer, 320, 220)

	layout.Put(MulButtom, 30, 100)
	layout.Put(AddButtom, 30, 60)
	layout.Put(SubButtom, 30, 140)
	layout.Put(DivButtom, 30, 180)

	//信号处理

	Answer.Connect("pressed", func() {

		fmt.Println("答案按钮被按下")

	})

	//信号处理2

	tmp := 10
	Generate.Connect("pressed", HandleButton, &tmp)
	//显示控件
	win.ShowAll()
	//主事件循环

	gtk.Main()

}
