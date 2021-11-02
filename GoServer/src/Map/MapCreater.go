package Map

import (
	"GameServer/utils"
	"fmt"
	"math/rand"
	"time"
)

type MapCreater struct {
	//随机种子，当前所有存在的行的都保存
	mapRandSeed[utils.MapRow] int64
	mapRandSeedHead int
	//最底层计数器
	oldestLinePos int32
	//间隔
	interval int64
	//计时
	timeRecord int64
}

func NewMapCreater() *MapCreater{
	nc := &MapCreater{
		timeRecord: 0,
		mapRandSeedHead: 0,
		oldestLinePos: 0,
	}
	nc.interval = 5000
	for i := 0; i < utils.MapRow; i++ {
		nc.mapRandSeed[i] = int64(i)
	}
	rand.Seed(time.Now().UnixNano())
	return nc
}

func (mc *MapCreater) GetMapRandSeedList() [utils.MapRow]int64  {
	var res[utils.MapRow] int64
	index := mc.mapRandSeedHead
	for i := 0; i < utils.MapRow; i++{
		res[i] = mc.mapRandSeed[index]
		index++
		if index >= utils.MapRow{
			index = 0
		}
	}
	return res
}

func (mc *MapCreater) GetInfo() ([utils.MapRow]int64, int32) {
	return mc.mapRandSeed, mc.oldestLinePos
}

func (mc *MapCreater) Update(millisecond int64){
	timestamp := utils.GetUnixMillis()     //时间戳
	if timestamp - mc.timeRecord > mc.interval {
		mc.MakeNewLine(timestamp)
		mc.timeRecord = timestamp
	}
}
//产生新的一行
func (mc *MapCreater) MakeNewLine(time int64)  {
	fmt.Println("MakeNewLine")
	mc.mapRandSeed[mc.mapRandSeedHead] = time
	mc.mapRandSeedHead++
	if mc.mapRandSeedHead >= utils.MapRow{
		mc.mapRandSeedHead = 0
	}
	mc.oldestLinePos += (utils.MapLineInterval + utils.MapLineHeight)
}
