/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-21 16:36:25
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-29 16:51:48
 * @Description: option test
 */
package test

import (
	"fmt"
	"testing"

	"github.com/yangjerry110/mytool/http"
)

func TestOption(t *testing.T) {

	headers := map[string]string{}
	headers["aaaaaa"] = "bbbbbbbb"

	httpOptionsObj := http.HttpOptions{}
	httpOptions := []http.HttpOptionFunc{
		httpOptionsObj.SetHeaders(headers),
	}
	options := httpOptionsObj.SetOptions(httpOptions)
	setHeaders := httpOptionsObj.GetHeaders(options)
	fmt.Printf("setHeaders : %v", setHeaders)

}
