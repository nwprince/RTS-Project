package cli

import (
	"fmt"

	"github.com/gookit/color"
)

// PostStatus will show which package is generating a msg and that message
func PostStatus(src string, msg string) {
	src = getColor(src)

	fmt.Println(src, msg)
}

// Error prints the error code and shows the package responsible
func Error(src string, msg string) {
	e := getColor("error")
	src = getColor(src)
	fmt.Printf("%s -> %s : %s\n", e, src, msg)
}

func getColor(src string) string {

	switch src {
	case "app":
		return color.New(color.FgWhite, color.BgGreen, color.OpBold).Render("AppName")
	case "downloader":
		return color.New(color.FgWhite, color.BgYellow, color.OpBold).Render("Downloader")
	case "error":
		return color.New(color.FgWhite, color.BgRed, color.OpBold).Render("ERROR")

	case "web":
		return color.New(color.FgWhite, color.BgMagenta, color.OpBold).Render("Web")

	case "worker":
		return color.New(color.FgWhite, color.BgCyan, color.OpBold).Render("Worker")
	}

	return ""
}
