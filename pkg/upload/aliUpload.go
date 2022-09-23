/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 16:26:49
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-23 16:36:08
 * @Description: ali upload
 */
package upload

import "github.com/yangjerry110/mytool/upload"

type AliUploadPkgInterface interface {
	CreateAliUploadInterface(aliUploadInterface upload.AliUploadInterface) *UploadPkg
	CreateAliUploadOssInstance() upload.AliUploadInterface
	CreateAliUploadOssFromLocalFile() upload.AliUploadInterface
	CreateAliUploadFromFileUrl() upload.AliUploadInterface
}

type AliUploadPkg struct{}

type AliUploadOss struct {
	AccessKeyId     string
	AccessKeySecret string
	EndPoint        string
	Bucket          string
	FileName        string
	FileType        string
	FileData        string
	LocalFilePath   string
	FileUrl         string
	DownloadDoamin  string
}

/**
 * @description: CreateAliUploadInterface
 * @param {upload.AliUploadInterface} aliUploadInterface
 * @author: Jerry.Yang
 * @date: 2022-09-23 16:35:45
 * @return {*}
 */
func CreateAliUploadInterface(aliUploadInterface upload.AliUploadInterface) *UploadPkg {
	return &UploadPkg{AliUploadInterface: aliUploadInterface}
}

/**
 * @description: CreateAliUploadOssInstance
 * @author: Jerry.Yang
 * @date: 2022-09-23 16:35:53
 * @return {*}
 */
func (a *AliUploadOss) CreateAliUploadOssInstance() upload.AliUploadInterface {
	return CreateAliUploadInterface(&upload.AliOssUpload{
		AccessKeyId:     a.AccessKeyId,
		AccessKeySecret: a.AccessKeySecret,
		EndPoint:        a.EndPoint,
		Bucket:          a.Bucket,
		FileName:        a.FileName,
		FileType:        a.FileType,
		FileData:        a.FileData,
		DownloadDoamin:  a.DownloadDoamin,
	}).AliUploadInterface
}

/**
 * @description: CreateAliUploadOssFromLocalFile
 * @author: Jerry.Yang
 * @date: 2022-09-23 16:36:00
 * @return {*}
 */
func (a *AliUploadOss) CreateAliUploadOssFromLocalFile() upload.AliUploadInterface {
	return CreateAliUploadInterface(&upload.AliOssUploadFromLocalFile{
		AliOssUpload: upload.AliOssUpload{
			AccessKeyId:     a.AccessKeyId,
			AccessKeySecret: a.AccessKeySecret,
			EndPoint:        a.EndPoint,
			Bucket:          a.Bucket,
			FileName:        a.FileName,
			FileType:        a.FileType,
			FileData:        a.FileData,
			DownloadDoamin:  a.DownloadDoamin,
		},
		LocalFilePath: a.LocalFilePath,
	}).AliUploadInterface
}

/**
 * @description: CreateAliUploadFromFileUrl
 * @author: Jerry.Yang
 * @date: 2022-09-23 16:36:07
 * @return {*}
 */
func (a *AliUploadOss) CreateAliUploadFromFileUrl() upload.AliUploadInterface {
	return CreateAliUploadInterface(&upload.AliOssUpLoadFromFileUrl{
		AliOssUpload: upload.AliOssUpload{
			AccessKeyId:     a.AccessKeyId,
			AccessKeySecret: a.AccessKeySecret,
			EndPoint:        a.EndPoint,
			Bucket:          a.Bucket,
			FileName:        a.FileName,
			FileType:        a.FileType,
			FileData:        a.FileData,
			DownloadDoamin:  a.DownloadDoamin,
		},
		FileUrl: a.FileUrl,
	}).AliUploadInterface
}
