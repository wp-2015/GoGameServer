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
	seedList := Managers.GMInstance.GetRoom().GetMapCreater().GetMapRandSeedList()
	for _, v := range seedList {
		s2cMapInit.MapRandSeeds = append(s2cMapInit.MapRandSeeds, int32(v))
	}

	data, _ := proto.Marshal(s2cMapInit)
	code := uint16(protocol.MSGTYPE_Mt_S2C_MapInit)
	request.GetConnection().SendBuffMsg(code, data)
	request.GetConnection().SendBuffMsg(code, data)
	//request.GetConnection().SendBuffMsg(code, data)
	fmt.Println("收到Join",1)
}
