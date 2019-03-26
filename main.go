package main

import (
	"cli"
	"downloader"
	"web"
)

func main() {
	go cli.PostStatus("app", "Initializing...")
	/* media :=*/ downloader.Init()
	web.Init()
}
