/*
 * @Author: yangjie04@qutoutiao.net
 * @Date: 2022-09-20 19:34:59
 * @LastEditors: yangjie04@qutoutiao.net
 * @LastEditTime: 2022-09-20 19:36:36
 * @Description: decrty 解密
 */
package pkg

import "github.com/yangjerry110/mytool/perm"

type DecrtyPkgInterface interface {
	RasDecrty(permPath string, inputStr string) (string, error)
}

type DecrtyPkg struct{}

/**
 * @description: RasDecrty
 * @param {string} permPath
 * @param {string} inputStr
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-20 19:37:05
 * @return {*}
 */
func RasDecrty(permPath string, inputStr string) (string, error) {
	rasDecrtyObj := perm.RasDecrty{}
	outputStr, err := rasDecrtyObj.Decrty(permPath, inputStr)
	return outputStr, err
}
