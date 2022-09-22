/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 16:02:35
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 16:09:48
 * @Description: common
 */
package common

import "github.com/yangjerry110/mytool/common"

type QiweiCommonInterface interface {
	GetQiweiAccessToken(appId string, cropId string, cropSecret string) (string, error)
}

type QiweiCommon struct{}

/**
 * @description: GetQiweiAccessToken
 * @param {string} appId
 * @param {string} cropId
 * @param {string} cropSecret
 * @author: Jerry.Yang
 * @date: 2022-09-22 16:04:57
 * @return {*}
 */
func GetQiweiAccessToken(appId string, cropId string, cropSecret string) (string, error) {
	qiweiCommonObj := common.QiweiCommon{
		AppId:      appId,
		CropId:     cropId,
		CropSecret: cropSecret,
	}
	err := qiweiCommonObj.GetQiweiAccessToken()
	return qiweiCommonObj.AccessToken, err
}
