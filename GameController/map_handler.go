package GameController

import (
	"Test/Chunk"
	"Test/WorldMap"
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"net/http"
)

type requestMap struct {
	X        int
	Y        int
	PlayerID int
}
type pingPlayer struct {
	Name string `json:"name"`
	Chunk.Coordinate
}

func Map_Handler(W *WorldMap.WorldMap) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("MAP HANDLER")
		body, _ := ioutil.ReadAll(r.Body)

		rm := requestMap{}

		err := json.Unmarshal(body, &rm)
		if err != nil {
			fmt.Println("Error Marshaler")
		}
		fmt.Println(rm.X, rm.Y)

		c := WorldMap.GetChunkID(rm.X, rm.Y)
		d := WorldMap.GetCurrentPlayerMap(c)
		x := WorldMap.GetPlayerDrawChunkMap(d, W)
		playerMap := WorldMap.MapToJSON(x, rm.PlayerID)
		w.Header().Set("Content-Type", "application/json")
		w.Write(playerMap)

	}

}

func PlayerHandler(W *WorldMap.WorldMap) func(ws *websocket.Conn) {
	return func(ws *websocket.Conn) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Close Conn")
			}

		}()

		player := pingPlayer{}
		websocket.JSON.Receive(ws, &player)
		fmt.Println(player)

		//Game Loop
		fmt.Println("Connect Player", player.Name)
		for {
			err := websocket.JSON.Receive(ws, &player)
			if err != nil {
				err.Error()
				return
			}
			W.Player[player.Name].SetWalkPath(player.X, player.Y, W)
			pls := W.GetPlayers()
			websocket.JSON.Send(ws, pls)
			//fmt.Println("Ping", player.Name, player.X, player.Y)

		}

	}
}
