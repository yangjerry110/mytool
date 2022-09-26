/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-26 15:09:17
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 16:29:50
 * @Description: common
 */
package common

type CommonInterface interface {
	GetAccessToken() (string, error)
}

type QiweiCommon struct {
	AppId      string
	CropId     string
	CropSecret string
}
