package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/opensourceways/community-robot-lib/logrusutil"
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/repo-owners-cache/cache"
	"github.com/opensourceways/repo-owners-cache/grpc/server"
)

type options struct {
	port      string
	endpoint  string
	startTime string
}

func (o *options) addFlags(fs *flag.FlagSet) {
	fs.StringVar(&o.port, "port", "8888", "Port to listen on.")
	fs.StringVar(&o.endpoint, "endpoint", "", "The endpoint of repo file cache")
	fs.StringVar(&o.startTime, "start-time", "01:00", "Time to synchronize repo file for the first time. The format is Hour:Minute")
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

	if _, _, err := o.parseStartTime(); err != nil {
		return err
	}

	return nil
}

func (o *options) parseStartTime() (int, int, error) {
	format := "2006-01-02T"
	t, err := time.Parse(format+"15:04", format+o.startTime)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid start_time: %s", err.Error())
	}

	return t.Hour(), t.Minute(), nil
}

func (o *options) getStartTime() time.Duration {
	seconds := func(h, m int) int {
		return h*3600 + m*60
	}

	now := time.Now()
	t0 := seconds(now.Hour(), now.Minute())

	h, m, _ := o.parseStartTime()
	t1 := seconds(h, m)

	diff := t1 - t0
	if diff >= 0 {
		return time.Duration(diff) * time.Second
	}
	return time.Duration(seconds(24, 0)+diff) * time.Second
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

	stop := c.SyncPerDay(o.getStartTime())

	defer stop()

	if err := server.Start(":"+o.port, c); err != nil {
		logrus.WithError(err).Fatal("Error to start grpc server.")
	}
}
