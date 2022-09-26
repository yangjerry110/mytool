/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 16:02:35
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 18:39:03
 * @Description: common
 */
package common

import "github.com/yangjerry110/mytool/common"

type QiweiCommon struct{}

/**
 * @description: GetQiweiAccessToken
 * @param {string} appId
 * @param {string} cropId
 * @param {string} cropSecret
 * @author: Jerry.Yang
 * @date: 2022-09-23 14:29:43
 * @return {*}
 */
func GetQiweiAccessToken(appId string, cropId string, cropSecret string) (string, error) {
	return CreateCommonInterface(&common.QiweiCommon{AppId: appId, CropId: cropId, CropSecret: cropSecret}).CommonInterface.GetAccessToken()
}
