package gcontrl

import (
	"Test/chunk"
	"Test/wmap"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
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
	chunk.Coordinate
}
func (p pingPlayer) GetCoordinate() chunk.Coordinate {
	return chunk.Coordinate{X:p.X, Y:p.Y}
}
func (p pingPlayer) GetID() string {
	return p.Name
}

func Map_Handler(W *wmap.WorldMap) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("MAP HANDLER")
		body, _ := ioutil.ReadAll(r.Body)

		rm := requestMap{}

		err := json.Unmarshal(body, &rm)
		if err != nil {
			log.WithFields(log.Fields{
				"package": "GameController",
				"func":    "InitHandler",
				"error":   err,
				"data":    body,
			}).Error("Error Marshal data")
		}
		fmt.Println(rm.X, rm.Y)

		c := wmap.GetChunkID(rm.X, rm.Y)
		d := wmap.GetCurrentPlayerMap(c)
		x := wmap.GetPlayerDrawChunkMap(d, W)
		playerMap := wmap.MapToJSON(x, rm.PlayerID)
		w.Header().Set("Content-Type", "application/json")
		w.Write(playerMap)

	}

}

func PlayerHandler(W *wmap.WorldMap) func(ws *websocket.Conn) {
	return func(ws *websocket.Conn) {
		defer func() {
			if err := recover(); err != nil {
				log.WithFields(log.Fields{
					"package": "GameController",
					"func":    "PlayerHandler",
					"error":   err,
				}).Error("Error ws")
			}

		}()

		player := pingPlayer{}
		websocket.JSON.Receive(ws, &player)

		//Game Loop
		log.WithFields(log.Fields{
			"package": "GameController",
			"func":    "PlayerHandler",
			"player":  player,
		}).Info("Connect player")

		for {
			err := websocket.JSON.Receive(ws, &player)
			if err != nil {
				log.WithFields(log.Fields{
					"package": "GameController",
					"func":    "PlayerHandler",
					"error":   err,
				}).Error("Connect cancel")
				return
			}

			W.UpdatePlayer(player)
			pls := W.GetPlayers()
			websocket.JSON.Send(ws, pls)
			log.WithFields(log.Fields{
				"Player": player,
			}).Info("Player log")

		}

	}
}
