<!--
 * @Author: Jerry.Yang
 * @Date: 2022-09-19 17:46:05
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 18:11:34
 * @Description: 
-->
# my-tool
## 1.cmd
### 脚本文件
> 需要单独启用一个项目对cmd中的方法进行引用，再进行go build生成可执行的二进制文件

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
>对一些配置文件进行解析，需要提前定义到需要解析的文件和解析对应的结构体
1.配置文件定义(以yaml文件示例)

```
conf_path: "conf_path"
```

2.引用实例：

```
package main

import "github.com/yangjerry110/mytool/pkg/conf"

func main() {

    type MyConf struct {
        ConfPath string `yaml:"conf_path"`
    }

    // 解析yaml的配置到MyConf结构体
    conf.YamlParse(yamlFilePath,&MyConf)
    myConfPath := MyConf.ConfPath

}
```

## 4.http
> http调用相关的  
> 需要提前定义好input, output   
> ps : options 配置需要引导 http.Option下   

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

> 通知相关   
> 以企微通知示例   
> 支持 text,markdown,image,news,card    
> bot通知 支持 text,markdown,image,news

```
package main 

import "github.com/yangjerry110/mytool/pkg/notice"

func main() {

    /**
    * @param {string} AppId
    * @param {string} MsgType
    * @param {string} CropId
    * @param {string} CropSecret
    * @param {string} AgentId
    * @param {string} DepartmentIds
    * @param {string} TagIds
    * @param {string} UserIds
    * @param {int32} Safe
    * @param {string} SendMsg
    * @param {string} MediaData
    * @param {string} MediaType
    * @param {string} Title
    * @param {string} Description
    * @param {string} Url
    * @param {string} PicUrl
    * @param {int32} EnableIdTrans
    * @param {string} Btntxt
    * @param {string} AppletId
    * @param {string} AppletPagepath
    * @param {string} QiweiFilePath
     */
    result,err := notice.QiweiNotifyMessage(AppId,MsgType,CropId,CropSecret,AgentId,DepartmentIds,TagIds,userIds,safe,SendMsg,MediaData,MediaType,Title,Description,url,PicUrl,EnableIdTrans,Btntxt,AppletId,AppletPagepath,QiweiFilePath)

}
```

## 6.perm 

> 加密解密相关    
> PS: 使用rsa加解密之前，需要先生成公钥和私钥

```

package main 

import "github.com/yangjerry110/mytool/pkg/perm"

func main() {
    encrtyStr,err := perm.RsaEncrty(permPath,inputStr)
}

```

## 7.upload

> 以ali的oss示例    
> PS： 上传的fileData需要是以base64编码的文件内容

```

package main 

import "github.com/yangjerry110/mytool/pkg/upload"

func main() {

    aliUploadOssResp,err := upload.AliUploadOss(AccessKeyId , AccessKeySecret , EndPoint , Bucket , FileName , FileType , FileData , DownloadDoamin )

}

```






