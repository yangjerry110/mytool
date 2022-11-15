/*
 * @Author: Jerry.Yang
 * @Date: 2022-11-11 15:14:45
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-15 14:26:56
 * @Description: init
 */
package conf

import (
	"crypto/md5"
	"errors"
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

type InitConf struct {
	FilePath       string
	FileName       string
	FileType       string
	Intervals      time.Duration
	LastModityTime time.Time
	NotifyerList   []Notifyer
}

/**
 * @description: 需要通知更新的对象
 * @author: Jerry.Yang
 * @date: 2022-11-10 16:58:23
 * @return {*}
 */
type Notifyer interface {
	Callback(c *InitConf, configData interface{}) error
}

/**
 * @description: ConfigStore
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:13:33
 * @return {*}
 */
type ConfigStore struct {
	Config atomic.Value
}

/**
 * @description: ConfigStore
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:13:26
 * @return {*}
 */
var configStore = &ConfigStore{}

/**
 * @description: Init
 * @param {interface{}} conf
 * @author: Jerry.Yang
 * @date: 2022-11-11 16:02:02
 * @return {*}
 */
func (i *InitConf) Init(conf interface{}) ConfInterface {

	/**
	 * @step
	 * @获取一遍数据
	 **/
	i.GetConf(conf)

	/**
	 * @step
	 * @添加观察者
	 **/
	i.AddNotifyer(&ConfigStore{})

	/**
	 * @step
	 * @开启一个协程，刷新配置
	 **/
	go i.ReloadConf(conf)
	return i
}

/**
 * @description: GetNewConf
 * @author: Jerry.Yang
 * @date: 2022-11-11 16:01:53
 * @return {*}
 */
func (i *InitConf) GetNewConf() error {
	return nil
}

/**
 * @description: GetParseConf
 * @author: Jerry.Yang
 * @date: 2022-11-11 16:01:38
 * @return {*}
 */
func (i *InitConf) GetParseConf() interface{} {
	atomicData := configStore.Config.Load().(map[string]interface{})
	return atomicData[i.GetStoreKey()]
}

/**
 * @description: GetConf
 * @param {interface{}} conf
 * @author: Jerry.Yang
 * @date: 2022-11-11 16:01:45
 * @return {*}
 */
func (i *InitConf) GetConf(conf interface{}) error {
	/**
	 * @step
	 * @根据type，获取文件内容
	 **/
	switch i.FileType {
	case "yaml":
		yamlConf := YamlConf{FilePath: i.FilePath, FileName: i.FileName, Conf: conf}
		yamlConfErr := yamlConf.GetConf(conf)
		if yamlConfErr != nil {
			return yamlConfErr
		}
	default:
		return errors.New("GetConf Err : no match fileType")
	}

	/**
	 * @step
	 * @保存数据
	 **/
	md5Key := i.GetStoreKey()
	configStore.Config.Store(map[string]interface{}{md5Key: conf})
	return nil
}

/**
 * @description: ReloadConf
 * @param {interface{}} conf
 * @author: Jerry.Yang
 * @date: 2022-11-11 16:02:15
 * @return {*}
 */
func (i *InitConf) ReloadConf(conf interface{}) error {
	/**
	 * @step
	 * @创建一个定时器
	 **/
	timeTickers := time.NewTicker(time.Second * i.Intervals)

	/**
	 * @step
	 * @定时读取配置
	 **/
	for range timeTickers.C {

		/**
		 * @step
		 * @打开文件
		 **/
		fileObj, err := os.Open(fmt.Sprintf("%s/%s", i.FilePath, i.FileName))
		if err != nil {
			//timeTickers.Stop()
			return err
		}
		defer fileObj.Close()

		/**
		 * @step
		 * @获取文件详情
		 **/
		fileInfo, err := fileObj.Stat()
		if err != nil {
			//timeTickers.Stop()
			return err
		}

		/**
		 * @step
		 * @获取当前文件的修改时间
		 * @当发生更新的时候，执行
		 **/
		if fileInfo.ModTime().Unix() > i.LastModityTime.Unix() {
			/**
			 * @step
			 * @当发生更新的时候
			 * @赋值时间
			 **/
			i.LastModityTime = fileInfo.ModTime()

			/**
			 * @step
			 * @当更新的时候，重新解析
			 **/
			err = i.GetConf(conf)
			if err != nil {
				return err
			}

			/**
			 * @step
			 * @通知订阅者
			 **/
			for _, notifyer := range i.NotifyerList {
				notifyer.Callback(i, conf)
			}
		}

	}
	//timeTickers.Stop()
	return nil
}

/**
 * @description: GetStoreKey
 * @author: Jerry.Yang
 * @date: 2022-11-11 15:04:21
 * @return {*}
 */
func (i *InitConf) GetStoreKey() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s/%s", i.FilePath, i.FileName))))
}

/**
 * @description: AddNotifyer
 * @param {Notifyer} n
 * @author: Jerry.Yang
 * @date: 2022-11-10 17:56:09
 * @return {*}
 */
func (i *InitConf) AddNotifyer(n Notifyer) {
	i.NotifyerList = append(i.NotifyerList, n)
}

/**
 * @description: Callback
 * @param {*InitConf} conf
 * @param {interface{}} configData
 * @author: Jerry.Yang
 * @date: 2022-11-11 16:02:26
 * @return {*}
 */
func (c *ConfigStore) Callback(conf *InitConf, configData interface{}) error {

	/**
	 * @step
	 * @生成md5
	 **/
	md5Key := conf.GetStoreKey()
	storeVal := map[string]interface{}{md5Key: configData}
	configStore.Config.Store(storeVal)
	return nil
}
