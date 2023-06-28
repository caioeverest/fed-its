package model

type Envelope struct {
	Provider string `json:"provider"`
	Version  string `json:"version"`
	Result   any    `json:"result"`
}
