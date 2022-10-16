package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

const port = ":8080"

func main() {

	router := gin.New()

	// 优雅退出
	stopC := make(chan os.Signal, 1)
	signal.Notify(stopC, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-stopC
		fmt.Println("准备优雅退出")
		// 这里可以准备清理资源。。。
		if value, ok := s.(syscall.Signal); ok {
			os.Exit(int(value))
		} else {
			os.Exit(0)
		}
	}()


	// 启动server
	defer func() {
		server := fmt.Sprintf("服务对外端口%v",port)
		fmt.Println(server)
		_ = router.Run(port)
	}()


	//_ = router.SetTrustedProxies(nil)

	register(router)

}
