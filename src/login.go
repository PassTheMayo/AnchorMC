package main

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"github.com/anchormc/anchor/src/packets"
)

func Login(server *Server, client *Client) error {
	// [S <- C] Login Start
	{
		p := &packets.ServerboundLoginStart{}

		if err := client.ReadPacket(p); err != nil {
			return err
		}
	}

	// [S -> C] Encryption Request
	{
		privatePem := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(server.privateKey),
		})

		publicPem := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(server.publicKey),
		})

		p := &packets.ClientboundEncryptionRequest{
			ServerID:    "",
			PublicKey:   []byte(fmt.Sprintf("-----BEGIN PUBLIC KEY-----\n%s\n-----END PUBLIC KEY-----", base64.RawStdEncoding.EncodeToString(publicPem))),
			VerifyToken: nil, // TODO
		}

		if err := client.WritePacket(p); err != nil {
			return err
		}

		// TODO verify encryption works
	}

	return nil
}
