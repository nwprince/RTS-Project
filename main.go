package main

import (
	"cli"
	"downloader"
	"p2phost"
	"web"
)

func main() {
	go cli.PostStatus("app", "Initializing...")
	/* media :=*/ downloader.Init()
	go web.Init()
	p2phost.InitP2P()

}
