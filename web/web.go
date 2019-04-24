package web

import (
	"encoding/json"
	"net/http"

	shell "github.com/ipfs/go-ipfs-api"

	"github.com/nwprince/RTS-Project/cli"
)

var sh *shell.Shell
var cid string

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func addMedia(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
}

func handshake(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	jData := &Base{
		Initialized: false,
	}
	if cid != "" {
		jData.Initialized = true
	}
	msg, _ := json.Marshal(jData)

	w.Header().Set("Content-Type", "application/json")
	w.Write(msg)
}

// Init will start the web service
func Init() {
	go cli.PostStatus("web", "Initializing web service...")

	//Configure fileserver for serving frontend
	port := ":8080"
	fs := http.FileServer(http.Dir("./web/public/build"))
	http.Handle("/", fs)
	http.HandleFunc("/handshake", handshake)

	//Configure a ipfs connection
	go func() {
		sh = shell.NewShell("localhost:5001")
		cli.PostStatus("web", "IPFS Shell is initialized")
	}()

	go cli.PostStatus("web", "You can now connect to the service")
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
