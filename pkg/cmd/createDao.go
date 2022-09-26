/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-26 15:03:45
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 18:37:31
 * @Description: createDao
 */
package cmdPkg

import "github.com/yangjerry110/mytool/cmd"

type CreateDaoPkg struct{}

/**
 * @description: CreateDao
 * @param {string} projectName
 * @param {string} daoName
 * @param {string} modelName
 * @param {...string} authorName
 * @author: Jerry.Yang
 * @date: 2022-09-26 15:08:41
 * @return {*}
 */
func CreateDao(projectName string, daoName string, modelName string, authorName ...string) (string, error) {
	return CreateCmdInterface(&cmd.CreateDao{ProjectName: projectName, DaoName: daoName, ModelName: modelName, AuthorName: authorName}).cmdInterface.CreateContent()
}
