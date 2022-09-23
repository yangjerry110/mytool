/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 18:59:34
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-23 14:42:17
 * @Description: createApp
 */
package cmd

import "github.com/yangjerry110/mytool/cmd"

type CreateAppCmdInterface interface {
	CreateAppInterface(createAppInstance cmd.CreateAppInterface) *CreatePkg
	CreateCmdCreateAppInstance() cmd.CreateAppInterface
}

/**
 * @description: CreateAppInterface
 * @param {cmd.CreateAppInterface} createAppInstance
 * @author: Jerry.Yang
 * @date: 2022-09-23 14:20:53
 * @return {*}
 */
func CreateAppInterface(createAppInstance cmd.CreateAppInterface) *CreatePkg {
	return &CreatePkg{CreateAppInterface: createAppInstance}
}

/**
 * @description: CreateCmdCreateAppInstance
 * @author: Jerry.Yang
 * @date: 2022-09-23 14:21:03
 * @return {*}
 */
func CreateCmdCreateAppInstance() cmd.CreateAppInterface {
	return CreateAppInterface(&cmd.CreateApp{}).CreateAppInterface
}
