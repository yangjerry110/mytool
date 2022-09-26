/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-26 17:05:18
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 18:44:21
 * @Description: ali
 */
package ali

type AliInterface interface {
	Upload() (string, error)
}

/**
 * @description: Ali
 * @author: Jerry.Yang
 * @date: 2022-09-26 18:44:29
 * @return {*}
 */
type Ali struct{}

/**
 * @description: AliOssUpload
 * @author: Jerry.Yang
 * @date: 2022-09-26 18:44:09
 * @return {*}
 */
type AliOssUpload struct {
	AccessKeyId     string
	AccessKeySecret string
	EndPoint        string
	Bucket          string
	FileName        string
	FileType        string
	FileData        string
	DownloadDoamin  string
}

/**
 * @description: AliOssUploadFromLocalFile
 * @author: Jerry.Yang
 * @date: 2022-09-26 18:44:16
 * @return {*}
 */
type AliOssUploadFromLocalFile struct {
	AliOssUpload
	LocalFilePath string
}

/**
 * @description: AliOssUpLoadFromFileUrl
 * @author: Jerry.Yang
 * @date: 2022-09-26 18:44:20
 * @return {*}
 */
type AliOssUpLoadFromFileUrl struct {
	AliOssUpload
	FileUrl string
}
