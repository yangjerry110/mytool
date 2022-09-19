/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-19 14:34:03
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-19 18:17:20
 * @Description: createDao
 */
package pkg

import "github.com/yangjerry110/mytool/create"

type CreateInterface interface {
	CreateDao(projectName string, daoName string, modelName string, authorName ...string)
	CreateApp(projectName string, appName string, method string)
}

type Create struct{}

/**
 * @description: CreateDaoFunc
 * @param {string} projectName
 * @param {string} daoName
 * @param {string} modelName
 * @param {...string} authorName
 * @author: Jerry.Yang
 * @date: 2022-09-19 14:36:42
 * @return {*}
 */
func CreateDao(projectName string, daoName string, modelName string, authorName ...string) {
	createDaoObj := create.CreateDao{}
	createDaoObj.CreatedDao(projectName, daoName, modelName, authorName...)
}

/**
 * @description: CreateApp
 * @param {string} projectName
 * @param {string} appName
 * @param {string} method
 * @author: Jerry.Yang
 * @date: 2022-09-19 17:43:23
 * @return {*}
 */
func CreateApp(projectName string, appName string, method string) {
	createAppObj := create.CreateApp{}
	createAppObj.CreateApp(projectName, appName, method)
}
