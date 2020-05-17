package master

import (
	"net"
	"net/http"
	"time"
)

// 单例
var G_ApiServer *ApiServer


// 任务的http接口
type ApiServer struct {
	httpServer *http.Server
}

func handleJobSave(w http.ResponseWriter, r *http.Request) {

}

// 初始化服务
func InitApiServer() (err error) {
	var (
		mux        *http.ServeMux
		listener   net.Listener
		httpServer *http.Server
	)

	// 配置路由
	mux = http.NewServeMux()
	mux.HandleFunc("/job/save", handleJobSave)

	// 启动http监听
	if listener, err = net.Listen("tcp", ":8070"); err != nil {
		return
	}

	// 创建一个http服务
	httpServer = &http.Server{
		ReadTimeout:  5 * time.Microsecond,
		WriteTimeout: 5 * time.Millisecond,
		Handler:      mux,
	}

	// 赋值单例
	G_ApiServer = &ApiServer{httpServer:httpServer}

	// 启动服务端
	go httpServer.Serve(listener)
	return
}
