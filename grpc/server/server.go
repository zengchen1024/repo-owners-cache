package server

import (
	"net"

	"github.com/opensourceways/community-robot-lib/interrupts"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/opensourceways/repo-owners-cache/cache"
	"github.com/opensourceways/repo-owners-cache/protocol"
)

func Start(port string, c *cache.Cache) error {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	protocol.RegisterRepoOwnersServer(server, &repoOwnersServer{c: c})

	return run(server, listen)
}

func run(server *grpc.Server, listen net.Listener) error {
	defer interrupts.WaitForGracefulShutdown()

	interrupts.OnInterrupt(func() {
		logrus.Errorf("grpc server exit...")
		server.Stop()
	})

	return server.Serve(listen)
}
