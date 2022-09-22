/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-21 17:23:18
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 11:01:22
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
	qiweiNotice := notice.QiweiNotice{
		AppId:      "kuhe",
		CropId:     "ww2cb3d75879c33965",
		CropSecret: "gnbzD0ZypFfLVdJG6Rwd1Rpw03xJlcZ3zTgbqPWRbyY",
		UserIds:    "Jerry.Yang",
		MsgType:    "markdown",
		AgentId:    "1000077",
		SendMsg:    "3333",
	}
	result, err := pkg.QiweiNotifyMessage(&qiweiNotice)
	fmt.Printf("err : %v", err)
	fmt.Print("\r\n")

	fmt.Printf("result : %v", result)
	fmt.Print("\r\n")
}
