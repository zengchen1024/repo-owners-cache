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

		k := c.getBranchKey()

		v, err := cli.GetFiles(c.branch, "OWNERS", false)
		if err != nil {
			log.Errorf(
				"load file for branch:%s, err:%s", k, err.Error(),
			)
			return nil, err
		}

		o := newRepoOwnerInfo()

		for i := range v.Files {
			parseOwnerConfig(k, o, &v.Files[i], log)
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

		owner := c.getOwner()
		if owner == nil {
			return fmt.Errorf("unable to refresh")
		}

		k := c.getBranchKey()

		v, err := cli.GetFiles(c.branch, "OWNERS", false)
		if err != nil {
			log.Errorf(
				"load file for branch:%s, err:%s", k, err.Error(),
			)
			return err
		}

		no := newRepoOwnerInfo()
		for i := range v.Files {
			f := &v.Files[i]
			dir := f.Path.Dir()

			if owner.getFileSHA(dir) != f.SHA {
				parseOwnerConfig(k, no, f, log)
			}
		}

		if no.isEmpty() {
			return nil
		}

		owner.copyOwnerFiles(no)

		c.setOwner(no)

		return nil

	default:
		return fmt.Errorf("no chance to init repo owner")
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

func parseOwnerConfig(branch string, o *RepoOwnerInfo, file *models.File, log *logrus.Entry) {
	if err := o.parseOwnerConfig(file.Dir(), file.Content, file.SHA, log); err != nil {
		log.Errorf(
			"parse file:%s of branch:%s, err:%s",
			file.Dir(), branch, err.Error(),
		)
	}
}
