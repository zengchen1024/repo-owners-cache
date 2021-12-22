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

// The second return value means whether can retry
func (c *cacheEntry) init(cli *filecache.SDK, log *logrus.Entry) (RepoOwner, bool) {
	select {
	case c.start <- empty:
		defer func() {
			<-c.start
		}()

		if d := c.getOwner(); d != nil {
			return d, false
		}

		r, err := loadOwners(cli, c.branch, log)
		if err != nil {
			return nil, false
		}

		if r.IsEmpty() {
			return nil, false
		}

		c.setOwner(r)

		return r, false

	default:
		return nil, true
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

func loadOwners(cli *filecache.SDK, b RepoBranch, log *logrus.Entry) (*RepoOwnerInfo, error) {
	v, err := cli.GetFiles(b, "OWNERS", false)
	if err != nil {
		log.Errorf("load file for branch:%s, err:%s", branchToKey(b), err.Error())
		return nil, err
	}

	o := newRepoOwnerInfo()
	k := branchToKey(b)

	for _, item := range v.Files {
		if err := o.parseOwnerConfig(item.Dir(), item.Content, log); err != nil {
			log.Errorf(
				"parse file:%s of branch:%s, err:%s",
				item.Dir(), k, err.Error(),
			)
		}
	}

	return o, nil
}
