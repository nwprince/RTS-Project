package web

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

var ipTable = make(map[string]bool)

func startConnection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Here")
	enableCors(&w)
	newIP := strings.Split(r.RemoteAddr, ":")[0]
	if ipTable[newIP] {
		fmt.Println("Web: this ip is already in the network")
	} else {
		ipTable[newIP] = true
		go manageIP(newIP)
	}
}

func manageIP(ip string) {
	worker := NewWorker(3 * time.Second)
	go worker.Run()
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ping")
	enableCors(&w)
}
func Init() {
	fmt.Println("Web: Starting web service")

	fs := http.FileServer(http.Dir("./web/public"))
	http.Handle("/", fs)
	http.HandleFunc("/init", startConnection)
	http.HandleFunc("/ping", ping)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
