/*
 * @Author: yangjie04@qutoutiao.net
 * @Date: 2022-09-20 10:51:10
 * @LastEditors: yangjie04@qutoutiao.net
 * @LastEditTime: 2022-09-20 11:25:48
 * @Description: pkg
 */
package pkg

import (
	"github.com/yangjerry110/mytool/create"
)

type CreateAppPkgInterface interface {
	CreateApp(projectName string, appName string, method string)
}

type CreateAppPkg struct{}

/**
 * @description: CreateApp
 * @param {string} projectName
 * @param {string} appName
 * @param {string} method
 * @author: Jerry.Yang
 * @date: 2022-09-19 17:43:23
 * @return {*}
 */
func CreateApp(projectName string, appName string, method string) {
	createAppObj := create.CreateApp{}
	createAppObj.CreateApp(projectName, appName, method)
}
