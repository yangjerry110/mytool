/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-23 15:57:51
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-23 16:15:28
 * @Description: perm
 */
package perm

import "github.com/yangjerry110/mytool/perm"

type PermPkgInterface interface {
	RsaCreatePerm(byteSize int32, permPath string) (bool, error)
	RsaDecrty(permPath string, inputStr string) (string, error)
	RsaEncrty(permPath string, inputStr string) (string, error)
}

type PermPkg struct {
	CreatePermInterface perm.CreatePermInterface
	DecrtyPermInterface perm.DecrtyInterface
	EncrtyPermInterface perm.EncrtyInterface
}

/**
 * @description: RsaCreatePerm
 * @param {int32} byteSize
 * @param {string} permPath
 * @author: Jerry.Yang
 * @date: 2022-09-23 16:15:35
 * @return {*}
 */
func RsaCreatePerm(byteSize int32, permPath string) (bool, error) {
	return CreateRasCreatePermInstance().CreatePerm(byteSize, permPath)
}

/**
 * @description: RsaDecrty
 * @param {string} permPath
 * @param {string} inputStr
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:23:13
 * @return {*}
 */
func RsaDecrty(permPath string, inputStr string) (string, error) {
	return CreateRasDecrtyPermInstance().Decrty(permPath, inputStr)
}

/**
 * @description: RsaEncrty
 * @param {string} permPath
 * @param {string} inputStr
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:23:21
 * @return {*}
 */
func RsaEncrty(permPath string, inputStr string) (string, error) {
	return CreateRasEncrtyInstance().Encrty(permPath, inputStr)
}
