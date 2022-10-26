package freecache

import (
	"errors"
	"time"

	"github.com/coocood/freecache"
	"github.com/yangjerry110/mytool/cache"
	"github.com/yangjerry110/mytool/conf"
)

type FreeCache struct {
	FreeCache    *freecache.Cache
	FreeCacheVal interface{}
	FreeCacheErr error
	Size         int
	Timer        freecache.Timer
}

func (f *FreeCache) Client(configPath string) cache.CacheInterface {
	yamlConf := conf.YamlConf{YamlFilePath: configPath, Conf: &f}
	err := yamlConf.GetConf()
	if err != nil {
		f.FreeCacheErr = err
		return f
	}

	/**
	 * @step
	 * @check
	 **/
	err = f.CheckConfig()
	if err != nil {
		f.FreeCacheErr = err
		return f
	}
	f.FreeCache = freecache.NewCacheCustomTimer(f.Size, f.Timer)
	return f
}

/**
 * @description: CreateClient
 * @param {int} size
 * @param {freecache.Timer} timer
 * @author: Jerry.Yang
 * @date: 2022-10-25 18:01:15
 * @return {*}
 */
func (f *FreeCache) CreateDefaultClient() cache.CacheInterface {
	f.FreeCache = freecache.NewCacheCustomTimer(f.Size, f.Timer)
	return f
}

/**
 * @description: CheckConfig
 * @author: Jerry.Yang
 * @date: 2022-10-26 15:21:29
 * @return {*}
 */
func (f *FreeCache) CheckConfig() error {
	if f.Size == 0 {
		return errors.New("freecache check : size is not set")
	}
	return nil
}

/**
 * @description: FreeCache
 * @author: Jerry.Yang
 * @date: 2022-10-25 18:06:53
 * @return {*}
 */
func (f *FreeCache) Now() uint32 {
	return uint32(time.Now().Add(60 * time.Minute).Unix())
}
