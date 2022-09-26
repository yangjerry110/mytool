/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 15:53:38
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 18:35:21
 * @Description: create
 */
package cmdPkg

import "github.com/yangjerry110/mytool/cmd"

type CreatePkgInterface interface {
	CreateInterface(cmdInterface cmd.CmdInterface) cmd.CmdInterface
	CreateAppContent(projectName string, appName string, method string) (string, error)
	CreateAppInputVoContent(projectName string, appName string, method string, dirName string, fileName string) (string, error)
	CreateAppOutputVoContent(projectName string, appName string, method string, dirName string, fileName string) (string, error)
	CreateAppRouteContent(projectName string, appName string, method string, dirName string, fileName string) (string, error)
	CreateAppServiceContent(projectName string, appName string, method string, dirName string, fileName string) (string, error)
	CreateDao(projectName string, daoName string, modelName string, authorName ...string) (string, error)
}

type CreatePkg struct {
	cmdInterface cmd.CmdInterface
}

/**
 * @description: CreateInteface
 * @param {cmd.CmdInterface} cmdInterface
 * @author: Jerry.Yang
 * @date: 2022-09-23 18:58:47
 * @return {*}
 */
func CreateCmdInterface(cmdInterface cmd.CmdInterface) *CreatePkg {
	return &CreatePkg{cmdInterface: cmdInterface}
}
