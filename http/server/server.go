package main

import (
	"fmt"
	"net/http"

	_ "net/http/pprof"
)

func requestHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("server running..."))
}

func main()  {
	http.HandleFunc("/", requestHandler)
	// 启动后，访问 localhost:8888/debug/pprof 查看信息
	// 查看30s CPU使用率：go tool pprof http://localhost:8888/debug/pprof/profile
	// 等结果出来以后，输入web,就可以查看流程图
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
