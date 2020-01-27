package GameController

import (
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
type requestInit struct {

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

		c := WorldMap.GetChankID(rm.X, rm.Y)
		d := WorldMap.GetCurrentPlayerMap(c)
		x := WorldMap.GetPlayerDrawChunkMap(d, W)
		playerMap := WorldMap.MapToJSON(x, rm.PlayerID)
		w.Header().Set("Content-Type", "application/json")
		w.Write(playerMap)

	}

}

func Player_Handler(W *WorldMap.WorldMap) func(ws *websocket.Conn) {
	return func(ws *websocket.Conn) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Close Conn")
			}

		}()

		player := WorldMap.Player{}
		websocket.JSON.Receive(ws, &player)
		fmt.Println()
		W.AddPlayer(player)

		//Game Loop
		fmt.Println("Connect Player", player.Name)
		for {
			websocket.JSON.Receive(ws, &player)
			W.UpdatePlayer(player)
			pls := W.GetPlayers()
			websocket.JSON.Send(ws, pls)

		}

	}
}

func Init_Handler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("MAP HANDLER")
		body, _ := ioutil.ReadAll(r.Body)

		rm := requestMap{}

		err := json.Unmarshal(body, &rm)
		if err != nil {
			fmt.Println("Error Marshaler")
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write()
	}
}