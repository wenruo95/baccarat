package main

import (
	"log"
	"os"
)

func main() {
	log.Printf("start process! pid:%d", os.Getpid())
}
