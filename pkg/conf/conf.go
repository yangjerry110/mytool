/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-23 14:44:20
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-23 15:30:50
 * @Description: conf
 */
package conf

import "github.com/yangjerry110/mytool/conf"

type ConfPkgInterface interface {
	ParseYamlConf(yamlFilePath string, config interface{}) error
}

type ConfPkg struct {
	YamlConfInterface conf.YamlConfInterface
}

/**
 * @description: ParseYamlConf
 * @param {string} yamlFilePath
 * @param {interface{}} config
 * @author: Jerry.Yang
 * @date: 2022-09-23 14:50:45
 * @return {*}
 */
func (y *YamlConfPkg) ParseYamlConf() error {
	return y.CreateYamlConfInstance().Parse()
}
