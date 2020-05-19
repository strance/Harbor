package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// CTRL+ALT+V 自动生成表达式返回值
func main()  {
	requestWithSelfClient()
}

func defaultRequest()  {
	url := "https://www.baidu.com/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	response, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("response:", string(response))
}

// 在请求时带上请求头，模拟手机浏览器
func requestWhitHeaders()  {
	url := "http://www.baidu.com/"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("user-agent","Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	bytes, err := httputil.DumpResponse(response, true)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(string(bytes))
}

// 使用自定义客户端发送请求
func requestWithSelfClient()  {
	url := "http://www.baidu.com/"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("user-agent","Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	client := http.Client{
		Transport:     nil, // 设置代理服务器
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req.URL)
			return nil
		}, // 重定向函数, via存放所有重定向的路径，req是重定向的目标
		Jar:           nil, // cookie
		Timeout:       0,  
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	bytes, err := httputil.DumpResponse(response, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}
