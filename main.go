package main

import "downloader"

// var shouldDownload = false
// var mediaType string
// var title string

// func getAllItems() {
// 	url := baseURL + "Users/" + userID + "/Items?api_key=" + apiKey + "&Recursive=true&IncludeItemTypes="
// 	if mediaType == "m" {
// 		url = url + "Movie"
// 	} else {
// 		url = url + "Episode"
// 	}

// 	getJson(url, mEmbyAllItems)
// }

// func getItem() {
// 	for _, item := range mEmbyAllItems.Items {
// 		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(title)) {
// 			downloadItem(item.ID)
// 		}
// 	}
// }

// func downloadItem(itemID string) error {

// 	getMediaFileInfo(itemID)
// 	url := baseURL + "Items/" + itemID + "/File?api_key=" + apiKey

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	out, err := os.Create(title + "." + mFileInfo.MediaSources[0].Container)
// 	if err != nil {
// 		return err
// 	}
// 	defer out.Close()

// 	_, err = io.Copy(out, resp.Body)
// 	fmt.Println("Done downloading")
// 	return err
// }

// func getMediaFileInfo(itemID string) {
// 	url := baseURL + "Items/" + itemID + "/PlaybackInfo?api_key=" + apiKey
// 	getJson(url, mFileInfo)
// }

// func getJson(url string, target interface{}) error {
// 	r, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer r.Body.Close()
// 	return json.NewDecoder(r.Body).Decode(target)
// }

// func serveMovie(w http.ResponseWriter, r *http.Request) {

// }

func main() {
	// handleArgs()
	downloader.Init()
	// http.HandleFunc("/", serveMovie)
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	panic(err)
	// }
}
