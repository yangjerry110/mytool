/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 16:05:20
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-23 15:02:22
 * @Description: http
 */
package http

import (
	"io"

	"github.com/yangjerry110/mytool/http"
)

type HttpClientPkgInterface interface {
	CreateHttpClientInterface(httpClientInterface http.HttpClientInterface) *HttpPkg
	CreateHttpClientInstance() http.HttpClientInterface
}

type HttpClientPkg struct {
	Method  string
	Url     string
	Body    io.Reader
	Output  interface{}
	Options []http.HttpOptionFunc
}

/**
 * @description: CreateHttpClientInterface
 * @param {http.HttpClientInterface} httpClientInterface
 * @author: Jerry.Yang
 * @date: 2022-09-23 15:02:55
 * @return {*}
 */
func CreateHttpClientInterface(httpClientInterface http.HttpClientInterface) *HttpPkg {
	return &HttpPkg{HttpClientInterface: httpClientInterface}
}

/**
 * @description: CreateHttpClientInstance
 * @author: Jerry.Yang
 * @date: 2022-09-23 15:03:03
 * @return {*}
 */
func (h *HttpClientPkg) CreateHttpClientInstance() http.HttpClientInterface {
	return CreateHttpClientInterface(&http.HttpClient{Method: h.Method, Url: h.Url, Body: h.Body, Output: &h.Output, Options: h.Options}).HttpClientInterface
}
