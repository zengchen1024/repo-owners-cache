package cache

import (
	"fmt"
	"sync"

	"github.com/opensourceways/repo-file-cache/models"
	filecache "github.com/opensourceways/repo-file-cache/sdk"
	"github.com/sirupsen/logrus"
)

var empty = struct{}{}

type RepoBranch = models.Branch

func branchToKey(b RepoBranch) string {
	return fmt.Sprintf("%s/%s/%s:%s", b.Platform, b.Org, b.Repo, b.Branch)
}

type cacheEntry struct {
	owner  *RepoOwnerInfo
	branch RepoBranch

	lock  sync.RWMutex
	start chan struct{}
}

func newCacheEntry(b RepoBranch) *cacheEntry {
	return &cacheEntry{
		start:  make(chan struct{}, 1),
		branch: b,
	}
}

func (c *cacheEntry) init(cli *filecache.SDK, log *logrus.Entry) (RepoOwner, error) {
	select {
	case c.start <- empty:
		defer func() {
			<-c.start
		}()

		if d := c.getOwner(); d != nil {
			return d, nil
		}

		l := log.WithField("init for branch", c.getBranchKey())

		v, err := cli.GetFiles(c.branch, "OWNERS", false)
		if err != nil {
			l.Errorf("load file, err:%s", err.Error())

			return nil, err
		}

		o := newRepoOwnerInfo()

		for _, f := range v.Files {
			o.parseOwnerConfig(f.Dir(), f.Content, f.SHA, l)
		}

		if o.isEmpty() {
			return nil, nil
		}

		c.setOwner(o)

		return o, nil

	default:
		return nil, fmt.Errorf("no chance to init repo owner")
	}
}

func (c *cacheEntry) refresh(cli *filecache.SDK, log *logrus.Entry) error {
	select {
	case c.start <- empty:
		defer func() {
			<-c.start
		}()

		l := log.WithField("refresh for branch", c.getBranchKey())

		owner := c.getOwner()
		if owner == nil {
			l.Error("it should init instead of refreshing")

			return nil
		}

		v, err := cli.GetFiles(c.branch, "OWNERS", false)
		if err != nil {
			l.Errorf("load file, err:%s", err.Error())

			return err
		}

		no := newRepoOwnerInfo()

		for i := range v.Files {
			f := &v.Files[i]
			dir := f.Dir()

			if owner.getFileSHA(dir) != f.SHA {
				no.parseOwnerConfig(dir, f.Content, f.SHA, l)
			}
		}

		if no.isEmpty() {
			return nil
		}

		owner.copyOwnerFiles(no)
		c.setOwner(no)

		return nil

	default:
		return fmt.Errorf("no chance to refresh repo owner")
	}
}

func (c *cacheEntry) getOwner() *RepoOwnerInfo {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.owner
}

func (c *cacheEntry) setOwner(d *RepoOwnerInfo) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.owner = d
}

func (c *cacheEntry) getBranchKey() string {
	return branchToKey(c.branch)
}
