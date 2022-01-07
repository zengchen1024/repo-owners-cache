package cache

import (
	"sync"
	"time"

	"github.com/opensourceways/community-robot-lib/utils"
	filecache "github.com/opensourceways/repo-file-cache/sdk"
	"github.com/sirupsen/logrus"
)

type Cache struct {
	cli *filecache.SDK
	log *logrus.Entry
	t   utils.Timer

	lock sync.RWMutex
	data map[string]*cacheEntry
}

func NewCache(endpoint string, log *logrus.Entry) *Cache {
	return &Cache{
		log:  log,
		cli:  filecache.NewSDK(endpoint, 3),
		data: make(map[string]*cacheEntry),
	}
}

func (c *Cache) SyncPerDay(start time.Duration) func() {
	c.t = utils.NewTimer()

	go func(delay time.Duration) {
		c.t.Start(c.refresh, 24*time.Hour, delay)
	}(start)

	return c.t.Stop
}

func (c *Cache) get(k string) *cacheEntry {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.data[k]
}

func (c *Cache) getOrNewAnEntry(b RepoBranch) *cacheEntry {
	k := branchToKey(b)

	if entry := c.get(k); entry != nil {
		return entry
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	if e, ok := c.data[k]; ok {
		return e
	}

	e := newCacheEntry(b)
	c.data[k] = e

	return e
}

func (c *Cache) LoadRepoOwners(b RepoBranch) (RepoOwner, error) {
	e := c.getOrNewAnEntry(b)

	var r *RepoOwnerInfo
	var err error

	for i := 0; i < 10; i++ {
		if r = e.getOwner(); r != nil {
			return r, nil
		}

		if r, err = e.init(c.cli, c.log); err == nil {
			if r != nil {
				return r, nil
			}

			return nil, nil
		}

		time.Sleep(time.Second)
	}

	if r = e.getOwner(); r != nil {
		return r, nil
	}

	return nil, err
}

func (c *Cache) refresh() {
	c.log.Info("refresh starts")

	c.lock.RLock()
	all := make([]string, 0, len(c.data))
	for k := range c.data {
		all = append(all, k)
	}
	c.lock.RUnlock()

	f := func(k string) {
		e := c.get(k)
		if e == nil {
			return
		}

		for i := 0; i < 10; i++ {
			if err := e.refresh(c.cli, c.log); err == nil {
				return
			}

			time.Sleep(time.Second)
		}

		e.refresh(c.cli, c.log)
	}

	for _, k := range all {
		f(k)
	}

	c.log.Info("refresh ends")
}
