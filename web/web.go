package web

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"

	"github.com/nwprince/RTS-Project/cli"
)

var sh *shell.Shell
var cid string = ""

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func addMedia() {
	var mediaExists bool = false
	var shouldConvert bool = true
	var filename string
	var outputDir string
	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".mkv" || filepath.Ext(path) == ".mp4" {
			filename = info.Name()
			name := strings.Split(info.Name(), ".")[0]
			outputDir = "./" + name
			if _, err := os.Stat(outputDir); os.IsNotExist(err) {
				os.Mkdir(outputDir, 0777)
			}
			if err != nil {
				return err
			}
			mediaExists = true
		}
		if filepath.Ext(path) == ".m3u8" {
			shouldConvert = false
			mediaExists = true
		}
		return nil
	})

	if err != nil {
		log.Panicf("walk error [%v]\n", err)
	}

	if !mediaExists {
		cli.PostStatus("error", "Ensure file is in this directory")
		os.Exit(1)
	} else if shouldConvert {
		dir, _ := os.Getwd()
		fullpath := dir + "/" + filename
		outputDir = dir + "/" + strings.Split(filename, ".")[0] + "/master.m3u8"

		cmd := exec.Command("ffmpeg", "-i", fullpath, "-profile:v", "baseline", "-level", "3.0", "-start_number", "0", "-hls_time", "5", "-hls_list_size", "0", "-f", "hls", outputDir)
		if err := cmd.Run(); err != nil {
			log.Print("error: ", err)
		}
	}
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
	cli.PostStatus("web", "Initializing web service...")

	//Configure fileserver for serving frontend
	port := ":8081"
	fs := http.FileServer(http.Dir("./web/public/build"))
	http.Handle("/", fs)
	http.HandleFunc("/handshake", handshake)

	addMedia()

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
