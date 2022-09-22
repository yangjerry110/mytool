/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 16:26:49
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 18:09:55
 * @Description: ali upload
 */
package upload

import "github.com/yangjerry110/mytool/upload"

type (
	AliUploadInterface interface {
		AliUploadOss(AccessKeyId string, AccessKeySecret string, EndPoint string, Bucket string, FileName string, FileType string, FileData string, DownloadDoamin string) (*AliUploadOssResp, error)
		AliUploadOssFromLocalFile(AccessKeyId string, AccessKeySecret string, EndPoint string, Bucket string, FileName string, FileType string, FileData string, DownloadDoamin string, localFilePath string) (*AliUploadOssResp, error)
		AliUploadOssFromFileUrl(AccessKeyId string, AccessKeySecret string, EndPoint string, Bucket string, FileName string, FileType string, FileData string, DownloadDoamin string, fileUrl string) (*AliUploadOssResp, error)
	}

	AliUpload struct{}

	AliUploadOssResp struct {
		FileName  string
		FileType  string
		Bucket    string
		AliOssUrl string
	}
)

/**
 * @description: AliUploadOss
 * @param {string} OssUrl
 * @param {string} AccessKeyId
 * @param {string} AccessKeySecret
 * @param {string} EndPoint
 * @param {string} Bucket
 * @param {string} FileName
 * @param {string} FileType
 * @param {string} FileData
 * @param {string} DownloadDoamin
 * @author: Jerry.Yang
 * @date: 2022-09-22 16:34:04
 * @return {*}
 */
func AliUploadOss(AccessKeyId string, AccessKeySecret string, EndPoint string, Bucket string, FileName string, FileType string, FileData string, DownloadDoamin string) (*AliUploadOssResp, error) {
	aliUploadOssObj := upload.AliOssUpload{
		AccessKeyId:     AccessKeyId,
		AccessKeySecret: AccessKeySecret,
		EndPoint:        EndPoint,
		Bucket:          Bucket,
		FileName:        FileName,
		FileType:        FileType,
		FileData:        FileData,
		DownloadDoamin:  DownloadDoamin,
	}
	err := aliUploadOssObj.Upload()
	return &AliUploadOssResp{
		FileName:  FileName,
		FileType:  FileType,
		Bucket:    Bucket,
		AliOssUrl: aliUploadOssObj.OssUrl,
	}, err
}

/**
 * @description: AliUploadOssFromLocalFile
 * @param {string} AccessKeyId
 * @param {string} AccessKeySecret
 * @param {string} EndPoint
 * @param {string} Bucket
 * @param {string} FileName
 * @param {string} FileType
 * @param {string} FileData
 * @param {string} DownloadDoamin
 * @param {string} localFilePath
 * @author: Jerry.Yang
 * @date: 2022-09-22 16:39:30
 * @return {*}
 */
func AliUploadOssFromLocalFile(AccessKeyId string, AccessKeySecret string, EndPoint string, Bucket string, FileName string, FileType string, FileData string, DownloadDoamin string, localFilePath string) (*AliUploadOssResp, error) {
	aliUploadOssFormLocalFileObj := upload.AliOssUploadFromLocalFile{
		AliOssUpload: upload.AliOssUpload{
			AccessKeyId:     AccessKeyId,
			AccessKeySecret: AccessKeySecret,
			EndPoint:        EndPoint,
			Bucket:          Bucket,
			FileName:        FileName,
			FileType:        FileType,
			FileData:        FileData,
			DownloadDoamin:  DownloadDoamin,
		},
		LocalFilePath: localFilePath,
	}
	err := aliUploadOssFormLocalFileObj.Upload()
	return &AliUploadOssResp{
		FileName:  FileName,
		FileType:  FileType,
		Bucket:    Bucket,
		AliOssUrl: aliUploadOssFormLocalFileObj.OssUrl,
	}, err
}

/**
 * @description: AliUploadOssFromFileUrl
 * @param {string} AccessKeyId
 * @param {string} AccessKeySecret
 * @param {string} EndPoint
 * @param {string} Bucket
 * @param {string} FileName
 * @param {string} FileType
 * @param {string} FileData
 * @param {string} DownloadDoamin
 * @param {string} fileUrl
 * @author: Jerry.Yang
 * @date: 2022-09-22 16:39:39
 * @return {*}
 */
func AliUploadOssFromFileUrl(AccessKeyId string, AccessKeySecret string, EndPoint string, Bucket string, FileName string, FileType string, FileData string, DownloadDoamin string, fileUrl string) (*AliUploadOssResp, error) {
	aliUploadOssFormFileUrlObj := upload.AliOssUpLoadFromFileUrl{
		AliOssUpload: upload.AliOssUpload{
			AccessKeyId:     AccessKeyId,
			AccessKeySecret: AccessKeySecret,
			EndPoint:        EndPoint,
			Bucket:          Bucket,
			FileName:        FileName,
			FileType:        FileType,
			FileData:        FileData,
			DownloadDoamin:  DownloadDoamin,
		},
		FileUrl: fileUrl,
	}
	err := aliUploadOssFormFileUrlObj.Upload()
	return &AliUploadOssResp{
		FileName:  FileName,
		FileType:  FileType,
		Bucket:    Bucket,
		AliOssUrl: aliUploadOssFormFileUrlObj.OssUrl,
	}, err
}
