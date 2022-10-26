/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-21 17:23:18
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-10-26 18:03:33
 * @Description: notice test
 */
package test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/yangjerry110/mytool/external/ali"
)

func TestQiweiNotice(t *testing.T) {

	aliUploadLocaFile := ali.AliOssUploadFromLocalFile{
		LocalFilePath: "/Users/admin/Desktop/1.png",
	}

	fileContent, _ := aliUploadLocaFile.GetFileData()
	fileContentBase64 := base64.StdEncoding.EncodeToString(fileContent)

	a := ali.AliDingdingNotice{
		RedisConfPath:    "/Users/admin/go/src/my-tool/test/conf/redis.yaml",
		DingdingFilePath: "",
		AppId:            "test",
		AppKey:           "dinggxetcehjiqxpgedx",
		AgentId:          "1946178798",
		AppSecret:        "UjvQWwvO0nxQzYqECsx6LkqKN9QvOI-tJOBuvsA755BTt40Tchro8vjGLJfzSFvO",
		MsgType:          "news",
		Title:            "test markdown ",
		FileType:         "image",
		MediaType:        "png",
		MediaData:        fileContentBase64,
		MessageUrl:       "https://www.baidu.com",
		SingleTitle:      "查看详情",
		SingleUrl:        "https://www.baidu.com",
		Msg:              "# 这是支持markdown的文本   \n   ## 标题2    \n   * 列表1   \n  ![alt 啊](https://img.alicdn.com/tps/TB1XLjqNVXXXXc4XVXXXXXXXXXX-170-64.png)",
		UserIds:          "manager4791",
	}
	accessToken, err := a.NotifyMessage()

	fmt.Printf("accessToken : %v", accessToken)
	fmt.Print("\r\n")

	fmt.Printf("err : %v", err)
	fmt.Print("\r\n")
}
