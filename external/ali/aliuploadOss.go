/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 14:05:52
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 17:27:46
 * @Description: ali oss
 */
package ali

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	aliyunOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
)

/**
 * @description: AliOssUpload Upload
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:08:04
 * @return {*}
 */
func (a *AliOssUpload) Upload() (string, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := a.CheckParams()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @创建client
	 **/
	client, err := a.CreateClient()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @创建bucket
	 **/
	bucket, err := a.CreateBucket(client)

	/**
	 * @step
	 * @objectName
	 **/
	objectName := fmt.Sprintf("%d_%s.%s", time.Now().Unix(), uuid.New().String(), a.FileType)

	/**
	 * @step
	 * @option
	 **/
	options := []aliyunOss.Option{
		aliyunOss.ObjectACL(aliyunOss.ACLPublicRead),
	}

	/**
	 * @step
	 * @解密base64数据
	 **/
	decodFileData, err := base64.StdEncoding.DecodeString(a.FileData)
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @上传bytes数据到oss
	 **/
	err = bucket.PutObject(objectName, bytes.NewReader(decodFileData), options...)
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @获取实际访问的oss地址
	 **/
	ossUrl := fmt.Sprintf("%s/%s", a.DownloadDoamin, objectName)
	return ossUrl, nil
}

/**
 * @description: AliOssUploadFormLocalFile Upload
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:34:56
 * @return {*}
 */
func (a AliOssUploadFromLocalFile) Upload() (string, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := a.CheckParams()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @创建client
	 **/
	client, err := a.CreateClient()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @创建bucket
	 **/
	bucket, err := a.CreateBucket(client)

	/**
	 * @step
	 * @objectName
	 **/
	objectName := fmt.Sprintf("%d_%s.%s", time.Now().Unix(), uuid.New().String(), a.FileType)

	/**
	 * @step
	 * @option
	 **/
	options := []aliyunOss.Option{
		aliyunOss.ObjectACL(aliyunOss.ACLPublicRead),
	}

	/**
	 * @step
	 * @获取文件数据
	 **/
	decodFileData, err := a.GetFileData()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @上传bytes数据到oss
	 **/
	err = bucket.PutObject(objectName, bytes.NewReader(decodFileData), options...)
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @获取实际访问的oss地址
	 **/
	ossUrl := fmt.Sprintf("%s/%s", a.DownloadDoamin, objectName)
	return ossUrl, nil
}

/**
 * @description: AliOssUpLoadFormFileUrl
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:38:03
 * @return {*}
 */
func (a *AliOssUpLoadFromFileUrl) Upload() (string, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := a.CheckParams()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @创建client
	 **/
	client, err := a.CreateClient()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @创建bucket
	 **/
	bucket, err := a.CreateBucket(client)

	/**
	 * @step
	 * @objectName
	 **/
	objectName := fmt.Sprintf("%d_%s.%s", time.Now().Unix(), uuid.New().String(), a.FileType)

	/**
	 * @step
	 * @option
	 **/
	options := []aliyunOss.Option{
		aliyunOss.ObjectACL(aliyunOss.ACLPublicRead),
	}

	/**
	 * @step
	 * @获取文件数据
	 **/
	decodFileData, err := a.GetFileData()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @上传bytes数据到oss
	 **/
	err = bucket.PutObject(objectName, bytes.NewReader(decodFileData), options...)
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @获取实际访问的oss地址
	 **/
	ossUrl := fmt.Sprintf("%s/%s", a.DownloadDoamin, objectName)
	return ossUrl, nil
}

/**
 * @description: CreateClient
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:12:39
 * @return {*}
 */
func (a *AliOssUpload) CreateClient() (*aliyunOss.Client, error) {
	return aliyunOss.New(a.EndPoint, a.AccessKeyId, a.AccessKeySecret)
}

/**
 * @description: CreateBucket
 * @param {*aliyunOss.Client} aliyunOssClient
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:12:47
 * @return {*}
 */
func (a *AliOssUpload) CreateBucket(aliyunOssClient *aliyunOss.Client) (*aliyunOss.Bucket, error) {
	return aliyunOssClient.Bucket(a.Bucket)
}

/**
 * @description: AliOssUploadFromLocalFile GetFileData
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:33:04
 * @return {*}
 */
func (a *AliOssUploadFromLocalFile) GetFileData() ([]byte, error) {
	/**
	 * @step
	 * @读取本地文件
	 **/
	res, err := os.Open(a.LocalFilePath)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	localFileBuf := bufio.NewReader(res)

	/**
	 * @step
	 * @读取本地文件二进制流
	 **/
	chunks := make([]byte, 0)
	buf := make([]byte, 1024) //一次读取多少个字节
	for {
		n, err := localFileBuf.Read(buf)

		/**
		 * @step
		 * @拼接内容
		 **/
		chunks = append(chunks, buf[:n]...)

		/**
		 * @step
		 * @报错
		 **/
		if err != nil && err != io.EOF {
			return nil, err
		}

		/**
		 * @step
		 * @读取完成
		 **/
		if n == 0 {
			break
		}
	}
	return chunks, nil
}

/**
 * @description: AliOssUpLoadFromFileUrl GetFileData
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:36:41
 * @return {*}
 */
func (a *AliOssUpLoadFromFileUrl) GetFileData() ([]byte, error) {
	/**
	 * @step
	 * @获取网络图片相关的资源
	 **/
	res, err := http.Get(a.FileUrl)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	/**
	 * @step
	 * @定义返回内容
	 **/
	var buffer []byte

	/**
	 * @step
	 * @读取网络图片的内容到byte中,每次读取1024
	 **/
	buf := make([]byte, 1024)
	for {
		n, err := res.Body.Read(buf)

		/**
		 * @step
		 * @拼接内容
		 **/
		buffer = append(buffer, buf[:n]...)

		/**
		 * @step
		 * @报错
		 **/
		if err != nil && err != io.EOF {
			return nil, err
		}

		/**
		 * @step
		 * @读取完成
		 **/
		if n == 0 {
			break
		}
	}
	return buffer, nil
}

/**
 * @description: CheckParams
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:43:53
 * @return {*}
 */
func (a *AliOssUpload) CheckParams() error {
	if a.FileName == "" {
		return errors.New("AliOssUpload Err : FileName is empty!")
	}

	if a.FileType == "" {
		return errors.New("AliOssUpload Err : FileType is empty!")
	}

	if a.FileData == "" {
		return errors.New("AliOssUpload Err : FileData is empty!")
	}

	if a.AccessKeyId == "" {
		return errors.New("AliOssUpload Err : AccessKeyId is empty!")
	}

	if a.AccessKeySecret == "" {
		return errors.New("AliOssUpload Err : AccessKeySecret is empty!")
	}

	if a.Bucket == "" {
		return errors.New("AliOssUpload Err : Bucket is empty!")
	}

	if a.EndPoint == "" {
		return errors.New("AliOssUpload Err : EndPoint is empty!")
	}

	if a.DownloadDoamin == "" {
		return errors.New("AliOssUpload Err : DownloadDoamin is empty!")
	}
	return nil
}

/**
 * @description: AliOssUploadFromLocalFile CheckParams
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:47:00
 * @return {*}
 */
func (a *AliOssUploadFromLocalFile) CheckParams() error {
	err := a.AliOssUpload.CheckParams()
	if err != nil {
		return err
	}

	if a.LocalFilePath == "" {
		return errors.New("AliOssUploadFormLocalFile Err : LocalFilePath is empty!")
	}
	return nil
}

/**
 * @description: AliOssUpLoadFromFileUrl CheckParams
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:47:51
 * @return {*}
 */
func (a *AliOssUpLoadFromFileUrl) CheckParams() error {
	err := a.AliOssUpload.CheckParams()
	if err != nil {
		return err
	}

	if a.FileUrl == "" {
		return errors.New("AliOssUpLoadFormFileUrl Err : FileUrl is empty!")
	}
	return nil
}
