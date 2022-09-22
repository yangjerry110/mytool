/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-21 17:18:14
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 14:27:15
 * @Description: http client pkg
 */
package pkg

import (
	mytoolHttp "github.com/yangjerry110/mytool/http"
)

type HttpPkgInterface interface {
	HttpRequest(httpClient *mytoolHttp.HttpClient) error
}

type HttpPkg struct{}

/**
 * @description: HttpRequest
 * @param {*http.HttpClient} httpClient
 * @author: Jerry.Yang
 * @date: 2022-09-21 17:20:41
 * @return {*}
 */
func HttpRequest(httpClient *mytoolHttp.HttpClient) error {
	return httpClient.HttpRequest()
}
