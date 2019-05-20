package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"github.com/nwprince/RTS-Project/cli"
	"github.com/nwprince/RTS-Project/ipfs"
)

var cid string

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func mediaInit() {
	var mediaExists = false
	var shouldConvert = true
	var filename string
	var name string
	var outputDir string
	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".mkv" || filepath.Ext(path) == ".mp4" {
			filename = info.Name()
			name = strings.Split(info.Name(), ".")[0]
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

	out, errPin := ipfs.CheckForPin(name + "/")
	if errPin != "" {
		fmt.Println(errPin)
		return
	}

	if out == "" {
		var errAdd string
		cid, errAdd = ipfs.AddDir(name + "/")
		if errAdd != "" {
			fmt.Println(errAdd)
			return
		}

		_, errPin := ipfs.Pin(cid)
		if errPin != "" {
			fmt.Println(errPin)
			return
		}
	}

	cid = out

	return
}

func handshake(c echo.Context) error {
	jData := &Base{
		Initialized: false,
	}
	if cid != "" {
		jData.Initialized = true
		jData.Hash = cid
	}
	// msg, _ := json.Marshal(jData)
	return c.JSON(http.StatusOK, jData)
}

// Init will start the web service
func Init() {
	cli.PostStatus("web", "Initializing web service...")

	r := echo.New()
	r.Use(middleware.Recover())
	r.Use(middleware.CORS())
	r.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	r.Static("/", "/home/nwprince/Projects/Go/media/src/web/public/build/")

	// r.StaticFile("/", "/home/nwprince/Projects/Go/media/src/web/public/build/")
	r.GET("/handshake", handshake)

	//Configure fileserver for serving frontend
	port := ":8081"

	go mediaInit()

	go cli.PostStatus("web", "You can now connect to the service at localhost"+port)

	r.Logger.Fatal(r.Start(port))
}
