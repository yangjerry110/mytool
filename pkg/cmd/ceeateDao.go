/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-23 10:49:32
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-23 14:41:42
 * @Description: create dao
 */
package cmd

import "github.com/yangjerry110/mytool/cmd"

type CreateDaoCmdInterface interface {
	CreateDaoInterface() *CreatePkg
	CreateCmdCreateDaoInstance() cmd.CreateDaoInterface
}

/**
 * @description: CreateDaoInterface
 * @param {cmd.CreateDaoInterface} createDao
 * @author: Jerry.Yang
 * @date: 2022-09-23 11:36:13
 * @return {*}
 */
func CreateDaoInterface(createDaoInterface cmd.CreateDaoInterface) *CreatePkg {
	return &CreatePkg{CreateDaoInterface: createDaoInterface}
}

/**
 * @description: CreateCmdCreateDaoInstance
 * @author: Jerry.Yang
 * @date: 2022-09-23 11:36:21
 * @return {*}
 */
func CreateCmdCreateDaoInstance() cmd.CreateDaoInterface {
	return CreateDaoInterface(&cmd.CreateDao{}).CreateDaoInterface
}
