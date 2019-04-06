package web

import (
	"cli"
	"encoding/json"
	"log"
	"net/http"
	"p2phost"
)

// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// }

func handshake(w http.ResponseWriter, r *http.Request) {
	id := p2phost.Host.ID().Pretty()
	if id != "" {
		jData, err := json.Marshal(id)
		if err != nil {
			log.Panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(jData))
	}
}

// Init will start the web service
func Init() {
	cli.PostStatus("web", "Initializing web service...")

	port := ":8080"
	fs := http.FileServer(http.Dir("./web/public/build"))
	http.Handle("/", fs)
	http.HandleFunc("/handshake", handshake)
	cli.PostStatus("web", "You can now connect to the service")
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
	select {}
}
