/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-21 17:23:18
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 14:55:12
 * @Description: notice test
 */
package test

import (
	"fmt"
	"testing"

	"github.com/yangjerry110/mytool/notice"
	"github.com/yangjerry110/mytool/notice/pkg"
)

func TestQiweiNotice(t *testing.T) {
	qiweiNotice := notice.QiweiNotice{}
	result, err := pkg.QiweiNotifyMessage(&qiweiNotice)
	fmt.Printf("err : %v", err)
	fmt.Print("\r\n")

	fmt.Printf("result : %v", result)
	fmt.Print("\r\n")
}
