package client

import (
	"google.golang.org/grpc"

	"github.com/opensourceways/repo-owners-cache/protocol"
)

func NewClient(endpoint string) (*Client, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:             conn,
		RepoOwnersClient: protocol.NewRepoOwnersClient(conn),
	}, nil
}

type Client struct {
	protocol.RepoOwnersClient

	conn *grpc.ClientConn
}

func (c *Client) Disconnect() error {
	return c.conn.Close()
}
