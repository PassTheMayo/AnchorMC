package types

import "github.com/anchormc/anchor/src/chat"

type StatusResponse struct {
	Version     StatusResponseVersion `json:"version"`
	Players     StatusResponsePlayers `json:"players"`
	Description chat.Chat             `json:"description"`
	Favicon     *string               `json:"favicon"`
}

type StatusResponseVersion struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type StatusResponsePlayers struct {
	Online int                          `json:"online"`
	Max    int                          `json:"max"`
	Sample []StatusResponseSamplePlayer `json:"sample,omitempty"`
}

type StatusResponseSamplePlayer struct {
	Username string `json:"name"`
	UUID     string `json:"id"`
}
