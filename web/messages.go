package web

// Base is the base json message
type Base struct {
	Initialized bool        `json:"init"`
	Type        string      `json:"type,omitempty"`
	MediaInfo   interface{} `json:"media_info,omitempty"`
	Hash        string      `json:"hash"`
}
