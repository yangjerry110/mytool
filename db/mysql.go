/*
 * @Author: Jerry.Yang
 * @Date: 2022-11-07 15:01:27
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-15 10:26:29
 * @Description:  mysql
 */
package db

import (
	"errors"
	"fmt"

	"github.com/yangjerry110/mytool/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbMysqlInteraface interface{}

type DbMysql struct {
	Dsn string `yaml:"dsn"`
}

/**
 * @description: Client
 * @param {string} yamlFilePath
 * @param {string} yamlFileName
 * @param {string} dbName
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:25:40
 * @return {*}
 */
func (d *DbMysql) Client(yamlFilePath string, yamlFileName string, dbName string) (*gorm.DB, error) {

	/**
	 * @step
	 * @定义mysqlconf
	**/
	mysqlConfs := make(map[string]DbMysql, 0)

	/**
	 * @step
	 * @获取配置
	 **/
	yamlConf := conf.YamlConf{FilePath: yamlFilePath, FileName: yamlFileName, Conf: &mysqlConfs}
	err := yamlConf.GetConf(mysqlConfs)
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @判断是否存在
	 **/
	mysqlConf, mysqlConfIsExist := mysqlConfs[dbName]
	if !mysqlConfIsExist {
		return nil, errors.New(fmt.Sprintf("mysql client Err : %s mysqlConf is not set", dbName))
	}

	/**
	 * @step
	 * @判断dsn
	 **/
	if mysqlConf.Dsn == "" {
		return nil, errors.New(fmt.Sprintf("mysql client Err : %s dsn is not set", dbName))
	}

	/**
	 * @step
	 * @创建db链接
	 **/
	return gorm.Open(mysql.Open(mysqlConf.Dsn), &gorm.Config{})
}
