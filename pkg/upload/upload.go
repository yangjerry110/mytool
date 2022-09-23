/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-23 16:16:52
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-23 16:41:04
 * @Description: upload
 */
package upload

import "github.com/yangjerry110/mytool/upload"

type UploadPkgInterface interface {
	AliUploadOss() (string, error)
	AliUploadOssFromLocalFile() (string, error)
	AliUploadOssFromFileUrl() (string, error)
	QiweiUploadMedia() (string, error)
}

type UploadPkg struct {
	AliUploadInterface   upload.AliUploadInterface
	QiweiUploadInterface upload.QiweiUploadInterface
}

/**
 * @description: AliUploadOss
 * @author: Jerry.Yang
 * @date: 2022-09-23 16:35:13
 * @return {*}
 */
func (a *AliUploadOss) AliUploadOss() (string, error) {
	return a.CreateAliUploadOssInstance().Upload()
}

/**
 * @description: AliUploadOssFromLocalFile
 * @author: Jerry.Yang
 * @date: 2022-09-23 16:35:21
 * @return {*}
 */
func (a *AliUploadOss) AliUploadOssFromLocalFile() (string, error) {
	return a.CreateAliUploadOssFromLocalFile().Upload()
}

/**
 * @description: AliUploadOssFromFileUrl
 * @author: Jerry.Yang
 * @date: 2022-09-23 16:35:28
 * @return {*}
 */
func (a *AliUploadOss) AliUploadOssFromFileUrl() (string, error) {
	return a.CreateAliUploadFromFileUrl().Upload()
}

/**
 * @description: QiweiUploadMedia
 * @author: Jerry.Yang
 * @date: 2022-09-23 16:41:12
 * @return {*}
 */
func (q *QiweiUploadMedia) QiweiUploadMedia() (string, error) {
	return q.CreateQiweiUploadMediaInstance().Upload()
}
