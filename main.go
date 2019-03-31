package main

import (
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

	Generate := gtk.NewButtonWithLabel("Generate test")
	Answer := gtk.NewButtonWithLabel("Show Answer")

	Generate.SetSizeRequest(140, 80) //设置按钮的大小
	Answer.SetSizeRequest(140, 80)

	//创建选项按钮 添加四个加减乘除

	MathButtom := gtk.NewButtonWithLabel("100以内")

	AddButtom := gtk.NewButtonWithLabel("加法")

	SubButtom := gtk.NewButtonWithLabel("减法")

	MulButtom := gtk.NewButtonWithLabel("乘法")

	DivButtom := gtk.NewButtonWithLabel("除法")

	//设置四个按钮的大小

	MathButtom.SetSizeRequest(100, 80)

	AddButtom.SetSizeRequest(100, 80)

	SubButtom.SetSizeRequest(100, 80)

	DivButtom.SetSizeRequest(100, 80)
	// 按钮添加到布局中
	layout.Put(Generate, 170, 220)
	layout.Put(Answer, 320, 220)

	layout.Put(MathButtom, 150, 180)
	//显示控件
	win.ShowAll()
	//主事件循环

	gtk.Main()

}
