package MessageHandles

import (
	protocol "GameServer/src/PB"
	"GameServer/ziface"
)

func AddHandles(server ziface.IServer){
	addFuncHandle(server, protocol.MSGTYPE_Mt_C2S_Join, C2SJoin)
}

func addFuncHandle(server ziface.IServer, id protocol.MSGTYPE, funcHandle func(request ziface.IRequest)){
	code := uint16(id)
	server.AddFuncRouter(code, funcHandle)
}