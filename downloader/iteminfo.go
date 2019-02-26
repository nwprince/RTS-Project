package downloader

type EmbyFileInfo struct {
	MediaSources []struct {
		Protocol              string `json:"Protocol"`
		ID                    string `json:"Id"`
		Path                  string `json:"Path"`
		Type                  string `json:"Type"`
		Container             string `json:"Container"`
		Size                  int64  `json:"Size"`
		Name                  string `json:"Name"`
		IsRemote              bool   `json:"IsRemote"`
		RunTimeTicks          int64  `json:"RunTimeTicks"`
		ReadAtNativeFramerate bool   `json:"ReadAtNativeFramerate"`
		DiscardCorruptPts     bool   `json:"DiscardCorruptPts"`
		FillWallClockDts      bool   `json:"FillWallClockDts"`
		IgnoreDts             bool   `json:"IgnoreDts"`
		IgnoreIndex           bool   `json:"IgnoreIndex"`
		SupportsTranscoding   bool   `json:"SupportsTranscoding"`
		SupportsDirectStream  bool   `json:"SupportsDirectStream"`
		SupportsDirectPlay    bool   `json:"SupportsDirectPlay"`
		IsInfiniteStream      bool   `json:"IsInfiniteStream"`
		RequiresOpening       bool   `json:"RequiresOpening"`
		RequiresClosing       bool   `json:"RequiresClosing"`
		RequiresLooping       bool   `json:"RequiresLooping"`
		SupportsProbing       bool   `json:"SupportsProbing"`
		MediaStreams          []struct {
			Codec                  string `json:"Codec"`
			Language               string `json:"Language,omitempty"`
			ColorTransfer          string `json:"ColorTransfer,omitempty"`
			ColorPrimaries         string `json:"ColorPrimaries,omitempty"`
			ColorSpace             string `json:"ColorSpace,omitempty"`
			TimeBase               string `json:"TimeBase,omitempty"`
			CodecTimeBase          string `json:"CodecTimeBase,omitempty"`
			Title                  string `json:"Title,omitempty"`
			VideoRange             string `json:"VideoRange,omitempty"`
			DisplayTitle           string `json:"DisplayTitle,omitempty"`
			NalLengthSize          string `json:"NalLengthSize,omitempty"`
			IsInterlaced           bool   `json:"IsInterlaced"`
			IsAVC                  bool   `json:"IsAVC,omitempty"`
			BitRate                int    `json:"BitRate,omitempty"`
			BitDepth               int    `json:"BitDepth,omitempty"`
			RefFrames              int    `json:"RefFrames,omitempty"`
			IsDefault              bool   `json:"IsDefault"`
			IsForced               bool   `json:"IsForced"`
			Height                 int    `json:"Height,omitempty"`
			Width                  int    `json:"Width,omitempty"`
			AverageFrameRate       int    `json:"AverageFrameRate,omitempty"`
			RealFrameRate          int    `json:"RealFrameRate,omitempty"`
			Profile                string `json:"Profile,omitempty"`
			Type                   string `json:"Type"`
			AspectRatio            string `json:"AspectRatio,omitempty"`
			Index                  int    `json:"Index"`
			IsExternal             bool   `json:"IsExternal"`
			IsTextSubtitleStream   bool   `json:"IsTextSubtitleStream"`
			SupportsExternalStream bool   `json:"SupportsExternalStream"`
			PixelFormat            string `json:"PixelFormat,omitempty"`
			Level                  int    `json:"Level,omitempty"`
			IsAnamorphic           bool   `json:"IsAnamorphic,omitempty"`
			ChannelLayout          string `json:"ChannelLayout,omitempty"`
			Channels               int    `json:"Channels,omitempty"`
			SampleRate             int    `json:"SampleRate,omitempty"`
			Path                   string `json:"Path,omitempty"`
		} `json:"MediaStreams"`
		Formats             []interface{} `json:"Formats"`
		Bitrate             int           `json:"Bitrate"`
		RequiredHTTPHeaders struct {
		} `json:"RequiredHttpHeaders"`
	} `json:"MediaSources"`
	PlaySessionID string `json:"PlaySessionId"`
}
