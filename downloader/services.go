package downloader

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/nwprince/RTS-Project/cli"
)

var baseURL = "http://r06nwpkbd.device.mst.edu:8096/emby/"
var userID = "a6721c39bc1d4d8fb356cac7d45f8db5"
var apiKey = "b2d4af7255f94a7ca73fd85a84ebbd3a"

func shouldDownload() bool {
	if os.Args[1] == "prepare" {
		return true
	}
	return false
}
func buildURI(endpoint string, requireUser bool, requireKey bool, query string) string {
	url := baseURL + endpoint
	if requireUser == true {
		url = strings.Replace(url, "*", userID, 1)
	}
	if requireKey == true || query != "" {
		url = url + "?"
	}

	if requireKey == true {
		url = url + "api_key=" + apiKey + "&"
	}
	if query != "" {
		url = url + query
	}
	return url
}

func initDownload(media *Media) {
	var allItems = new(EmbyItems)
	url := buildURI("Users/*/Items", true, true, "Recursive=true&IncludeItemTypes="+media.MediaType)
	getJSON(url, allItems)
	if allItems == nil {
		panic("Error: No items in list")
	}
	for _, item := range allItems.Items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(media.Title)) {
			media.ID = item.ID
			getContainer(media)
			downloadItem(media)
		}
	}
}

func getContainer(media *Media) {
	url := buildURI("Items/"+media.ID+"/PlaybackInfo", false, true, "")
	fileInfo := new(EmbyFileInfo)
	getJSON(url, fileInfo)
	media.Container = fileInfo.MediaSources[0].Container
}

func downloadItem(media *Media) {
	url := buildURI("Items/"+media.ID+"/File", false, true, "")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	out, err := os.Create(media.Title + "." + media.Container)
	if err != nil {
		panic(err)
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
}

func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func checkArgs() bool {
	if len(os.Args) < 3 {
		return false
	}
	return true
}

func Init() Media {
	var media = new(Media)

	if checkArgs() == false {
		cli.Error("downloader", "Please use proper params")
		return Media{}
	}

	if shouldDownload() == true {
		cli.PostStatus("downloader", "Will download new media")
		media.Prepared = false
		media.MediaType = os.Args[2]
		media.Title = os.Args[3]
		initDownload(media)
		fmt.Println("downloader", "Finished downloading "+media.Title)
	} else {
		media.Prepared = true
		media.Title = os.Args[1]
	}

	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".mkv" || filepath.Ext(path) == ".mp4" {
			media.Container = filepath.Ext(path)
			media.Title = strings.TrimSuffix(info.Name(), filepath.Ext(path))
			media.FileName = info.Name()
		}
		return nil
	})
	if err != nil {
		log.Panicf("walk error [%v]\n", err)
	}

	return *media
}
