package main

import (
	"fmt"
	"go-rbac-example/internal/initialize"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initialize.Init("./config.yaml")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	fmt.Println("退出")
}
