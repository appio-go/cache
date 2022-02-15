package cache

import (
	"github.com/patrickmn/go-cache"
	"os"
	"strconv"
	"time"
)

type Singleton interface {
	Get(string) (interface{}, bool)
	Set(string, interface{})
}

type singleton struct {
	cache    *cache.Cache
	lifetime time.Duration
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
		instance.lifetime = 5 * time.Minute

		lifetimeString, _ := os.LookupEnv("CACHE_LIFETIME")
		lifetime, err := strconv.ParseInt(lifetimeString, 10, 32)

		if err == nil {
			instance.lifetime = time.Duration(lifetime) * time.Minute
		}

		instance.cache = cache.New(instance.lifetime, instance.lifetime)
	}
	return instance
}
func (s *singleton) Get(st string) (interface{}, bool) {
	return s.cache.Get(st)
}

func (s *singleton) Set(st string, i interface{}) {
	//При добавление чего нить в кеш чистим все просроченное. Чтобы поитогу не протерять всю память
	s.cache.DeleteExpired()
	s.cache.Set(st, i, instance.lifetime)
}
