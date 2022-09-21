/*
 * @Author: yangjie04@qutoutiao.net
 * @Date: 2022-09-20 11:25:03
 * @LastEditors: yangjie04@qutoutiao.net
 * @LastEditTime: 2022-09-21 10:40:39
 * @Description: createPerm
 */
package pkg

import "github.com/yangjerry110/mytool/perm"

type CreatePermPkgInterface interface {
	CreatePerm(byteSize int32, permPath string) (bool, error)
}

type CreatePermPkg struct{}

/**
 * @description: CreatePerm
 * @param {int32} byteSize
 * @param {string} permPath
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-20 11:28:54
 * @return {*}
 */
func CreatePerm(byteSize int32, permPath string) (bool, error) {
	createPermObj := perm.CreateRasPerm{}
	result, err := createPermObj.CreatePerm(byteSize, permPath)
	return result, err
}
