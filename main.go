package main

import (
	"downloader"
	"fmt"
	"web"
)

func main() {
	media := downloader.Init()
	fmt.Println(media)
	web.Init()
}
