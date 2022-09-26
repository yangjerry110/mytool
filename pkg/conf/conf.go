/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-23 14:44:20
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 18:39:31
 * @Description: conf
 */
package conf

import "github.com/yangjerry110/mytool/conf"

type ConfInterface interface {
	CreateConfInterface(confInterface conf.ConfInterface) conf.ConfInterface
	GetYamlConf(yamlConfPath string, config interface{}) error
}

type Conf struct {
	ConfInterface conf.ConfInterface
}

/**
 * @description: CreateConfInterface
 * @param {conf.ConfInterface} confInterface
 * @author: Jerry.Yang
 * @date: 2022-09-26 18:39:40
 * @return {*}
 */
func CreateConfInterface(confInterface conf.ConfInterface) *Conf {
	return &Conf{ConfInterface: confInterface}
}
