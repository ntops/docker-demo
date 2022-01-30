package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	// 初始化日志输出位置
	filePath := "./log/" + time.Now().Format("20060102") + ".txt"
	logFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.SetPrefix("[helloworld] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}

func main() {
	// 日志测试
	log.Println("helloworld")

	// http 服务器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world, url: %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
