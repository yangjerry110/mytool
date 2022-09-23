/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 16:40:39
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-23 16:41:47
 * @Description: qiwei upload
 */
package upload

import "github.com/yangjerry110/mytool/upload"

type QiweiUploadInterface interface {
	CreateQiweiUploadInterface(qiweiUploadInterface upload.QiweiUploadInterface) *UploadPkg
	CreateQiweiUploadMediaInstance() upload.QiweiUploadInterface
}

type QiweiUpload struct{}

type QiweiUploadMedia struct {
	AppId         string
	CropId        string
	CropSecret    string
	MediaData     string
	MediaType     string
	QiweiFilePath string
}

/**
 * @description: CreateQiweiUploadInterface
 * @param {upload.QiweiUploadInterface} qiweiUploadInterface
 * @author: Jerry.Yang
 * @date: 2022-09-23 16:41:36
 * @return {*}
 */
func CreateQiweiUploadInterface(qiweiUploadInterface upload.QiweiUploadInterface) *UploadPkg {
	return &UploadPkg{QiweiUploadInterface: qiweiUploadInterface}
}

/**
 * @description: CreateQiweiUploadMediaInstance
 * @author: Jerry.Yang
 * @date: 2022-09-23 16:41:46
 * @return {*}
 */
func (q *QiweiUploadMedia) CreateQiweiUploadMediaInstance() upload.QiweiUploadInterface {
	return CreateQiweiUploadInterface(&upload.QiweiUploadMedia{
		AppId:         q.AppId,
		CropId:        q.CropId,
		CropSecret:    q.CropSecret,
		MediaData:     q.MediaData,
		MediaType:     q.MediaType,
		QiweiFilePath: q.QiweiFilePath,
	}).AliUploadInterface
}
