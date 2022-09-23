/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 15:53:38
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-23 15:20:01
 * @Description: create
 */
package cmd

import "github.com/yangjerry110/mytool/cmd"

type CreateInterPkgface interface {
	CreateApp(projectName string, appName string, method string) error
	CreateDao(projectName string, daoName string, modelName string, authorName ...string) error
}

type CreatePkg struct {
	CreateDaoInterface cmd.CreateDaoInterface
	CreateAppInterface cmd.CreateAppInterface
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
func CreateApp(projectName string, appName string, method string) error {
	return CreateCmdCreateAppInstance().CreateApp(projectName, appName, method)
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
func CreateDao(projectName string, daoName string, modelName string, authorName ...string) error {
	return CreateCmdCreateDaoInstance().CreatedDao(projectName, daoName, modelName, authorName...)
}
