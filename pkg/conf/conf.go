/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-23 14:44:20
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-15 18:59:38
 * @Description: conf
 */
package conf

import (
	"time"

	mytoolConf "github.com/yangjerry110/mytool/conf"
)

type ConfInterface interface {
	CreateConfInterface(confInterface mytoolConf.ConfInterface) mytoolConf.ConfInterface
	GetYamlConf(yamlConfPath string, config interface{}) error
	GetConf(filepath string, fileName string, fileType string, intervals time.Duration, conf interface{}) error
	Init(filepath string, fileName string, fileType string, intervals time.Duration, conf interface{}) mytoolConf.ConfInterface
	GetParseConf(filepath string, fileName string) interface{}
}

type Conf struct {
	ConfInterface mytoolConf.ConfInterface
}

/**
 * @description: CreateConfInterface
 * @param {conf.ConfInterface} confInterface
 * @author: Jerry.Yang
 * @date: 2022-09-26 18:39:40
 * @return {*}
 */
func CreateConfInterface(confInterface mytoolConf.ConfInterface) *Conf {
	return &Conf{ConfInterface: confInterface}
}

/**
 * @description: GetConf
 * @param {string} filepath
 * @param {string} fileName
 * @param {string} fileType
 * @param {time.Duration} intervals
 * @param {interface{}} conf
 * @author: Jerry.Yang
 * @date: 2022-11-15 14:20:33
 * @return {*}
 */
func GetConf(filepath string, fileName string, fileType string, intervals time.Duration, conf interface{}) error {
	mytoolConf := &mytoolConf.Conf{
		FilePath:  filepath,
		FileName:  fileName,
		FileType:  fileType,
		Intervals: intervals,
		Data:      conf,
	}
	return CreateConfInterface(mytoolConf).ConfInterface.GetNewConf()
}

/**
 * @description: Init
 * @param {string} filepath
 * @param {string} fileName
 * @param {string} fileType
 * @param {time.Duration} intervals
 * @param {interface{}} conf
 * @author: Jerry.Yang
 * @date: 2022-11-15 14:45:24
 * @return {*}
 */
func Init(filepath string, fileName string, fileType string, intervals time.Duration, conf interface{}) mytoolConf.ConfInterface {
	mytoolInitConf := &mytoolConf.InitConf{
		FilePath:  filepath,
		FileName:  fileName,
		FileType:  fileType,
		Intervals: intervals,
	}
	return CreateConfInterface(mytoolInitConf).ConfInterface.Init(conf)
}

/**
 * @description: GetParseConf
 * @param {string} filepath
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-11-15 19:00:04
 * @return {*}
 */
func GetParseConf(filepath string, fileName string) interface{} {
	mytoolInitConf := &mytoolConf.InitConf{
		FilePath: filepath,
		FileName: fileName,
	}
	return CreateConfInterface(mytoolInitConf).ConfInterface.GetParseConf()
}
