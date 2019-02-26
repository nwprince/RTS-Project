package main

import (
	"downloader"
	"fmt"
)

func main() {
	media := downloader.Init()
	fmt.Println(media)
}
