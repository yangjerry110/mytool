/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 17:19:49
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 17:20:56
 * @Description: yaml conf
 */
package conf

import "github.com/yangjerry110/mytool/conf"

type YamlConfInterface interface{}

type YamlConf struct{}

/**
 * @description: ParseYaml
 * @param {string} yamlFilePath
 * @param {interface{}} config
 * @author: Jerry.Yang
 * @date: 2022-09-22 17:21:40
 * @return {*}
 */
func ParseYaml(yamlFilePath string, config interface{}) error {
	yamlConfObj := conf.YamlConf{YamlFilePath: yamlFilePath, Conf: &config}
	err := yamlConfObj.Parse()
	return err
}
