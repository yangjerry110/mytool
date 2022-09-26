/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-26 15:15:25
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 16:32:35
 * @Description: conf
 */
package conf

type ConfInterface interface {
	GetConf() error
}

type YamlConf struct {
	YamlFilePath string
	Conf         interface{}
}
