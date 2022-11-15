/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 17:19:49
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-11 16:05:39
 * @Description: yaml conf
 */
package conf

import "github.com/yangjerry110/mytool/conf"

type YamlConf struct{}

/**
 * @description: GetYamlConf
 * @param {string} filePath
 * @param {string} fileName
 * @param {interface{}} config
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:19:01
 * @return {*}
 */
func GetYamlConf(filePath string, fileName string, config interface{}) error {
	return CreateConfInterface(&conf.YamlConf{FilePath: filePath, FileName: fileName, Conf: config}).ConfInterface.GetConf(config)
}
