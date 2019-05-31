package main

import (
	"fmt"
	"log"
)

const (
	CARD_NUM            = 52 // 一副牌多少牌 没有大小王
	CARD_COLLECTION_NUM = 8  // 多少副牌
	DEFAULT_PLAY_CNT    = 55 // 默认回合数 (52 * 8) / 6 =
	RESULT_BANKER_WIN   = 0
	RESULT_PLAYER_WIN   = 1
	RESULT_PEACE        = 2
)

type Collections struct {
	curCnt   int     // 当前回合数
	maxCnt   int     // 最多可玩次数
	curPos   int     // 当前pos
	maxPos   int     // 最大位置
	cards    []int32 // 集合
	results  []int   // 0: 庄 1: 闲 2: 和
	banker11 []bool  // 庄对子
	player11 []bool  // 闲对子
}

func NewCollections(maxCnt int) *Collections {
	maxPos := CARD_NUM * CARD_COLLECTION_NUM
	if maxCnt <= 0 || maxCnt > maxPos/6 {
		maxCnt = DEFAULT_PLAY_CNT
	}
	collection := &Collections{
		curCnt:   1,
		maxCnt:   maxCnt,
		maxPos:   maxPos,
		results:  make([]int, 0),
		banker11: make([]bool, 0),
		player11: make([]bool, 0),
	}
	collection.initCards()
	return collection
}

func (this *Collections) initCards() {
	cards := make([]int32, 0)
	for i := 0; i < CARD_COLLECTION_NUM; i++ {
		for card := 0; card < CARD_NUM; card++ {
			cards = append(cards, int32(card))
		}
	}
	this.cards = FishedYates(cards)
}

func (this *Collections) Run() {
	for this.curCnt <= this.maxCnt {
		bankerCards, playerCards := this.dealCards()
		// 结算结果
		this.curCnt = this.curCnt + 1
		bankerPoint, bankerPair := CalculateCardsAndPair(bankerCards)
		playerPoint, playerPair := CalculateCardsAndPair(playerCards)
		this.banker11 = append(this.banker11, bankerPair)
		this.player11 = append(this.player11, playerPair)
		//
		var result int
		if bankerPoint > playerPoint {
			result = RESULT_BANKER_WIN
		} else if bankerPoint == playerPoint {
			result = RESULT_PEACE
		} else {
			result = RESULT_PLAYER_WIN
		}
		this.results = append(this.results, result)
	}
	this.finished()
}

func (this *Collections) getANewCard() int32 {
	old := this.curPos
	this.curPos = this.curPos + 1
	return this.cards[old]
}

/* 先发闲家
0	补一张牌	补一张牌
1	补一张牌	补一张牌
2	补一张牌	补一张牌
3	补一张牌	如果闲家补得第三张牌（非三张牌点数相加，下同）是8点，不须补牌，其他则需补牌
4	补一张牌	如果闲家补得第三张牌是0,1,8,9点，不须补牌，其他则需补牌
5	补一张牌	如果闲家补得第三张牌是0,1,2,3,8,9点，不须补牌，其他则需补牌
6	不须补牌	如果闲家需补牌（即前提是闲家为1至5点）而补得第三张牌是6或7点，补一张牌，其他则不需补牌
7	不须补牌	不须补牌
8	天牌，不须补牌	天牌，不须补牌
9	天牌，不须补牌	天牌，不须补牌
*/
func (this *Collections) dealCards() ([]int32, []int32) {
	bankerCards, playerCards := make([]int32, 0), make([]int32, 0)
	for i := 0; i < 2; i++ {
		playerCards = append(playerCards, this.getANewCard())
		bankerCards = append(bankerCards, this.getANewCard())
	}
	playerPoint := CalculateCards(playerCards)
	lastCard := CardToPoint(playerCards[len(playerCards)-1])
	// 闲家补牌规则
	switch playerPoint {
	case 0, 1, 2:
		playerCards = append(playerCards, this.getANewCard())
	case 3:
		if lastCard != 8 {
			playerCards = append(playerCards, this.getANewCard())
		}
	case 4:
		if lastCard != 0 && lastCard != 1 && lastCard != 8 && lastCard != 9 {
			playerCards = append(playerCards, this.getANewCard())
		}
	case 5:
		if lastCard != 0 && lastCard != 1 && lastCard != 2 && lastCard != 3 && lastCard != 9 && lastCard != 9 {
			playerCards = append(playerCards, this.getANewCard())
		}
	case 6, 7, 8, 9:
		// pass
	}
	playerPoint = CalculateCards(playerCards)
	bankerPoint := CalculateCards(bankerCards)
	if playerPoint == 8 || playerPoint == 9 {
		return bankerCards, playerCards
	}
	if bankerPoint < playerPoint {
		bankerCards = append(bankerCards, this.getANewCard())
	}
	return bankerCards, playerCards
}

func (this *Collections) finished() {
	log.Printf("result:%v bank11:%v play11:%v", len(this.results), len(this.banker11), len(this.player11))
	for i := 0; i < this.maxCnt; i++ {
		result := "庄赢"
		if this.results[i] == RESULT_PLAYER_WIN {
			result = "闲赢"
		} else if this.results[i] == RESULT_PEACE {
			result = "和"
		}
		info := fmt.Sprintf("第%02d盘: %s", i+1, result)
		if this.banker11[i] {
			info = info + " 庄对子"
		}
		if this.player11[i] {
			info = info + " 闲对子"
		}
		log.Printf("%s", info)
	}
}
