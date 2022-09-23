/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-23 11:00:49
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-23 14:43:29
 * @Description: common
 */
package common

import "github.com/yangjerry110/mytool/common"

type CommonPkgInterface interface {
	GetQiweiAccessToken(appId string, cropId string, cropSecret string) (string, error)
}

type CommonPkg struct {
	QiweiCommonInterface common.QiweiCommonInterface
}

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
	return CreateCommonQiweiCommonInstance(appId, cropId, cropSecret).GetQiweiAccessToken()
}
