package service

import (
	"testing"
)

func Test_FishedYates(t *testing.T) {
	arrs := make([]int32, 0)
	for i := 0; i < 54; i++ {
		arrs = append(arrs, int32(i+1))
	}
	after := FishedYates(arrs)
	t.Logf("after1[%d]:%v\n", len(after), after)
}

func Test_KnuthDurstenfeld(t *testing.T) {
	arrs := make([]int32, 0)
	for i := 0; i < 54; i++ {
		arrs = append(arrs, int32(i+1))
	}
	after := KnuthDurstenfeld(arrs)
	t.Logf("after2[%d]:%v\n", len(after), after)
}

func Benchmark_fished_yates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrs := make([]int32, 0)
		for i := 0; i < 54; i++ {
			arrs = append(arrs, int32(i+1))
		}
		_ = FishedYates(arrs)
	}
}

func Benchmark_knuth_durtenfeld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrs := make([]int32, 0)
		for i := 0; i < 54; i++ {
			arrs = append(arrs, int32(i+1))
		}
		_ = KnuthDurstenfeld(arrs)
	}
}

/*
// card出现在每个位置的次数
func Test_Ratio_fished_yates() {
	count := 10000
	mp := make(map[int][]int) // card - pos - int
	for cnt := 0; cnt < count; cnt++ {
		arrs := make([]int32, 0)
		for i := 0; i < 54; i++ {
			arrs = append(arrs, int32(i+1))
		}
		after := FishedYates(arrs)
	}
}

func Test_Ratio_knuth_durtenfeld() {
}
*/
