/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-21 18:50:50
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-21 19:17:58
 * @Description: qiwei upload
 */
package upload

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	mytoolCommon "github.com/yangjerry110/mytool/common"
	mytoolCommonPkg "github.com/yangjerry110/mytool/common/pkg"
	mytoolHttp "github.com/yangjerry110/mytool/http"
	mytoolHttpPkg "github.com/yangjerry110/mytool/http/pkg"
)

type (
	QiweiUploadInterface interface {
		UploadMedia() error
	}

	QiweiUpload struct{}

	QiweiUploadMedia struct {
		AppId         string
		CropId        string
		CropSecret    string
		MediaId       string
		MediaData     string
		MediaType     string
		QiweiFilePath string
	}
)

/**
 * @description: UploadMedia
 * @author: Jerry.Yang
 * @date: 2022-09-21 18:56:23
 * @return {*}
 */
func (q *QiweiUploadMedia) UploadMedia() error {

	/**
	 * @step
	 * @判断appId
	**/
	if q.AppId == "" {
		return errors.New("QiweiUploadMedia Err : appId is empty!")
	}

	/**
	 * @step
	 * @判断qiweiFilePath
	 **/
	if q.QiweiFilePath == "" {
		return errors.New("QiweiUploadMedia Err : qiweiFilePath is empty!")
	}

	/**
	 * @step
	 * @获取accessToken
	 **/
	qiweiCommon := mytoolCommon.QiweiCommon{AppId: q.AppId, CropId: q.CropId, CropSecret: q.CropSecret}
	err := mytoolCommonPkg.GetQiweiAccessToken(&qiweiCommon)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @解密数据
	 **/
	decodMediaData, err := base64.StdEncoding.DecodeString(q.MediaData)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @构建body参数
	 **/
	body := &bytes.Buffer{}

	/**
	 * @step
	 * @实例化multipart
	 **/
	writer := multipart.NewWriter(body)

	/**
	 * @step
	 * @创建name
	 **/
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), uuid.New().String())

	/**
	 * @step
	 * @创建path
	 **/
	filePath := fmt.Sprintf("%s.%s", q.QiweiFilePath+fileName, q.MediaType)

	/**
	 * @step
	 * @创建multipart 文件字段
	 * @创建一个文件的write对象
	 **/
	formFileWriter, err := writer.CreateFormFile(fileName, filepath.Base(filePath))
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @写入传入的文件数据
	 **/
	formFileWriter.Write(decodMediaData)

	/**
	 * @step
	 * @close
	 **/
	err = writer.Close()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @构建请求url
	 **/
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s", qiweiCommon.AccessToken, q.MediaType)

	/**
	 * @step
	 * @构建返回体结构体
	 **/
	type UploadMediaOutput struct {
		ErrCode   int32  `json:"errcode"`
		ErrMsg    string `json:"errmsg"`
		Type      string `json:"type"`
		MediaId   string `json:"media_id"`
		CreatedAt string `json:"created_at"`
	}

	/**
	 * @step
	 * @发送请求
	 **/
	resp := &UploadMediaOutput{}
	/**
	 * @step
	 * @设置httpOption
	 **/
	mytoolHttpOptions := mytoolHttp.HttpOptions{}
	httpOptions := []mytoolHttp.HttpOptionFunc{
		mytoolHttpOptions.SetHeaders(map[string]string{
			"Content-Type": "multipart/form-data",
		}),
	}

	/**
	 * @step
	 * @组装请求参数
	 **/
	httpClient := mytoolHttp.HttpClient{
		Method:  "POST",
		Url:     url,
		Body:    body,
		Options: httpOptions,
		Output:  resp,
	}
	mytoolHttpPkg.HttpRequest(&httpClient)

	/**
	 * @step
	 * @判断结果
	 **/
	if resp.ErrCode != 0 {
		return errors.New(resp.ErrMsg)
	}

	/**
	 * @step
	 * @复制mediaId
	 **/
	q.MediaId = resp.MediaId
	return nil
}
