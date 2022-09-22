/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 16:10:30
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 16:20:19
 * @Description: qiwei notice
 */
package notice

import "github.com/yangjerry110/mytool/notice"

type QiweiNoticeInterface interface {
	QiweiNotifyMessage(
		AppId string,
		MsgType string,
		CropId string,
		CropSecret string,
		AgentId string,
		DepartmentIds string,
		TagIds string,
		UserIds string,
		Safe int32,
		SendMsg string,
		MediaData string,
		MediaType string,
		Title string,
		Description string,
		Url string,
		PicUrl string,
		EnableIdTrans int32,
		Btntxt string,
		AppletId string,
		AppletPagepath string,
		QiweiFilePath string,
	) (bool, error)
	QiweiBotNotifyMessage(
		MsgType string,
		BotUrl string,
		SendMsg string,
		ImageData string,
		Title string,
		Description string,
		Url string,
		PicUrl string,
		MentionedList []string,
		MentionedMobileList []string,
	) (bool, error)
}

type QiweiNotice struct{}

/**
 * @description: QiweiNotifyMessage
 * @param {string} AppId
 * @param {string} MsgType
 * @param {string} CropId
 * @param {string} CropSecret
 * @param {string} AgentId
 * @param {string} DepartmentIds
 * @param {string} TagIds
 * @param {string} UserIds
 * @param {int32} Safe
 * @param {string} SendMsg
 * @param {string} MediaData
 * @param {string} MediaType
 * @param {string} Title
 * @param {string} Description
 * @param {string} Url
 * @param {string} PicUrl
 * @param {int32} EnableIdTrans
 * @param {string} Btntxt
 * @param {string} AppletId
 * @param {string} AppletPagepath
 * @param {string} QiweiFilePath
 * @author: Jerry.Yang
 * @date: 2022-09-22 16:17:05
 * @return {*}
 */
func QiweiNotifyMessage(AppId string, MsgType string, CropId string, CropSecret string, AgentId string, DepartmentIds string, TagIds string, UserIds string, Safe int32, SendMsg string, MediaData string, MediaType string, Title string, Description string, Url string, PicUrl string, EnableIdTrans int32, Btntxt string, AppletId string, AppletPagepath string, QiweiFilePath string) (bool, error) {
	qiweiNoticeObj := notice.QiweiNotice{
		AppId:          AppId,
		MsgType:        MsgType,
		CropId:         CropId,
		CropSecret:     CropSecret,
		AgentId:        AgentId,
		DepartmentIds:  DepartmentIds,
		TagIds:         TagIds,
		UserIds:        UserIds,
		Safe:           Safe,
		SendMsg:        SendMsg,
		MediaData:      MediaData,
		MediaType:      MediaType,
		Title:          Title,
		Description:    Description,
		Url:            Url,
		PicUrl:         PicUrl,
		EnableIdTrans:  EnableIdTrans,
		Btntxt:         Btntxt,
		AppletId:       AppletId,
		AppletPagepath: AppletPagepath,
		QiweiFilePath:  QiweiFilePath,
	}
	return qiweiNoticeObj.NotifyMessage()
}

/**
 * @description: QiweiBotNotifyMessage
 * @param {string} MsgType
 * @param {string} BotUrl
 * @param {string} SendMsg
 * @param {string} ImageData
 * @param {string} Title
 * @param {string} Description
 * @param {string} Url
 * @param {string} PicUrl
 * @param {[]string} MentionedList
 * @param {[]string} MentionedMobileList
 * @author: Jerry.Yang
 * @date: 2022-09-22 16:20:18
 * @return {*}
 */
func QiweiBotNotifyMessage(MsgType string, BotUrl string, SendMsg string, ImageData string, Title string, Description string, Url string, PicUrl string, MentionedList []string, MentionedMobileList []string) (bool, error) {
	qiweiBotNoticeObj := notice.QiweiBotNotice{
		MsgType:             MsgType,
		BotUrl:              BotUrl,
		SendMsg:             SendMsg,
		ImageData:           ImageData,
		Title:               Title,
		Description:         Description,
		Url:                 Url,
		PicUrl:              PicUrl,
		MentionedList:       MentionedList,
		MentionedMobileList: MentionedMobileList,
	}
	return qiweiBotNoticeObj.NotifyMessage()
}
