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
	lock   sync.RWMutex
	start  chan struct{}
	owner  RepoOwner
	branch RepoBranch
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

		v, err := cli.GetFiles(c.branch, "OWNERS", false)
		if err != nil {
			log.Errorf(
				"load file for branch:%s, err:%s",
				branchToKey(c.branch), err.Error(),
			)
			return nil, err
		}

		r := loadOwners(c.branch, v.Files, log)
		if r.isEmpty() {
			return nil, nil
		}

		c.setOwner(r)

		return r, nil

	default:
		return nil, fmt.Errorf("no chance to init repo owner")
	}
}

func (c *cacheEntry) getOwner() RepoOwner {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.owner
}

func (c *cacheEntry) setOwner(d RepoOwner) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.owner = d
}

func loadOwners(b RepoBranch, files []models.File, log *logrus.Entry) *RepoOwnerInfo {
	o := newRepoOwnerInfo()
	k := branchToKey(b)

	for _, item := range files {
		if err := o.parseOwnerConfig(item.Dir(), item.Content, log); err != nil {
			log.Errorf(
				"parse file:%s of branch:%s, err:%s",
				item.Dir(), k, err.Error(),
			)
		}
	}

	return o
}
