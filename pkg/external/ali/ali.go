/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-26 18:43:31
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 18:47:04
 * @Description: ali
 */
package ali

import "github.com/yangjerry110/mytool/external/ali"

type AliInterface interface {
	CreateAliInterface(aliInterface ali.AliInterface) *Ali
}

/**
 * @description: Ali
 * @author: Jerry.Yang
 * @date: 2022-09-26 18:46:00
 * @return {*}
 */
type Ali struct {
	AliInterface ali.AliInterface
}

/**
 * @description: CreateAliInterface
 * @param {ali.AliInterface} aliInterface
 * @author: Jerry.Yang
 * @date: 2022-09-26 18:45:53
 * @return {*}
 */
func CreateAliInterface(aliInterface ali.AliInterface) *Ali {
	return &Ali{AliInterface: aliInterface}
}

/**
 * @description: AliUploadOss
 * @param {string} accessKeyId
 * @param {string} accessKeySecret
 * @param {string} endPoint
 * @param {string} bucket
 * @param {string} filename
 * @param {string} fileType
 * @param {string} fileData
 * @param {string} downloadDomain
 * @author: Jerry.Yang
 * @date: 2022-09-26 16:51:56
 * @return {*}
 */
func AliUploadOss(accessKeyId string, accessKeySecret string, endPoint string, bucket string, filename string, fileType string, fileData string, downloadDomain string) (string, error) {
	return CreateAliInterface(&ali.AliOssUpload{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret, EndPoint: endPoint, Bucket: bucket, FileName: filename, FileType: fileType, FileData: fileData, DownloadDoamin: downloadDomain}).AliInterface.Upload()
}

/**
 * @description: AliUploadOssFromLocaFile
 * @param {string} accessKeyId
 * @param {string} accessKeySecret
 * @param {string} endPoint
 * @param {string} bucket
 * @param {string} filename
 * @param {string} fileType
 * @param {string} fileData
 * @param {string} downloadDomain
 * @param {string} localFilePath
 * @author: Jerry.Yang
 * @date: 2022-09-26 16:55:44
 * @return {*}
 */
func AliUploadOssFromLocaFile(accessKeyId string, accessKeySecret string, endPoint string, bucket string, filename string, fileType string, fileData string, downloadDomain string, localFilePath string) (string, error) {
	return CreateAliInterface(&ali.AliOssUploadFromLocalFile{AliOssUpload: ali.AliOssUpload{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret, EndPoint: endPoint, Bucket: bucket, FileName: filename, FileType: fileType, FileData: fileData, DownloadDoamin: downloadDomain}, LocalFilePath: localFilePath}).AliInterface.Upload()
}

/**
 * @description: AliUploadOssFromFileUrl
 * @param {string} accessKeyId
 * @param {string} accessKeySecret
 * @param {string} endPoint
 * @param {string} bucket
 * @param {string} filename
 * @param {string} fileType
 * @param {string} fileData
 * @param {string} downloadDomain
 * @param {string} fileUrl
 * @author: Jerry.Yang
 * @date: 2022-09-26 16:55:55
 * @return {*}
 */
func AliUploadOssFromFileUrl(accessKeyId string, accessKeySecret string, endPoint string, bucket string, filename string, fileType string, fileData string, downloadDomain string, fileUrl string) (string, error) {
	return CreateAliInterface(&ali.AliOssUpLoadFromFileUrl{AliOssUpload: ali.AliOssUpload{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret, EndPoint: endPoint, Bucket: bucket, FileName: filename, FileType: fileType, FileData: fileData, DownloadDoamin: downloadDomain}, FileUrl: fileUrl}).AliInterface.Upload()
}
