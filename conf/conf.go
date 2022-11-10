/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-26 15:15:25
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 18:18:24
 * @Description: conf
 */
package conf

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

type ConfInterface interface {
	GetHotUpdateConf() error
	GetConf() error
}

/**
 * @description: 需要通知更新的对象
 * @author: Jerry.Yang
 * @date: 2022-11-10 16:58:23
 * @return {*}
 */
type Notifyer interface {
	Callback(c *Conf) error
}

/**
 * @description: Conf
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:13:42
 * @return {*}
 */
type Conf struct {
	FilePath       string
	FileName       string
	FileType       string
	Intervals      time.Duration
	Data           interface{}
	LastModityTime time.Time
	NotifyerList   []Notifyer
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
 * @description: GetHotUpdateConf
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:13:19
 * @return {*}
 */
func (c *Conf) GetHotUpdateConf() error {

	/**
	 * @step
	 * @获取配置
	 **/
	err := c.GetConf()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @添加观察者
	 **/
	c.AddNotifyer(configStore)

	/**
	 * @step
	 * @添加一个协程，热更新配置文件
	 **/
	go c.GetConfTimer()
	return nil
}

/**
 * @description: GetConf
 * @author: Jerry.Yang
 * @date: 2022-11-10 17:15:51
 * @return {*}
 */
func (c *Conf) GetConf() error {
	/**
	 * @step
	 * @根据type，获取文件内容
	 **/
	switch c.FileType {
	case "yaml":
		yamlConf := YamlConf{FilePath: c.FilePath, FileName: c.FileName, Conf: c.Data}
		yamlConfErr := yamlConf.GetConf()
		if yamlConfErr != nil {
			return yamlConfErr
		}
	default:
	}
	return nil
}

/**
 * @description: GetConfTimer
 * @author: Jerry.Yang
 * @date: 2022-11-10 17:15:57
 * @return {*}
 */
func (c *Conf) GetConfTimer() error {

	/**
	 * @step
	 * @创建一个定时器
	 **/
	timeTickers := time.NewTicker(time.Second * c.Intervals)

	/**
	 * @step
	 * @定时读取配置
	 **/
	for range timeTickers.C {

		/**
		 * @step
		 * @打开文件
		 **/
		fileObj, err := os.Open(fmt.Sprintf("%s/%s", c.FilePath, c.FileName))
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
		if fileInfo.ModTime().Unix() > c.LastModityTime.Unix() {
			/**
			 * @step
			 * @当发生更新的时候
			 * @赋值时间
			 **/
			c.LastModityTime = fileInfo.ModTime()

			/**
			 * @step
			 * @当更新的时候，重新解析
			 **/
			err = c.GetConf()
			if err != nil {
				return err
			}

			/**
			 * @step
			 * @通知订阅者
			 **/
			for _, notifyer := range c.NotifyerList {
				notifyer.Callback(c)
			}
		}

	}
	//timeTickers.Stop()
	return nil
}

/**
 * @description: AddNotifyer
 * @param {Notifyer} n
 * @author: Jerry.Yang
 * @date: 2022-11-10 17:56:09
 * @return {*}
 */
func (c *Conf) AddNotifyer(n Notifyer) {
	c.NotifyerList = append(c.NotifyerList, n)
}

/**
 * @description: Callback
 * @param {*Conf} conf
 * @author: Jerry.Yang
 * @date: 2022-11-10 17:56:17
 * @return {*}
 */
func (c *ConfigStore) Callback(conf *Conf) error {
	configStore.Config.Store(conf.Data)
	return nil
}
