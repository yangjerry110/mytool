/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 16:02:35
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-23 14:42:52
 * @Description: common
 */
package common

import "github.com/yangjerry110/mytool/common"

type QiweiCommonPkgInterface interface {
	CreateQiweiInterface(qiweiCommonInterface common.QiweiCommonInterface) *CommonPkg
	CreateCommonQiweiInstance(appId string, cropId string, cropSecret string) common.QiweiCommonInterface
}

/**
 * @description: CreateQiweiInterface
 * @param {common.QiweiCommonInterface} qiweiCommonInterface
 * @author: Jerry.Yang
 * @date: 2022-09-23 14:37:30
 * @return {*}
 */
func CreateQiweiCommonInterface(qiweiCommonInterface common.QiweiCommonInterface) *CommonPkg {
	return &CommonPkg{QiweiCommonInterface: qiweiCommonInterface}
}

/**
 * @description: CreateQiweiInstance
 * @param {string} appId
 * @param {string} cropId
 * @param {string} cropSecret
 * @author: Jerry.Yang
 * @date: 2022-09-23 14:38:13
 * @return {*}
 */
func CreateCommonQiweiCommonInstance(appId string, cropId string, cropSecret string) common.QiweiCommonInterface {
	return CreateQiweiCommonInterface(&common.QiweiCommon{AppId: appId, CropId: cropId, CropSecret: cropSecret}).QiweiCommonInterface
}
