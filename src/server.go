package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"image/png"
	"os"

	"github.com/anchormc/anchor/src/logger"
)

type Server struct {
	Config          *Config
	socket          *Socket
	Clients         []*Client
	clientIDCounter uint
	statusIcon      []byte
	privateKey      *rsa.PrivateKey
	publicKey       *rsa.PublicKey
}

func NewServer() *Server {
	return &Server{
		Config:          DefaultConfig,
		socket:          NewSocket(),
		clientIDCounter: 0,
	}
}

func (s *Server) Initialize() error {
	key, err := rsa.GenerateKey(rand.Reader, 1024)

	if err != nil {
		return err
	}

	s.privateKey = key
	s.publicKey = key.Public().(*rsa.PublicKey)

	if err := s.Config.ReadFile("config.yml"); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}

		if err = s.Config.WriteFile("config.yml"); err != nil {
			return err
		}

		logger.Warn("config.yml did not exist, wrote default values")
	} else {
		logger.Info("Successfully read configuration file")
	}

	if data, err := os.ReadFile("server-icon.png"); err == nil {
		img, err := png.Decode(bytes.NewReader(data))

		if err != nil {
			return err
		}

		size := img.Bounds().Size()

		if size.X != 64 || size.Y != 64 {
			logger.Fatalf("server-icon.png has invalid dimensions (expected=64x64, received=%dx%d)", size.X, size.Y)
		}

		buf := &bytes.Buffer{}

		if err = png.Encode(buf, img); err != nil {
			return err
		}

		s.statusIcon = buf.Bytes()
	}

	return nil
}

func (s *Server) Start() error {
	if err := s.socket.Listen(s.Config.Host, s.Config.Port); err != nil {
		return err
	}

	logger.Infof("Listening on %s:%d", s.Config.Host, s.Config.Port)

	go s.handleConnections()

	return nil
}

func (s *Server) addClient(client *Client) {
	for _, c := range s.Clients {
		if c.ID != client.ID {
			continue
		}

		logger.Fatalf("A client already exists with the ID of %d", client.ID)
	}

	s.Clients = append(s.Clients, client)
}

func (s *Server) removeClient(id uint) {
	for k, c := range s.Clients {
		if c.ID != id {
			continue
		}

		s.Clients = append(s.Clients[0:k], s.Clients[k+1:]...)
	}
}

func (s *Server) handleConnections() error {
	for {
		conn, err := s.socket.Accept()

		if err != nil {
			logger.Errorf("Error while accepting socket connection: %s", err)

			continue
		}

		s.clientIDCounter++

		client := NewClient(s.clientIDCounter, conn)
		s.addClient(client)

		go s.handleClient(client)

		logger.Infof("Received a new connection from %s", client.RemoteAddr())
	}
}

func (s *Server) handleClient(client *Client) {
	if err := client.HandlePackets(s); err != nil {
		logger.Errorf("Error from client %d: %s", client.ID, err)
	}

	s.removeClient(client.ID)

	if err := client.Close(); err != nil {
		logger.Error(err)
	}
}

func (s *Server) Close() error {
	if err := s.socket.Close(); err != nil {
		return err
	}

	return nil
}
