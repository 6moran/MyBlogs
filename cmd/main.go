package main

import (
	"MyBlogs/internal/app"
	"log"
)

func main() {
	a := app.New()

	if err := a.Initialize(); err != nil {
		log.Fatalf("服务器初始化失败: %v\n", err)
		return
	}

	a.Run()
}
