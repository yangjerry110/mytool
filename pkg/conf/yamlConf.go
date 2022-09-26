/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 17:19:49
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 18:39:57
 * @Description: yaml conf
 */
package conf

import "github.com/yangjerry110/mytool/conf"

type YamlConf struct{}

/**
 * @description: GetYamlConf
 * @param {string} yamlConfPath
 * @param {interface{}} config
 * @author: Jerry.Yang
 * @date: 2022-09-23 17:54:56
 * @return {*}
 */
func GetYamlConf(yamlConfPath string, config interface{}) error {
	return CreateConfInterface(&conf.YamlConf{YamlFilePath: yamlConfPath, Conf: &config}).ConfInterface.GetConf()
}
