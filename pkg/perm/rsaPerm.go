/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 16:20:52
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 16:25:09
 * @Description: ras
 */
package perm

import "github.com/yangjerry110/mytool/perm"

type RasPermInterface interface {
	RsaCreatePerm(byteSize int32, permPath string) (bool, error)
	RsaDecrty(permPath string, inputStr string) (string, error)
	RsaEncrty(permPath string, inputStr string) (string, error)
}

type RasPerm struct{}

/**
 * @description: RsaCreatePerm
 * @param {int32} byteSize
 * @param {string} permPath
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:23:06
 * @return {*}
 */
func RsaCreatePerm(byteSize int32, permPath string) (bool, error) {
	createPermObj := perm.CreateRsaPerm{}
	return createPermObj.CreatePerm(byteSize, permPath)
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
	rasDecrtyObj := perm.RsaDecrty{}
	return rasDecrtyObj.Decrty(permPath, inputStr)
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
	rasEncrtyObj := perm.RsaEncrty{}
	return rasEncrtyObj.Encrty(permPath, inputStr)
}
