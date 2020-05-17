package main

import (
	master "Harbor/corntab/main"
	"fmt"
	"runtime"
)

func InitArgs()  {
	
}

func InitEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}	

func main() {
	var (
		err error
	)

	// 初始化线程数
	InitEnv()

	// 加载配置
	if err = master.InitConfig(); err != nil {
		goto ERR
	}

	// 初始化ApiServer
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}

	// 正常退出
	return
ERR:
	fmt.Println(err)
}
