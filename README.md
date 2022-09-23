<!--
 * @Author: Jerry.Yang
 * @Date: 2022-09-19 17:46:05
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-23 17:44:33
 * @Description: 
-->
# my-tool
## 1.cmd
### 脚本文件
> 1.需要单独启用一个项目对cmd中的方法进行引用，再进行go build生成可执行的二进制文件    
> 2.主要用作go generate自动化相关

实例参考 ： https://github.com/yangjerry110/createtool/tree/master 

## 2.common
> 本意是此项目的一些公用方法，考虑到也可以public就也放到pkg中了


```
package main

import "github.com/yangjerry110/mytool/pkg/common"

func main() {
    accessToken,err := common.GetQiweiAccessToken(appId,cropId,cropSecret)
}
```

## 3.conf
>1.对一些配置文件进行解析，需要提前定义到需要解析的文件和解析对应的结构体  
>2.配置文件定义(以yaml文件示例)

```
conf_path: "conf_path"
```

引用实例：

```
package main

import "github.com/yangjerry110/mytool/pkg/conf"

func main() {

    type MyConf struct {
        ConfPath string `yaml:"conf_path"`
    }

    yamlConf := conf.YamlConfPkg{
        YamlFilePath : "",
        Config: &MyConf,
    }

    // 解析yaml的配置到MyConf结构体
    yamlConf.ParseYamlConf()
    myConfPath := MyConf.ConfPath

}
```

## 4.http
> 1.http调用相关的  
> 2.需要提前定义好input, output   
> 3.ps : options 配置需要引导 http.Option下   

```
package main 

import (
    httpPkg "github.com/yangjerry110/mytool/pkg/http" 
    "github.com/yangjerry110/mytool/http"
)

func main() {

    // 以json数据示例
    type Input struct {
        InputStr string
    }

    type OutPut struct {
        RetCode int 
        RetMsg string
        RetResult bool
    }

    jsonReq,_ := json.Marshal(Input{InputStr:"inputStr"})

    // 定义options
    httpOptions := http.HttpOptions{}
    options := []http.HttpOptionFunc{
        httpOptions.SetHeaders(map[string]string{
            "Content-Type": "multipart/form-data",
        })
    }

    resp := &OutPut{}
    httpPkg.HttpRequest(method,url,bytes.NewBuffer(jsonReq),resp,options)
}

```

## 5.notice

> 1.通知相关   
> 2.以企微通知示例   
> 3.支持 text,markdown,image,news,card    
> 4.bot通知 支持 text,markdown,image,news

```
package main 

import "github.com/yangjerry110/mytool/pkg/notice"

func main() {

    /**
    * @param {string} AppId appId
    * @param {string} MsgType 消息类型
    * @param {string} CropId 公司id
    * @param {string} CropSecret 公司秘钥
    * @param {string} AgentId 通知的qiwei应用的id
    * @param {string} DepartmentIds 通知的组织架构集合
    * @param {string} TagIds 通知的Tag集合
    * @param {string} UserIds 通知的人
    * @param {int32} Safe safe
    * @param {string} SendMsg 通知的消息
    * @param {string} MediaData 通知的媒体消息内容
    * @param {string} MediaType 通知的媒体消息类型
    * @param {string} Title 标题
    * @param {string} Description 简介
    * @param {string} Url 链接
    * @param {string} PicUrl 图片链接
    * @param {int32} EnableIdTrans
    * @param {string} Btntxt news消息时候的按钮
    * @param {string} AppletId AppletId 小程序id
    * @param {string} AppletPagepath AppletPagepath 小程序链接
    * @param {string} QiweiFilePath 通知媒体消息的时候，存放媒体内容的地址
     */
    qiweiNotice := notice.QiweiNoticePkg{}

    result,err := qiweiNotice.QiweiNotice()

}
```

## 6.perm 

> 1.加密解密相关    
> 2.PS: 使用rsa加解密之前，需要先生成公钥和私钥

```

package main 

import "github.com/yangjerry110/mytool/pkg/perm"

func main() {
    encrtyStr,err := perm.RsaDecrty(permPath,inputStr)
}

```

## 7.upload

> 1.以ali的oss示例    
> 2.PS： 上传的fileData需要是以base64编码的文件内容

```

package main 

import "github.com/yangjerry110/mytool/pkg/upload"

func main() {

    aliUploadOss := upload.AliUploadOss{}

    ossUrl,err := aliUploadOss.AliUploadOss()

}

```






