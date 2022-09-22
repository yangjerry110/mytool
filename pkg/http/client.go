/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 16:05:20
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 16:09:06
 * @Description: http
 */
package http

import (
	"io"

	"github.com/yangjerry110/mytool/http"
)

type HttpClientInterface interface {
	HttpRequest(method string, url string, body io.Reader, output interface{}, options ...http.HttpOptionFunc) error
}

type HttpClient struct{}

/**
 * @description: HttpRequest
 * @param {string} method
 * @param {string} url
 * @param {io.Reader} body
 * @param {interface{}} output
 * @param {...http.HttpOptionFunc} option
 * @author: Jerry.Yang
 * @date: 2022-09-22 16:08:21
 * @return {*}
 */
func HttpRequest(method string, url string, body io.Reader, output interface{}, options ...http.HttpOptionFunc) error {
	httpClient := http.HttpClient{
		Method:  method,
		Url:     url,
		Body:    body,
		Options: options,
		Output:  &output,
	}
	return httpClient.HttpRequest()
}
