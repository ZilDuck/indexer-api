package cache

import (
	"github.com/gin-contrib/cache/persistence"
	"go.uber.org/zap"
	"reflect"
	"time"

	"github.com/robfig/go-cache"
)

//InMemoryStore represents the cache with memory persistence
type ReqStore struct {
	cache.Cache
}

// NewInMemoryStore returns a InMemoryStore
func NewReqStore(defaultExpiration time.Duration) *ReqStore {
	return &ReqStore{*cache.New(defaultExpiration, time.Minute)}
}

// Get (see CacheStore interface)
func (c *ReqStore) Get(key string, value interface{}) error {
	val, found := c.Cache.Get(key)
	if !found {
		return persistence.ErrCacheMiss
	}

	v := reflect.ValueOf(value)
	if v.Type().Kind() == reflect.Ptr && v.Elem().CanSet() {
		v.Elem().Set(reflect.ValueOf(val))
		zap.S().Infof("Cache Get: %s", key)
		return nil
	}
	return persistence.ErrNotStored
}

// Set (see CacheStore interface)
func (c *ReqStore) Set(key string, value interface{}, expires time.Duration) error {
	// NOTE: go-cache understands the values of DEFAULT and FOREVER
	zap.S().Infof("Cache Set: %s", key)
	c.Cache.Set(key, value, expires)
	return nil
}

// Add (see CacheStore interface)
func (c *ReqStore) Add(key string, value interface{}, expires time.Duration) error {
	err := c.Cache.Add(key, value, expires)
	if err == cache.ErrKeyExists {
		return persistence.ErrNotStored
	}
	return err
}

// Replace (see CacheStore interface)
func (c *ReqStore) Replace(key string, value interface{}, expires time.Duration) error {
	if err := c.Cache.Replace(key, value, expires); err != nil {
		return persistence.ErrNotStored
	}
	return nil
}

// Delete (see CacheStore interface)
func (c *ReqStore) Delete(key string) error {
	if found := c.Cache.Delete(key); !found {
		return persistence.ErrCacheMiss
	}
	return nil
}

// Increment (see CacheStore interface)
func (c *ReqStore) Increment(key string, n uint64) (uint64, error) {
	newValue, err := c.Cache.Increment(key, n)
	if err == cache.ErrCacheMiss {
		return 0, persistence.ErrCacheMiss
	}
	return newValue, err
}

// Decrement (see CacheStore interface)
func (c *ReqStore) Decrement(key string, n uint64) (uint64, error) {
	newValue, err := c.Cache.Decrement(key, n)
	if err == cache.ErrCacheMiss {
		return 0, persistence.ErrCacheMiss
	}
	return newValue, err
}

// Flush (see CacheStore interface)
func (c *ReqStore) Flush() error {
	c.Cache.Flush()
	return nil
}
