/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-23 14:52:39
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-23 15:00:47
 * @Description: http
 */
package http

import (
	"io"

	"github.com/yangjerry110/mytool/http"
)

type HttpPkgInterface interface {
	HttpRequest(method string, url string, body io.Reader, output interface{}, options ...http.HttpOptionFunc)
}

type HttpPkg struct {
	HttpClientInterface http.HttpClientInterface
}

/**
 * @description: HttpRequest
 * @param {string} method
 * @param {string} url
 * @param {io.Reader} body
 * @param {interface{}} output
 * @param {...http.HttpOptionFunc} options
 * @author: Jerry.Yang
 * @date: 2022-09-23 15:01:01
 * @return {*}
 */
func HttpRequest(method string, url string, body io.Reader, output interface{}, options ...http.HttpOptionFunc) error {
	httpClientPkg := HttpClientPkg{Method: method, Url: url, Body: body, Output: &output, Options: options}
	return httpClientPkg.CreateHttpClientInstance().HttpRequest()
}
