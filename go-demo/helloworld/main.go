package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

// 初始化 logger
func init() {
	// 初始化日志输出位置
	filePath := "./log/" + time.Now().Format("20060102") + ".txt"
	logFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}

	// zap
	writeSyncer := zapcore.AddSync(logFile)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger = zap.New(core)
}

func main() {
	// 日志测试
	logger.Debug("running...")

	// http 服务
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 输出日志
		logger.Info("recv a request, url: " + html.EscapeString(r.URL.Path))

		// 响应
		fmt.Fprintf(w, "hello world, url: %q", html.EscapeString(r.URL.Path))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Error(err.Error())
	}
}
