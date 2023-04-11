package model

type PingResponse struct {
	ID   string `json:"id""`
	Pong string `json:pong,omitempty`
}
