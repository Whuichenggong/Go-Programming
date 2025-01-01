package main

import (
	"fmt"
	"log"
	"net/http"

	"gopl.io/ch1/lissajous"
)

// !+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func main() {
	// 定义 handler 函数，处理 HTTP 请求
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 设置响应头为 GIF 图片格式
		w.Header().Set("Content-Type", "image/gif")
		// 调用 lissajous 函数生成 GIF 动画并发送到浏览器
		lissajous.Lissajous(w)
	})

	// 启动 HTTP 服务器，监听 localhost:8000
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
