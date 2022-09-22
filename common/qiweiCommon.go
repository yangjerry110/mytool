/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-21 17:42:36
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 16:50:30
 * @Description: qiwei common
 */
package common

import (
	"errors"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/yangjerry110/mytool/http"
)

type (
	QiweiCommonInterface interface{}

	QiweiCommon struct {
		AppId       string
		CropId      string
		CropSecret  string
		AccessToken string
	}
)

/**
 * @description: GetQiweiAccessToken
 * @author: Jerry.Yang
 * @date: 2022-09-22 16:02:25
 * @return {*}
 */
func (q *QiweiCommon) GetQiweiAccessToken() error {
	/**
	 * @step
	 * @初始化go-cache
	 **/
	goCache := cache.New(5*time.Minute, 10*time.Minute)

	/**
	 * @step
	 * @定义key
	 **/
	key := fmt.Sprintf("%s_qiwei_access_token", q.AppId)

	/**
	 * @step
	 * @获取缓存里面的accessToken
	 **/
	cacheAccessToken, cacheResult := goCache.Get(key)
	if cacheResult && cacheAccessToken.(string) == "" {
		q.AccessToken = cacheAccessToken.(string)
		return nil
	}

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
	httpClient := http.HttpClient{
		Method: "GET",
		Url:    url,
		Output: resp,
	}
	httpClient.HttpRequest()

	/**
	 * @step
	 * @判断accessToken和错误
	 **/
	if resp.Errcode != 0 {
		return errors.New(resp.Errmsg)
	}

	/**
	 * @step
	 * @设置accessToken缓存
	 **/
	goCache.Set(key, resp.AccessToken, time.Duration(resp.ExpiresIn)*time.Second)
	q.AccessToken = resp.AccessToken
	return nil
}
