package downloader

type EmbyItems struct {
	Items            []EmbyItemInfo `json:"Items"`
	TotalRecordCount int16          `json:"TotalRecordcount"`
}

type EmbyItemInfo struct {
	Name            string  `json:"Name"`
	ServerID        string  `json:"ServerId"`
	ID              string  `json:"Id"`
	HasSubtitles    bool    `json:"HasSubtitles"`
	PremiereDate    string  `json:"PremiereDate"`
	CriticRating    int8    `json:"CriticRating"`
	OfficialRating  string  `json:"OfficialRating"`
	CommunityRating float32 `json:"CommunityRating"`
	RunTimeTicks    int     `json:"RunTimeTicks"`
	ProductionYear  int16   `json:"ProductionYear"`
	IsFolder        bool    `json:"IsFolder"`
	Type            string  `json:"Type"`
	IUserData
	IImageTags
	BackdropImageTags []string `json:"BackdropImageTags"`
	MediaType         string   `json:"MediaType"`
}

type IUserData struct {
	PlaybackPositionTicks int8   `json:"PlaybackPositionTicks"`
	PlayCount             int8   `json:"PlayCount"`
	IsFavorite            bool   `json:"IsFavorite"`
	Played                bool   `json:"Played"`
	Key                   string `json:"Key"`
}

type IImageTags struct {
	Primary string `json:"Primary"`
	Logo    string `json:"Logo"`
	Thumb   string `json:"Thumb"`
}
