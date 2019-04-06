package cli

import (
	"testing"

	"github.com/gookit/color"
)

func TestColor(t *testing.T) {
	if getColor("app") != color.New(color.FgWhite, color.BgGreen, color.OpBold).Render("AppName") {
		t.Error("Expected AppName with a green bg")
	}

	if getColor("downloader") != color.New(color.FgWhite, color.BgYellow, color.OpBold).Render("Downloader") {
		t.Error("Expected Downloader with a yellow bg")
	}

	if getColor("error") != color.New(color.FgWhite, color.BgRed, color.OpBold).Render("ERROR") {
		t.Error("Expected ERROR with a red bg")
	}

	if getColor("web") != color.New(color.FgWhite, color.BgMagenta, color.OpBold).Render("Web") {
		t.Error("Expected Web with a magenta bg")
	}

	if getColor("worker") != color.New(color.FgWhite, color.BgCyan, color.OpBold).Render("Worker") {
		t.Error("Expected Worker with a cyan bg")
	}
}
