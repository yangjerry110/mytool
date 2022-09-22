/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-21 17:45:07
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 14:28:33
 * @Description: qiwei common pkg
 */
package pkg

import "github.com/yangjerry110/mytool/common"

type CommonPkgInterface interface {
	GetQiweiAccessToken(qiweiCommon *common.QiweiCommon) error
}

type CommonPkg struct{}

/**
 * @description: GetQiweiAccessToken
 * @param {*common.QiweiCommon} qiweiCommon
 * @author: Jerry.Yang
 * @date: 2022-09-21 17:47:48
 * @return {*}
 */
func GetQiweiAccessToken(qiweiCommon *common.QiweiCommon) error {
	return qiweiCommon.GetQiweiAccessToken()
}
