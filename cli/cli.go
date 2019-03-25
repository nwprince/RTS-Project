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
		c := color.New(color.FgWhite, color.BgGreen, color.OpBold).Render
		return c("AppName")
	case "downloader":
		c := color.New(color.FgWhite, color.BgYellow, color.OpBold).Render
		return c("Downloader")
	case "error":
		c := color.New(color.FgWhite, color.BgRed, color.OpBold).Render
		return c("ERROR")

	case "web":
		c := color.New(color.FgWhite, color.BgMagenta, color.OpBold).Render
		return c("Web")
	}

	return ""
}
