package downloader

import (
	"os"
	"testing"
)

var checkArgsTest = []struct {
	n        []string // input
	expected bool     // expected result
}{
	{[]string{"cmd"}, false},
	{[]string{"cmd", "prepare"}, false},
	{[]string{"cmd", "prepare", "blah"}, true},
}

func TestCheckArgs(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	for _, tt := range checkArgsTest {
		os.Args = tt.n
		actual := checkArgs()
		if actual != tt.expected {
			t.Errorf("Test failed, expected: '%t', got: '%t'", tt.expected, actual)
		}
	}
}

var shouldDownloadTest = []struct {
	n        []string // input
	expected bool     // expected result
}{
	{[]string{"cmd", "prepar"}, false},
	{[]string{"cmd", "prepare"}, true},
	{[]string{"cmd", "prepare", "blah"}, true},
}

func TestShouldDownload(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	for _, tt := range shouldDownloadTest {
		os.Args = tt.n
		actual := shouldDownload()
		if actual != tt.expected {
			t.Errorf("Test failed, expected: '%t', got: '%t'", tt.expected, actual)
		}
	}
}

type buildURIArgs struct {
	endpoint    string
	requireUser bool
	requireKey  bool
	query       string
}

var buildURITest = []struct {
	n        *buildURIArgs // input
	expected string        // expected result
}{
	{&buildURIArgs{"Users/*/Items", true, true, "Recursive=true&IncludeItemTypes=movie"}, "http://r06nwpkbd.device.mst.edu:8096/emby/Users/a6721c39bc1d4d8fb356cac7d45f8db5/Items?api_key=b2d4af7255f94a7ca73fd85a84ebbd3a&Recursive=true&IncludeItemTypes=movie"},
	{&buildURIArgs{"Users/*/Items", false, true, "Recursive=true&IncludeItemTypes=movie"}, "http://r06nwpkbd.device.mst.edu:8096/emby/Users/*/Items?api_key=b2d4af7255f94a7ca73fd85a84ebbd3a&Recursive=true&IncludeItemTypes=movie"},
	{&buildURIArgs{"Users/*/Items", true, false, "Recursive=true&IncludeItemTypes=movie"}, "http://r06nwpkbd.device.mst.edu:8096/emby/Users/a6721c39bc1d4d8fb356cac7d45f8db5/Items?Recursive=true&IncludeItemTypes=movie"},
	{&buildURIArgs{"Users/*/Items", true, false, ""}, "http://r06nwpkbd.device.mst.edu:8096/emby/Users/a6721c39bc1d4d8fb356cac7d45f8db5/Items"},
	{&buildURIArgs{"Items/1/File", false, true, ""}, "http://r06nwpkbd.device.mst.edu:8096/emby/Items/1/File?api_key=b2d4af7255f94a7ca73fd85a84ebbd3a&"},
	{&buildURIArgs{"Items/*/File", true, true, ""}, "http://r06nwpkbd.device.mst.edu:8096/emby/Items/a6721c39bc1d4d8fb356cac7d45f8db5/File?api_key=b2d4af7255f94a7ca73fd85a84ebbd3a&"},
	{&buildURIArgs{"Items/1/File", false, false, ""}, "http://r06nwpkbd.device.mst.edu:8096/emby/Items/1/File"},
	{&buildURIArgs{"Items/1/PlaybackInfo", false, true, ""}, "http://r06nwpkbd.device.mst.edu:8096/emby/Items/1/PlaybackInfo?api_key=b2d4af7255f94a7ca73fd85a84ebbd3a&"},
	{&buildURIArgs{"Items/1/PlaybackInfo", false, false, ""}, "http://r06nwpkbd.device.mst.edu:8096/emby/Items/1/PlaybackInfo"},
}

func TestGetContainer(t *testing.T) {
	for n, tt := range buildURITest {
		actual := buildURI(tt.n.endpoint, tt.n.requireUser, tt.n.requireKey, tt.n.query)
		if actual != tt.expected {
			t.Errorf("Test %d failed, expected: '%s', got: '%s'", n, tt.expected, actual)
		}
	}
}
