package web

import (
	"cli"
	"fmt"
	"log"
	"net"
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

func getHostIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}

// Init will start the web service
func Init() {
	cli.PostStatus("web", "Initializing web service...")

	port := ":8080"

	fs := http.FileServer(http.Dir("./web/public/build"))
	http.Handle("/", fs)
	http.HandleFunc("/init", startConnection)
	http.HandleFunc("/ping", ping)
	ip := getHostIP()
	cli.PostStatus("web", "You can now connect to the service at http://"+ip+port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
