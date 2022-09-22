/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 15:53:38
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 15:57:29
 * @Description: create
 */
package cmd

import "github.com/yangjerry110/mytool/cmd"

type CreateInterface interface{}

type Create struct{}

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
	createAppObj := cmd.CreateApp{}
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
	createDaoObj := cmd.CreateDao{}
	createDaoObj.CreatedDao(projectName, daoName, modelName, authorName...)
}
