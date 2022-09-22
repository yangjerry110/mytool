/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 16:10:30
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 18:14:20
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
 * @param {string} AppId appId
 * @param {string} MsgType 消息类型
 * @param {string} CropId 公司id
 * @param {string} CropSecret 公司秘钥
 * @param {string} AgentId 通知的qiwei应用的id
 * @param {string} DepartmentIds 通知的组织架构集合
 * @param {string} TagIds 通知的Tag集合
 * @param {string} UserIds 通知的人
 * @param {int32} Safe safe
 * @param {string} SendMsg 通知的消息
 * @param {string} MediaData 通知的媒体消息内容
 * @param {string} MediaType 通知的媒体消息类型
 * @param {string} Title 标题
 * @param {string} Description 简介
 * @param {string} Url 链接
 * @param {string} PicUrl 图片链接
 * @param {int32} EnableIdTrans
 * @param {string} Btntxt news消息时候的按钮
 * @param {string} AppletId AppletId 小程序id
 * @param {string} AppletPagepath AppletPagepath 小程序链接
 * @param {string} QiweiFilePath 通知媒体消息的时候，存放媒体内容的地址
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
 * @param {string} MsgType 消息类型
 * @param {string} BotUrl 通知的机器人的地址
 * @param {string} SendMsg 通知的消息内容
 * @param {string} ImageData 通知媒体消息的媒体内容
 * @param {string} Title 标题
 * @param {string} Description 简介
 * @param {string} Url url
 * @param {string} PicUrl 图片链接
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
