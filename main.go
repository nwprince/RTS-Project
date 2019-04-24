package main

import (
	"github.com/nwprince/RTS-Project/cli"
	"github.com/nwprince/RTS-Project/web"
)

func main() {
	go cli.PostStatus("app", "Initializing...")
	web.Init()
}
