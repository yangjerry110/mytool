/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-21 17:23:18
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-29 10:41:27
 * @Description: notice test
 */
package test

import (
	"testing"
)

func TestQiweiNotice(t *testing.T) {
	// qiweiNotice := notice.QiweiNotice{
	// 	// AppId:      "kuhe",
	// 	// CropId:     "ww2cb3d75879c33965",
	// 	// CropSecret: "gnbzD0ZypFfLVdJG6Rwd1Rpw03xJlcZ3zTgbqPWRbyY",
	// 	// UserIds:    "Jerry.Yang",
	// 	// MsgType:    "markdown",
	// 	// AgentId:    "1000077",
	// 	// SendMsg:    "3333",
	// }

	/**
	 * @step
	 * @获取accessToken的url
	 **/
	// url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", "", "")

	// /**
	//  * @step
	//  * @定义出参
	//  **/
	// type GetAccessTokenOutput struct {
	// 	Errcode     int32  `json:"errcode"`
	// 	Errmsg      string `json:"errmsg"`
	// 	AccessToken string `json:"access_token"`
	// 	ExpiresIn   int    `json:"expires_in"`
	// }

	// /**
	//  * @step
	//  * @获取accessToken
	//  **/
	// resp := &GetAccessTokenOutput{}
	// http.HttpRequest("GET", url, bytes.NewBuffer([]byte{}), resp)

	// fmt.Printf("resp : %v", resp)
	// fmt.Print("\r\n")

}
