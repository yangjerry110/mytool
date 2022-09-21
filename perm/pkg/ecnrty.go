/*
 * @Author: yangjie04@qutoutiao.net
 * @Date: 2022-09-20 11:25:03
 * @LastEditors: yangjie04@qutoutiao.net
 * @LastEditTime: 2022-09-20 19:37:16
 * @Description: ecnrty 加密
 */
package pkg

import "github.com/yangjerry110/mytool/perm"

type EncrtyPkgInterface interface {
	RasEncrty(permPath string, inputStr string) (string, error)
}

type EncrtyPkg struct{}

/**
 * @description: RasEncrty
 * @param {string} permPath
 * @param {string} inputStr
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-20 19:28:02
 * @return {*}
 */
func RasEncrty(permPath string, inputStr string) (string, error) {
	rasEncrtyObj := perm.RasEncrty{}
	outputStr, err := rasEncrtyObj.Encrty(permPath, inputStr)
	return outputStr, err
}
