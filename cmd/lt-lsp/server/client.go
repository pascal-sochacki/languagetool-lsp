package server

import "go.lsp.dev/protocol"

type Client struct {
	target protocol.Client
}

func NewClient() (Client, func(protocol.Client)) {
	a := Client{}

	return a, func(c protocol.Client) {
		a.target = c
	}

}
