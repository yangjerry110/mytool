/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 16:40:39
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 16:45:15
 * @Description: qiwei upload
 */
package upload

import "github.com/yangjerry110/mytool/upload"

type QiweiUploadInterface interface {
	QiweiUploadMedia(AppId string, CropId string, CropSecret string, MediaId string, MediaData string, MediaType string, QiweiFilePath string) (string, error)
}

type QiweiUpload struct{}

/**
 * @description: QiweiUploadMedia
 * @param {string} AppId
 * @param {string} CropId
 * @param {string} CropSecret
 * @param {string} MediaId
 * @param {string} MediaData
 * @param {string} MediaType
 * @param {string} QiweiFilePath
 * @author: Jerry.Yang
 * @date: 2022-09-22 16:45:55
 * @return {*}
 */
func QiweiUploadMedia(AppId string, CropId string, CropSecret string, MediaId string, MediaData string, MediaType string, QiweiFilePath string) (string, error) {
	qiweiUploadMediaObj := upload.QiweiUploadMedia{
		AppId:         AppId,
		CropId:        CropId,
		CropSecret:    CropSecret,
		MediaId:       MediaId,
		MediaData:     MediaData,
		MediaType:     MediaType,
		QiweiFilePath: QiweiFilePath,
	}
	err := qiweiUploadMediaObj.Upload()
	return qiweiUploadMediaObj.MediaId, err
}
