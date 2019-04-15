package main

import (
	"Ch_Problem/EgUrl"
	"Ch_Problem/Math"
	"fmt"
	//	"github.com/mattn/go-gtk/glib"
	"Ch_Problem/China"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

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

	Generate := gtk.NewButtonWithLabel("古诗文必备")
	Answer := gtk.NewButtonWithLabel("英语3000")

	Generate.SetSizeRequest(140, 40) //设置按钮的大小
	Answer.SetSizeRequest(140, 40)

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
	layout.Put(Generate, 30, 260)
	layout.Put(Answer, 30, 220)

	layout.Put(MulButtom, 30, 100)
	layout.Put(AddButtom, 30, 60)
	layout.Put(SubButtom, 30, 140)
	layout.Put(DivButtom, 30, 20)

	//信号处理 加法

	AddButtom.Connect("pressed", func() {
		path := "./作业生成/加法.txt"
		Add.WriteFileAdd(path)

	})

	//	信号处理 减法
	SubButtom.Connect("pressed", func() {
		path := "./作业生成/减法.txt"
		Add.WriteFileSub(path)

	})

	//信号处理 乘法
	MulButtom.Connect("pressed", func() {
		path := "./作业生成/乘法.txt"
		Add.WriteFileMul(path)

	})

	//信号处理  除法
	DivButtom.Connect("pressed", func() {
		path := "./作业生成/除法.txt"
		Add.WriteFileDiv(path)

	})

	Answer.Connect("pressed", func() {
		//======================获取到 url
		var start, end int

		fmt.Print("输出起始页")

		fmt.Scan(&start)

		fmt.Print("输入终止")

		fmt.Scan(&end)
		Url.ToWork(start, end)

	})

	//信号处理2
	Generate.Connect("pressed", func() {
		//选择爬取的页数
		var start, end int

		fmt.Print("输入您想爬取的起始页(共7页):")
		fmt.Scan(&start)
		fmt.Print("输入爬取的终止页(共7页):")

		fmt.Scan(&end)

		China.ToWork(start, end)

	})
	//显示控件
	win.ShowAll()
	//主事件循环

	gtk.Main()

}
