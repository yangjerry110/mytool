/*
 * @Author: yangjie04@qutoutiao.net
 * @Date: 2022-09-21 17:23:18
 * @LastEditors: yangjie04@qutoutiao.net
 * @LastEditTime: 2022-09-21 17:24:44
 * @Description: notice test
 */
package test

import (
	"testing"

	"github.com/yangjerry110/mytool/notice"
)

func TestQiweiNotice(t *testing.T) {
	qiweiNotice := notice.QiweiNotice{}
	qiweiNotice.GetAccessToken()
}
