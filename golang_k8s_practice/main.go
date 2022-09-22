package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const port = "8000"

func main() {

	router := gin.New()

	defer func() {
		server := fmt.Sprintf("服务对外端口:%d",port)
		fmt.Println(server)
		_ = router.Run(port)
	}()

	_ = router.SetTrustedProxies(nil)

	register(router)

}
