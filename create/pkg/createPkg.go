/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 14:25:44
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 14:26:59
 * @Description: pkg
 */
package pkg

import "github.com/yangjerry110/mytool/create"

type CreatePkgInterface interface {
	CreateApp(projectName string, appName string, method string)
	CreateDao(projectName string, daoName string, modelName string, authorName ...string)
}

type CreatePkg struct{}

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
