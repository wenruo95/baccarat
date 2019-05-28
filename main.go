package main

import (
	"log"
)

func test_fished_yates() {
	arrs := make([]int32, 0)
	for i := 0; i < 54; i++ {
		arrs = append(arrs, int32(i+1))
	}
	after := FishedYates(arrs)
	log.Printf("after[%d]:%v\n", len(after), after)
}

func test_knuth_durtenfeld() {
	arrs := make([]int32, 0)
	for i := 0; i < 54; i++ {
		arrs = append(arrs, int32(i+1))
	}
	log.Printf("after2[%d]:%v\n", len(arrs), arrs)
	after2 := KnuthDurstenfeld(arrs)
	log.Printf("after2[%d]:%v\n", len(after2), after2)
}

func test_bench_fished_yates() {
}

func test_bench_knuth_durtenfeld() {
}

func main() {
	test_fished_yates()
	test_fished_yates()
}
