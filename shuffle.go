package main

import (
	"math/rand"
	"time"
)

// Fisher-Yates Shuffle算法
// 1. 每次从扑克从随机一张牌出来，知道随即完所有的牌
func FishedYates(arrs []int32) []int32 {
	length := len(arrs)
	data := make([]int32, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		index := rand.Int() % (length - i)
		data = append(data, arrs[index])
		arrs[index] = arrs[length-i-1]
	}
	return data
}

// Knuth-Durstenfeld Shuffle
// 1. 从头至尾每次随意一张牌和当前牌交换，直到最后一张牌
func KnuthDurstenfeld(arrs []int32) []int32 {
	length := len(arrs)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		index := rand.Int() % length
		arrs[index], arrs[i] = arrs[i], arrs[index]
	}
	return arrs
}

// 0 ~ 8
func CardToPoint(card int32) int32 {
	card = card % 13
	if card > 8 {
		return 0
	}
	return card + 1
}

// 0~12
func CalculateCardsAndPair(cards []int32) (int32, bool) {
	sum, pair := int32(0), false
	mp := make(map[int32]bool)
	for _, card := range cards {
		validCard := card % 13
		sum = sum + CardToPoint(validCard)
		if _, ok := mp[validCard]; ok {
			pair = true
		} else {
			mp[validCard] = true
		}
	}
	return sum % 10, pair
}

func CalculateCards(cards []int32) int32 {
	sum := int32(0)
	for _, card := range cards {
		sum = sum + CardToPoint(card)
	}
	return sum % 10
}
