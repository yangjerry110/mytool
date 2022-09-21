/*
 * @Author: yangjie04@qutoutiao.net
 * @Date: 2022-09-21 17:18:14
 * @LastEditors: yangjie04@qutoutiao.net
 * @LastEditTime: 2022-09-21 17:26:18
 * @Description: http client pkg
 */
package pkg

import (
	"fmt"

	"github.com/yangjerry110/mytool/http"
)

type (
	HttpClientPkgInterface interface {
		HttpRequest(httpClient *http.HttpClient) error
	}

	HttpClientPkg struct{}
)

/**
 * @description: HttpRequest
 * @param {*http.HttpClient} httpClient
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-21 17:20:41
 * @return {*}
 */
func HttpRequest(httpClient *http.HttpClient) error {

	fmt.Printf("httpClient : %v", httpClient)
	fmt.Print("\r\n")

	err := httpClient.HttpRequest()
	if err != nil {
		return err
	}
	return nil
}
