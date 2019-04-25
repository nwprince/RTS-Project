package main

import (
	"github.com/nwprince/RTS-Project/cli"
	"github.com/nwprince/RTS-Project/web"
	"os/exec"
)

func init() {
	// Check for dependencies
	_, err := exec.LookPath("ffmpeg")
	if err != nil {
		cli.PostStatus("error", "ffmpeg could not be found. Please make sure it is installed and in your PATH")
	}
	_, err = exec.LookPath("ipfs")
	if err != nil {
		cli.PostStatus("error", "IPFS could not be found. Please make sure it is installed and in your PATH")
	}
}

func main() {
	cli.PostStatus("app", "Initializing...")
	web.Init()
}
