package main

import (
	"bytes"
	"fmt"
	"io"
	"net"

	"github.com/anchormc/anchor/src/enum"
	"github.com/anchormc/anchor/src/packets"
	"github.com/anchormc/anchor/src/protocol"
)

type Client struct {
	ID    uint
	State enum.ClientState
	conn  net.Conn
}

func NewClient(id uint, conn net.Conn) *Client {
	return &Client{
		ID:    id,
		State: enum.ClientStateHandshake,
		conn:  conn,
	}
}

func (c Client) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c Client) ReadPacket(packet packets.Packet) error {
	p, err := c.ReadPacketBuffer()

	if err != nil {
		return err
	}

	return packet.DecodePacket(p)
}

func (c Client) ReadPacketBuffer() (*protocol.PacketBuffer, error) {
	length, err := protocol.ReadVarInt(c.conn)

	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}

	if _, err := io.CopyN(buf, c.conn, int64(length)); err != nil {
		return nil, err
	}

	packetID, err := protocol.ReadVarInt(buf)

	if err != nil {
		return nil, err
	}

	return &protocol.PacketBuffer{
		ID:   packetID,
		Data: buf,
	}, nil
}

func (c *Client) WritePacket(packet packets.Packet) error {
	packetData := &bytes.Buffer{}

	if err := packet.EncodePacket(packetData); err != nil {
		return err
	}

	if err := protocol.WriteVarInt(c.conn, int32(packetData.Len())); err != nil {
		return err
	}

	if _, err := io.Copy(c.conn, packetData); err != nil {
		return err
	}

	return nil
}

func (c *Client) HandlePackets(server *Server) error {
	// [S <- C] Handshake
	{
		p := &packets.ServerboundHandshake{}

		if err := c.ReadPacket(p); err != nil {
			return err
		}

		c.State = enum.ClientState(p.NextState)
	}

	switch c.State {
	case enum.ClientStateStatus:
		{
			if err := Status(server, c); err != nil {
				return err
			}

			return nil
		}
	case enum.ClientStateLogin:
		{
			if err := Login(server, c); err != nil {
				return err
			}

			return nil
		}
	default:
		return fmt.Errorf("client is in unknown state: %d", c.State)
	}
}

func (c *Client) Close() error {
	return c.conn.Close()
}
