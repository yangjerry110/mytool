/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 17:15:12
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 18:17:36
 * @Description: yaml conf
 */
package conf

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

/**
 * @description: YamlConf
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:13:58
 * @return {*}
 */
type YamlConf struct {
	FilePath  string
	FileName  string
	FileType  string
	Intervals time.Duration
	Conf      interface{}
}

/**
 * @description: GetHotUpdateConf
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:17:57
 * @return {*}
 */
func (y *YamlConf) GetHotUpdateConf() error {
	conf := &Conf{
		FilePath:  y.FilePath,
		FileName:  y.FileName,
		FileType:  "yaml",
		Intervals: y.Intervals,
		Data:      y.Conf,
	}
	return conf.GetHotUpdateConf()
}

/**
 * @description: Parse()
 * @author: Jerry.Yang
 * @date: 2022-09-22 17:19:26
 * @return {*}
 */
func (y *YamlConf) GetConf() error {
	/**
	 * @step
	 * @获取yaml文件的数据
	 **/
	yamlData, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", y.FilePath, y.FileName))
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @渲染到结构体
	 **/
	decoder := yaml.NewDecoder(bytes.NewReader(yamlData))
	err = decoder.Decode(y.Conf)
	if err != nil {
		return err
	}
	return nil
}
