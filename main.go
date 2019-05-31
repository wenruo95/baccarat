package main

import (
	"log"
	"os"

	"github.com/wenruo95/baccarat/controller"
)

func main() {
	log.Printf("start process! pid:%d", os.Getpid())
	controller.NewController().Serve()
}
