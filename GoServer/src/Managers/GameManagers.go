package Managers

import (
	"GameServer/src/Room"
	"fmt"
)

type GameManager struct {
	room *Room.Room
}

var GMInstance GameManager

func init()  {
	fmt.Println("Init")
	GMInstance.room = Room.CreateRoom()
}

func(gm *GameManager) Update(time int64) {
	gm.room.Update(time)
}

func(gm *GameManager) GetRoom() *Room.Room{
	return gm.room
}
