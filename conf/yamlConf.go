/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 17:15:12
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 17:18:39
 * @Description: yaml conf
 */
package conf

import (
	"bytes"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type YamlConfInterface interface {
	Parse() error
}

type YamlConf struct {
	YamlFilePath string
	Conf         interface{}
}

/**
 * @description: Parse()
 * @author: Jerry.Yang
 * @date: 2022-09-22 17:19:26
 * @return {*}
 */
func (y *YamlConf) Parse() error {
	/**
	 * @step
	 * @获取yaml文件的数据
	 **/
	yamlData, err := ioutil.ReadFile(y.YamlFilePath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @渲染到结构体
	 **/
	decoder := yaml.NewDecoder(bytes.NewReader(yamlData))
	decoder.Decode(y.Conf)
	return nil
}
