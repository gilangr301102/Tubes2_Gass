package utils

import lru "github.com/hashicorp/golang-lru"

const CacheSize = 200000

var Articles *lru.Cache

func init() {
	cache, err := lru.New(CacheSize)
	PanicIfError(err)
	Articles = cache
}
