package lookup

import (
	"errors"
	"strings"
	"sync"
	"time"
)

// Could implement go-cache (https://github.com/patrickmn/go-cache) but lets create our own simple one for IPs and
// their associated info (https://hackernoon.com/in-memory-caching-in-golang)

type cachedIp struct {
	IpData
	expireAtTimestamp int64
}

type LocalCache struct {
	stop chan struct {}

	wg sync.WaitGroup
	mu sync.RWMutex
	ips map[string]cachedIp
}

// Create a new in memory cache
func NewLocalCache(cleanupInterval time.Duration) *LocalCache {
	lc := &LocalCache{
		ips: make(map[string]cachedIp),
		stop: make(chan struct{}),
	}

	lc.wg.Add(1)
	go func(cleanupInterval time.Duration) {
		defer lc.wg.Done()

		lc.cleanupLoop(cleanupInterval)
	}(cleanupInterval)

	return lc
}

// Filter cache for specific values
func (lc *LocalCache) Filter(loc IpLocation) []IpData {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	city := strings.ToLower(loc["city"])
	country := strings.ToLower(loc["country"])

	a := make([]IpData, 0)
	switch {
	case city == "":
		for _, v := range lc.ips {
			ipCou := strings.ToLower(v.IpData.Country)
			if ipCou == country {
				a = append(a, v.IpData)
			}
		}
		case country == "":
		for _, v := range lc.ips {
			ipCit := strings.ToLower(v.IpData.City)
			if ipCit == city {
				a = append(a, v.IpData)
			}
		}
		default:
		for _, v := range lc.ips {
			ipCit := strings.ToLower(v.IpData.City)
			ipCou := strings.ToLower(v.IpData.Country)
			if ipCit == city && ipCou == country {
				a = append(a, v.IpData)
			}
		}
	}

	return a
}

// Read all values in cache
func (lc *LocalCache) ReadAll() []IpData {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	a := make([]IpData, len(lc.ips))
	i := 0
	for _, v := range lc.ips {
		a[i] = v.IpData
		i++
	}

	return a
}

// Read key from cache
func (lc *LocalCache) ReadCache(i string) (IpData, error) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	ci, ok := lc.ips[i]
	if !ok {
		return IpData{}, errIpNotInCache
	}

	return ci.IpData, nil
}

var errIpNotInCache = errors.New("This IP is not in the cache")

// Update the cache with a new key-value
func (lc *LocalCache) UpdateCache(ipd IpData) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.ips[ipd.Ip] = cachedIp {
		IpData: ipd,
		expireAtTimestamp: time.Now().Unix(),
	}	
}

// Cleanup expired key-values
func (lc *LocalCache) cleanupLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-lc.stop:
			return
		case <-t.C:
			lc.mu.Lock()
			for ip, cu := range lc.ips {
				if cu.expireAtTimestamp <= time.Now().Unix() {
					delete(lc.ips, ip)
				}
			}
			lc.mu.Unlock()
		}
	}
}

// Stop cleaning the cache
func (lc *LocalCache) stopCleanup() {
	close(lc.stop)
	lc.wg.Wait()
}

// Delete key-value from cache
func (lc *LocalCache) delete(ip string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	delete(lc.ips, ip)
}

