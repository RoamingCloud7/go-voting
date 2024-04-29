package main

import (
	"voting-backend/router"
)

func main() {
	r := router.Router()

	r.Run(":8081") // 监听并在 0.0.0.0:8081 上启动服务
}
