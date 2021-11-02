package MessageHandles

import (
	"GameServer/src/Managers"
	"fmt"

	//"GameServer/src/Map"
	protocol "GameServer/src/PB"
	"GameServer/utils"
	"GameServer/ziface"
	"github.com/golang/protobuf/proto"
)

func init() {
}

func C2SJoin(request ziface.IRequest)  {
	fmt.Println("收到Join")
	s2cMapInit := new(protocol.S2CMapInit)

	s2cMapInit.MapRow = int32(utils.MapRow)
	s2cMapInit.MapLineHeight = utils.MapLineHeight
	s2cMapInit.MapLineInterval = utils.MapLineInterval
	seedList, oldestLinePos := Managers.GMInstance.GetRoom().GetMapCreater().GetInfo()
	for i := 0; i < len(seedList); i++{
		s2cMapInit.MapRandSeeds = append(s2cMapInit.MapRandSeeds, seedList[i])
	}
	s2cMapInit.MapOldestLinePos = oldestLinePos

	data, _ := proto.Marshal(s2cMapInit)
	code := uint16(protocol.MSGTYPE_Mt_S2C_MapInit)

	request.GetConnection().SendBuffMsg(code, data)
	fmt.Println("收到Join",1)
}
