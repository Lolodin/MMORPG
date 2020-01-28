package main

import (
	"Test/GameController"
	"Test/WorldMap"
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("start")
	World := WorldMap.NewCacheWorldMap()
	http.HandleFunc("/init", GameController.Init_Handler(&World))
	http.HandleFunc("/map", GameController.Map_Handler(&World))
	http.Handle("/player", websocket.Handler(GameController.Player_Handler(&World)))
	http.HandleFunc("/", indexHandler)
	http.Handle("/node_modules/phaser/dist/", http.StripPrefix("/node_modules/phaser/dist/", http.FileServer(http.Dir("./node_modules/phaser/dist/"))))
	http.Handle("/Client/", http.StripPrefix("/Client/", http.FileServer(http.Dir("./Client/"))))
	http.Handle("/Client/content/", http.StripPrefix("/Client/content/", http.FileServer(http.Dir("./Client/content/"))))
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("indexAction")
	t, _ := template.ParseFiles("test.html")
	err := t.Execute(w, "test")
	if err != nil {
		fmt.Println(err.Error())
	}
}
