package xkv

import (
	"log"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	// getAndDelScript 获取并删除key所关联的值lua脚本
	getAndDelScript = `local current = redis.call('GET', KEYS[1]);
if (current) then
    redis.call('DEL', KEYS[1]);
end
return current;`
)

type Store struct {
	kv.Store

	Redis *redis.Redis
}

func NewStore(c kv.KvConf) *Store {
	if len(c) == 0 || cache.TotalWeights(c) <= 0 {
		log.Fatal("no cache nodes")
	}
	cn := redis.MustNewRedis(c[0].RedisConf)
	return &Store{
		Store: kv.NewStore(c),
		Redis: cn,
	}
}
