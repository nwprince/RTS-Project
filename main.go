package main

import (
	"cli"
	"downloader"
	"web"
)

func main() {
	cli.PostStatus("app", "Initializing...")
	/* media :=*/ downloader.Init()
	web.Init()
}
