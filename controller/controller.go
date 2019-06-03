package controller

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/wenruo95/baccarat/service"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (this *Controller) Serve() {
	for i := 0; i < 10000; i++ {
		collection := service.NewCollections(service.DEFAULT_PLAY_CNT)
		collection.Run()
		collection.PrintPaint()
		fmt.Println()
	}
	//this.signHandle()
}

func (this *Controller) Close(reason string) {
	log.Printf("baccarat closed! reason:%v", reason)
}

func (this *Controller) signHandle() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
	select {
	case sig := <-signalChan:
		this.Close(fmt.Sprintf("SIG-%v[CTRL+C]", sig))
	}
	os.Exit(0)
}
