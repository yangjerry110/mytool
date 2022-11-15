/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-23 14:44:20
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-15 14:31:29
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
	GetConfResult(filepath string, fileName string, fileType string, intervals time.Duration, conf interface{}) interface{}
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
 * @description: GetConfResult
 * @param {string} filepath
 * @param {string} fileName
 * @param {string} fileType
 * @param {time.Duration} intervals
 * @param {interface{}} conf
 * @author: Jerry.Yang
 * @date: 2022-11-15 14:31:28
 * @return {*}
 */
func GetConfResult(filepath string, fileName string, fileType string, intervals time.Duration, conf interface{}) interface{} {
	mytoolInitConf := &mytoolConf.InitConf{
		FilePath:  filepath,
		FileName:  fileName,
		FileType:  fileType,
		Intervals: intervals,
	}
	return CreateConfInterface(mytoolInitConf).ConfInterface.Init(conf).GetParseConf()
}
