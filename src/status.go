package main

import (
	"encoding/hex"
	"fmt"

	"github.com/anchormc/anchor/src/chat"
	"github.com/anchormc/anchor/src/packets"
	"github.com/anchormc/anchor/src/types"
)

func Status(server *Server, client *Client) error {
	// [S <- C] Status Request
	{
		p := &packets.ServerboundStatusRequest{}

		if err := client.ReadPacket(p); err != nil {
			return err
		}
	}

	// [S -> C] Status Response
	{
		var icon *string = nil

		if server.statusIcon != nil {
			// TODO test that this works

			icon = pointerOf(fmt.Sprintf("data:image/png;base64,%s", hex.EncodeToString(server.statusIcon)))
		}

		p := &packets.ClientboundStatusResponse{
			Payload: types.StatusResponse{
				Version: types.StatusResponseVersion{
					Name:     "1.20.2",
					Protocol: 764,
				},
				Players: types.StatusResponsePlayers{
					Online: 0,
					Max:    64,
					Sample: nil,
				},
				Description: chat.Chat{
					Text: "Hello, world!",
				},
				Favicon: icon,
			},
		}

		if err := client.WritePacket(p); err != nil {
			return err
		}
	}

	pingPongPacket := &packets.BidirectionalPingPong{}

	// [S <- C] Ping
	{
		if err := client.ReadPacket(pingPongPacket); err != nil {
			return err
		}
	}

	// [S -> C] Pong
	{
		if err := client.WritePacket(pingPongPacket); err != nil {
			return err
		}
	}

	return nil
}
