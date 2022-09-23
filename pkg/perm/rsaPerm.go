/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 16:20:52
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-23 16:16:16
 * @Description: ras
 */
package perm

import "github.com/yangjerry110/mytool/perm"

type RasPermInterface interface {
	CreateCreatePermInterface(createPermInterface perm.CreatePermInterface) *PermPkg
	CreateRasCreatePermInstance() perm.CreatePermInterface
	CreateDecrtyPermInterface(decrtyPermInterface perm.DecrtyInterface) *PermPkg
	CreateRasDecrtyPermInstance() perm.DecrtyInterface
	CreateEncrtyPermInterface(encrtyPermInterface perm.EncrtyInterface) *PermPkg
	CreateRasEncrtyInstance() perm.EncrtyInterface
}

type RasPerm struct{}

/**
 * @description: CreateCreatePermInterface
 * @param {perm.CreatePermInterface} createPermInterface
 * @author: Jerry.Yang
 * @date: 2022-09-23 16:16:01
 * @return {*}
 */
func CreateCreatePermInterface(createPermInterface perm.CreatePermInterface) *PermPkg {
	return &PermPkg{CreatePermInterface: createPermInterface}
}

/**
 * @description: CreateRasCreatePermInstance
 * @author: Jerry.Yang
 * @date: 2022-09-23 16:16:08
 * @return {*}
 */
func CreateRasCreatePermInstance() perm.CreatePermInterface {
	return CreateCreatePermInterface(&perm.CreateRsaPerm{}).CreatePermInterface
}

/**
 * @description: CreateDecrtyPermInterface
 * @param {perm.DecrtyInterface} decrtyPermInterface
 * @author: Jerry.Yang
 * @date: 2022-09-23 16:16:15
 * @return {*}
 */
func CreateDecrtyPermInterface(decrtyPermInterface perm.DecrtyInterface) *PermPkg {
	return &PermPkg{DecrtyPermInterface: decrtyPermInterface}
}

/**
 * @description: CreateRasDecrtyPermInstance
 * @author: Jerry.Yang
 * @date: 2022-09-23 16:16:25
 * @return {*}
 */
func CreateRasDecrtyPermInstance() perm.DecrtyInterface {
	return CreateDecrtyPermInterface(&perm.RsaDecrty{}).DecrtyPermInterface
}

/**
 * @description: CreateEncrtyPermInterface
 * @param {perm.EncrtyInterface} encrtyPermInterface
 * @author: Jerry.Yang
 * @date: 2022-09-23 16:16:32
 * @return {*}
 */
func CreateEncrtyPermInterface(encrtyPermInterface perm.EncrtyInterface) *PermPkg {
	return &PermPkg{EncrtyPermInterface: encrtyPermInterface}
}

/**
 * @description: CreateRasEncrtyInstance
 * @author: Jerry.Yang
 * @date: 2022-09-23 16:16:40
 * @return {*}
 */
func CreateRasEncrtyInstance() perm.EncrtyInterface {
	return CreateEncrtyPermInterface(&perm.RsaEncrty{}).EncrtyPermInterface
}
