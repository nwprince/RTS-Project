package main

import (
	"github.com/nwprince/RTS-Project/cli"
	"github.com/nwprince/RTS-Project/downloader"
	"github.com/nwprince/RTS-Project/p2phost"
	"github.com/nwprince/RTS-Project/web"
)

func main() {
	go cli.PostStatus("app", "Initializing...")
	/* media :=*/ downloader.Init()
	go web.Init()
	p2phost.InitP2P()

}
