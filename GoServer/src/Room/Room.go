package Room

import "GameServer/src/Map"

type Room struct {
	mapCreater *Map.MapCreater
}

func CreateRoom() *Room{
	res := &Room{}
	//地图创建
	res.mapCreater = Map.NewMapCreater()
	return res
}

func(r *Room) Update(time int64)  {
	r.mapCreater.Update(time)
}

func(r *Room) GetMapCreater() *Map.MapCreater{
	return r.mapCreater
}
