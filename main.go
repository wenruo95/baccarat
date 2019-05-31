package main

import (
	"log"
	"os"
)

func main() {
	log.Printf("start process! pid:%d", os.Getpid())
	//
	for i := 0; i < 1; i++ {
		log.Printf("start new collections")
		collection := NewCollections(DEFAULT_PLAY_CNT)
		collection.Run()
		log.Printf("end a collections")
	}
	log.Printf("end process! pid:%v", os.Getpid())
}
