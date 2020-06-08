package main

import (
	"Test/gcontrl"
	"Test/wmap"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
	"html/template"
	"net/http"
	"os"
)

func init() {

	filelog, e := os.Create("log")
	if e != nil {
		panic("error create log file")
	}
	log.SetOutput(filelog)
}
func main() {
	log.WithFields(log.Fields{
		"package": "main",
		"func":    "main",
	}).Info("Server start")
	World := wmap.NewCacheWorldMap()
	http.HandleFunc("/init", gcontrl.InitHandler(&World))
	http.HandleFunc("/map", gcontrl.Map_Handler(&World))
	http.Handle("/player", websocket.Handler(gcontrl.PlayerHandler(&World)))
	http.HandleFunc("/", indexHandler)

	//static
	http.Handle("/node_modules/phaser/dist/", http.StripPrefix("/node_modules/phaser/dist/", http.FileServer(http.Dir("./node_modules/phaser/dist/"))))
	http.Handle("/Client/", http.StripPrefix("/Client/", http.FileServer(http.Dir("./Client/"))))
	http.Handle("/Client/content/", http.StripPrefix("/Client/content/", http.FileServer(http.Dir("./Client/content/"))))
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.WithFields(log.Fields{
			"package": "main",
			"func":    "main",
			"error":   err,
		}).Fatal("Error start server")
	}

}

// Обработчик для index.html, здесь мы просто отдаем клиент пользователю
func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	err := t.Execute(w, "index")
	if err != nil {
		log.WithFields(log.Fields{
			"package": "main",
			"func":    "indexHandler",
			"error":   err,
		}).Error("Error get index.html")
	}
}
