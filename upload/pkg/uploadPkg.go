/*
 * @Author: yangjie04@qutoutiao.net
 * @Date: 2022-09-21 19:18:09
 * @LastEditors: yangjie04@qutoutiao.net
 * @LastEditTime: 2022-09-22 14:39:58
 * @Description: qiwei upload
 */
package pkg

import "github.com/yangjerry110/mytool/upload"

type UploadPkgInterface interface {
	QiweiUploadMedia(qiweiUploadMedia *upload.QiweiUploadMedia) error
	AliOssUpload(AliOssUpload *upload.AliOssUpload) error
	AliOssUploadFromLocalFile(aliOssUploadFormLocalFile upload.AliOssUploadFormLocalFile) error
	AliOssUploadFormFileUrl(aliOssUploadFormFileUrl upload.AliOssUpLoadFormFileUrl) error
}

type UploadPkg struct{}

/**
 * @description: QiweiUploadMedia
 * @param {*upload.QiweiUploadMedia} qiweiUploadMedia
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-21 19:19:59
 * @return {*}
 */
func QiweiUploadMedia(qiweiUploadMedia *upload.QiweiUploadMedia) error {
	return qiweiUploadMedia.UploadMedia()
}

/**
 * @description: AliOssUpload
 * @param {*upload.AliOssUpload} AliOssUpload
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-22 14:20:11
 * @return {*}
 */
func AliOssUpload(AliOssUpload *upload.AliOssUpload) error {
	return AliOssUpload.Upload()
}

/**
 * @description: AliOssUploadFromLocalFile
 * @param {upload.AliOssUploadFormLocalFile} aliOssUploadFormLocalFile
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-22 14:39:09
 * @return {*}
 */
func AliOssUploadFromLocalFile(aliOssUploadFormLocalFile upload.AliOssUploadFormLocalFile) error {
	return aliOssUploadFormLocalFile.Upload()
}

/**
 * @description: AliOssUploadFormFileUrl
 * @param {upload.AliOssUpLoadFormFileUrl} aliOssUploadFormFileUrl
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-22 14:39:57
 * @return {*}
 */
func AliOssUploadFormFileUrl(aliOssUploadFormFileUrl upload.AliOssUpLoadFormFileUrl) error {
	return aliOssUploadFormFileUrl.Upload()
}
