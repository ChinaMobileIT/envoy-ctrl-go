package cache

import (
	"github.com/ChinaMobileIT/envoy-ctrl-go/internal/example"
	cachev3 "github.com/ChinaMobileIT/envoy-ctrl-go/pkg/cache/v3"
)

var (
	cache cachev3.SnapshotCache
)

func NewCache(l example.Logger) cachev3.SnapshotCache {
	// Create a Cache
	cache = cachev3.NewSnapshotCache(false, cachev3.IDHash{}, l)
	return cache
}
func GetCache() cachev3.SnapshotCache {
	return cache
}
