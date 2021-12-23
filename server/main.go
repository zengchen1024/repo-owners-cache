package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/opensourceways/community-robot-lib/logrusutil"
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/repo-owners-cache/cache"
	"github.com/opensourceways/repo-owners-cache/grpc/server"
)

type options struct {
	port     string
	endpoint string
}

func (o *options) addFlags(fs *flag.FlagSet) {
	fs.StringVar(&o.port, "port", "8888", "Port to listen on.")
	fs.StringVar(&o.endpoint, "endpoint", "", "The endpoint of repo file cache")
}

func (o *options) validate() error {
	if o.endpoint == "" {
		return fmt.Errorf("missing endpoint")
	}

	v, err := url.Parse(o.endpoint)
	if err != nil {
		return err
	}
	o.endpoint = v.String()

	return nil
}

func gatherOptions(fs *flag.FlagSet, args ...string) options {
	var o options
	o.addFlags(fs)
	_ = fs.Parse(args)
	return o
}

func main() {
	logrusutil.ComponentInit("repo-owners-cache")

	o := gatherOptions(flag.NewFlagSet(os.Args[0], flag.ExitOnError), os.Args[1:]...)
	if err := o.validate(); err != nil {
		logrus.WithError(err).Fatal("Invalid options")
	}

	c := cache.NewCache(o.endpoint, logrus.NewEntry(logrus.StandardLogger()))

	if err := server.Start(":"+o.port, c); err != nil {
		logrus.WithError(err).Fatal("Error to start grpc server.")
	}
}
