package Map

import (
	"GameServer/utils"
	"fmt"
)

type MapCreater struct {
	row int //当前行
	//随机种子，当前所有存在的行的都保存
	mapRandSeed[utils.MapRow] int
	mapRandSeedHead int

	//间隔
	interval int64
	//计时
	timeRecord int64
}

func NewMapCreater() *MapCreater{
	nc := &MapCreater{
		row: 0,
		timeRecord: 0,
		mapRandSeedHead: 0,
	}
	nc.interval = 5000
	for i := 0; i < utils.MapRow; i++ {
		nc.mapRandSeed[i] = i
	}
	return nc
}

func (mc *MapCreater) GetMapRandSeedList() [utils.MapRow]int  {
	return mc.mapRandSeed
}

func (mc *MapCreater) Update(millisecond int64){
	timestamp := utils.GetUnixMillis()     //时间戳
	if timestamp - mc.timeRecord > mc.interval {
		mc.MakeNewLine(timestamp)
		mc.timeRecord = timestamp
	}
}

func (mc *MapCreater) MakeNewLine(time int64)  {
	fmt.Println("MakeNewLine")
	mc.mapRandSeed[mc.mapRandSeedHead] = int(time)
	mc.mapRandSeedHead++
	if mc.mapRandSeedHead >= utils.MapRow{
		mc.mapRandSeedHead = 0
	}
	mc.SortSeed()
}

func (mc *MapCreater) SortSeed(){
	var res[utils.MapRow] int
	index := mc.mapRandSeedHead
	for i := utils.MapRow - 1; i > -1; i--{
		res[i] = mc.mapRandSeed[index]
		index++
		if index >= utils.MapRow{
			index = 0
		}
	}
}
