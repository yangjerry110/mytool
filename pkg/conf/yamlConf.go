/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 17:19:49
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-23 14:56:11
 * @Description: yaml conf
 */
package conf

import "github.com/yangjerry110/mytool/conf"

type YamlConfPkgInterface interface {
	CreateYamlConfInterface(yamlConfInterface conf.YamlConfInterface) *ConfPkg
	CreateConfYamlConfInstance() conf.YamlConfInterface
}

type YamlConfPkg struct {
	YamlFilePath string
	Config       interface{}
}

/**
 * @description: CreateYamlConfInterface
 * @param {conf.YamlConfInterface} yamlConfInterface
 * @author: Jerry.Yang
 * @date: 2022-09-23 14:51:14
 * @return {*}
 */
func CreateYamlConfInterface(yamlConfInterface conf.YamlConfInterface) *ConfPkg {
	return &ConfPkg{YamlConfInterface: yamlConfInterface}
}

/**
 * @description: CreateConfYamlConfInstance
 * @author: Jerry.Yang
 * @date: 2022-09-23 14:51:22
 * @return {*}
 */
func (y *YamlConfPkg) CreateYamlConfInstance() conf.YamlConfInterface {
	return CreateYamlConfInterface(&conf.YamlConf{YamlFilePath: y.YamlFilePath, Conf: &y.Config}).YamlConfInterface
}
