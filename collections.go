package main

const (
	CARD_NUM            = 52 // 一副牌多少牌 没有大小王
	CARD_COLLECTION_NUM = 8  // 多少副牌
	DEFAULT_PLAY_CNT    = 55 // 默认回合数
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
	for this.curCnt < this.maxCnt {
		bankerCards, playerCards := this.dealCards()
		// 结算结果
		this.curCnt = this.curCnt + 1
		bankerPoint, bankerPair := CalculateCards(bankerCards)
		playerPoint, playerPair := CalculateCards(playerCards)
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
}

func (this *Collections) dealCards() ([]int32, []int32) {
	//
	return nil, nil
}

func (this *Collections) finished() {
	for i := 0; i < this.maxCnt; i++ {
	}
}
