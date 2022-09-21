/*
 * @Author: yangjie04@qutoutiao.net
 * @Date: 2022-09-21 15:30:12
 * @LastEditors: yangjie04@qutoutiao.net
 * @LastEditTime: 2022-09-21 17:24:40
 * @Description: 企微通知
 */
package notice

import (
	"fmt"

	"github.com/yangjerry110/mytool/http"
	"github.com/yangjerry110/mytool/http/pkg"
)

type QiweiNoticeInterface interface{}

type QiweiNotice struct {
	appId          string
	CropId         string
	CropSecret     string
	AgentId        string
	DepartmentIds  string
	TagIds         string
	UserIds        string
	Safe           int32
	SendMsg        string
	MediaData      string
	MediaType      string
	Title          string
	Description    string
	Url            string
	PicUrl         string
	EnableIdTrans  int32
	Btntxt         string
	AppletId       string
	AppletPagepath string
}

func (q *QiweiNotice) NotifyMessage() (bool, error) {
	/**
	 * @step
	 * @初始化一个通道,获取企微请求参数
	 **/
	//notifyQiweiReqChan := make(chan []byte)
	return false, nil
}

/**
 * @description: getAccessToken
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-21 15:42:47
 * @return {*}
 */
func (q *QiweiNotice) GetAccessToken() (string, error) {

	/**
	 * @step
	 * @初始化go-cache
	 **/
	//goCache := cache.New(5*time.Minute, 10*time.Minute)

	/**
	 * @step
	 * @获取accessToken的url
	 **/
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", q.CropId, q.CropSecret)

	/**
	 * @step
	 * @定义出参
	 **/
	type GetAccessTokenOutput struct {
		Errcode     int32  `json:"errcode"`
		Errmsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	/**
	 * @step
	 * @获取accessToken
	 **/
	resp := &GetAccessTokenOutput{}
	HttpClient := http.HttpClient{
		Method: "GET",
		Url:    url,
		Output: resp,
	}
	pkg.HttpRequest(&HttpClient)

	fmt.Printf("httpClient : %v", HttpClient)
	fmt.Print("\r\n")

	return "", nil
}
