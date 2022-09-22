/*
 * @Author: yangjie04@qutoutiao.net
 * @Date: 2022-09-22 14:20:20
 * @LastEditors: yangjie04@qutoutiao.net
 * @LastEditTime: 2022-09-22 14:27:43
 * @Description: pkg
 */
package pkg

import "github.com/yangjerry110/mytool/perm"

type PermPkgInterface interface {
	RasCreatePerm(byteSize int32, permPath string) (bool, error)
	RasDecrty(permPath string, inputStr string) (string, error)
	RasEncrty(permPath string, inputStr string) (string, error)
}

type PermPkg struct{}

/**
 * @description: RasCreatePerm
 * @param {int32} byteSize
 * @param {string} permPath
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-22 14:23:06
 * @return {*}
 */
func RasCreatePerm(byteSize int32, permPath string) (bool, error) {
	createPermObj := perm.CreateRasPerm{}
	return createPermObj.CreatePerm(byteSize, permPath)
}

/**
 * @description: RasDecrty
 * @param {string} permPath
 * @param {string} inputStr
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-22 14:23:13
 * @return {*}
 */
func RasDecrty(permPath string, inputStr string) (string, error) {
	rasDecrtyObj := perm.RasDecrty{}
	return rasDecrtyObj.Decrty(permPath, inputStr)
}

/**
 * @description: RasEncrty
 * @param {string} permPath
 * @param {string} inputStr
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-22 14:23:21
 * @return {*}
 */
func RasEncrty(permPath string, inputStr string) (string, error) {
	rasEncrtyObj := perm.RasEncrty{}
	return rasEncrtyObj.Encrty(permPath, inputStr)
}
