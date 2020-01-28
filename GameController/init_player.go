package GameController

import (
	"Test/WorldMap"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type requestPlayer struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
type responsePlayer struct {
	Error string `json:"error"`
	Name  string `json:"name"` // заменить на уникальный ид в будущем
	X     int    `json:"x"`
	Y     int    `json:"y"`
}

// Точка входа в игры, юзер отправляет нам свои данные, мы отдаем данные персонажа, уникальный ид или name через которое будет совершенно socket подключение
func Init_Handler(W *WorldMap.WorldMap) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Init")
		body, _ := ioutil.ReadAll(r.Body)
		rp := requestPlayer{}
		err := json.Unmarshal(body, &rp)
		if err != nil {
			fmt.Println("Error Marshaler")
		}
		w.Header().Set("Content-Type", "application/json")
		p, exile := W.GetPlayer(rp.Name)
		if exile {
			ok := p.ComparePassword(rp.Password)
			if ok {
				resPl := responsePlayer{Error: "null", X: p.X, Y: p.Y, Name: p.Name}
				res, err := json.Marshal(resPl)
				if err != nil {
					fmt.Println(err.Error())
					w.Write([]byte("{Error: error server}"))
					return
				}
				w.Write(res)
				return
			}
		} else {
			p := WorldMap.NewPlayer(rp.Name, rp.Password)
			W.AddPlayer(p)
			resPl := responsePlayer{Error: "null", X: p.X, Y: p.Y, Name: p.Name}
			res, err := json.Marshal(resPl)
			if err != nil {
				fmt.Println(err.Error())
				w.Write([]byte("{Error: error server}"))
				return
			}
			w.Write(res)
			return

		}


	}
}
