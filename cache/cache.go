package cache

import (
	"sync"
	"time"

	filecache "github.com/opensourceways/repo-file-cache/sdk"
	"github.com/sirupsen/logrus"
)

type Cache struct {
	cli *filecache.SDK
	log *logrus.Entry

	dataLock sync.RWMutex
	data     map[string]*cacheEntry
}

func NewCache(endpoint string, log *logrus.Entry) *Cache {
	return &Cache{
		log:  log,
		cli:  filecache.NewSDK(endpoint, 3),
		data: make(map[string]*cacheEntry),
	}
}

func (c *Cache) get(k string) *cacheEntry {
	c.dataLock.RLock()
	defer c.dataLock.RUnlock()

	return c.data[k]
}

func (c *Cache) getOrNewAnEntry(b RepoBranch) *cacheEntry {
	k := branchToKey(b)

	if entry := c.get(k); entry != nil {
		return entry
	}

	c.dataLock.Lock()
	defer c.dataLock.Unlock()

	if e, ok := c.data[k]; ok {
		return e
	}

	e := newCacheEntry(b)
	c.data[k] = e

	return e
}

func (c *Cache) LoadRepoOwners(b RepoBranch) RepoOwner {
	e := c.getOrNewAnEntry(b)

	if r := e.getData(); r != nil {
		return r
	}

	for i := 0; i < 10; i++ {
		if r, b := e.init(c.cli, c.log); !b {
			return r
		}

		time.Sleep(time.Second)
	}

	r, _ := e.init(c.cli, c.log)

	return r
}
