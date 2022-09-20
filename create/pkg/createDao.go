/*
 * @Author: yangjie04@qutoutiao.net
 * @Date: 2022-09-20 10:53:41
 * @LastEditors: yangjie04@qutoutiao.net
 * @LastEditTime: 2022-09-20 10:53:46
 * @Description: createDao
 */
package pkg

import "github.com/yangjerry110/mytool/create"

type CreateDaoInterface interface {
	CreateDao(projectName string, daoName string, modelName string, authorName ...string)
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
