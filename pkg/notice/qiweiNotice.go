/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 16:10:30
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-23 15:40:52
 * @Description: qiwei notice
 */
package notice

import "github.com/yangjerry110/mytool/notice"

type QiweiNoticePkgInterface interface {
	CreateQiweiNoticeInterface(qiweiNoticeInterface notice.QiweiNoticeInterface) *NoticePkg
	CreateQiweiNoticeInstance() notice.QiweiNoticeInterface
	CreateQiweiBotNoticeInterface(qiweiBotNoticeInterface notice.QiweiBotNoticeInterface) *NoticePkg
	CreateQiweiBotNoticeInstance() notice.QiweiBotNoticeInterface
}

type QiweiNoticePkg struct {
	AppId          string
	MsgType        string
	CropId         string
	CropSecret     string
	AgentId        string
	DepartmentIds  string
	TagIds         string
	UserIds        string
	Safe           int32
	SendMsg        string
	MediaData      string
	MediaType      string
	Title          string
	Description    string
	Url            string
	PicUrl         string
	EnableIdTrans  int32
	Btntxt         string
	AppletId       string
	AppletPagepath string
	QiweiFilePath  string
}

type QiweiBotNoticePkg struct {
	MsgType             string
	BotUrl              string
	SendMsg             string
	ImageData           string
	Title               string
	Description         string
	Url                 string
	PicUrl              string
	MentionedList       []string
	MentionedMobileList []string
}

/**
 * @description: CreateQiweiNoticeInterface
 * @param {notice.QiweiNoticeInterface} qiweiNoticeInterface
 * @author: Jerry.Yang
 * @date: 2022-09-23 15:40:59
 * @return {*}
 */
func (q *QiweiNoticePkg) CreateQiweiNoticeInterface(qiweiNoticeInterface notice.QiweiNoticeInterface) *NoticePkg {
	return &NoticePkg{QiweiNoticeInterface: qiweiNoticeInterface}
}

/**
 * @description: CreateQiweiNoticeInstance
 * @author: Jerry.Yang
 * @date: 2022-09-23 15:41:07
 * @return {*}
 */
func (q *QiweiNoticePkg) CreateQiweiNoticeInstance() notice.QiweiNoticeInterface {
	return q.CreateQiweiNoticeInterface(&notice.QiweiNotice{
		AppId:          q.AppId,
		MsgType:        q.MsgType,
		CropId:         q.CropId,
		CropSecret:     q.CropSecret,
		AgentId:        q.AgentId,
		DepartmentIds:  q.DepartmentIds,
		TagIds:         q.TagIds,
		UserIds:        q.UserIds,
		Safe:           q.Safe,
		SendMsg:        q.SendMsg,
		MediaData:      q.MediaData,
		MediaType:      q.MediaType,
		Title:          q.Title,
		Description:    q.Description,
		Url:            q.Url,
		PicUrl:         q.PicUrl,
		EnableIdTrans:  q.EnableIdTrans,
		Btntxt:         q.Btntxt,
		AppletId:       q.AppletId,
		AppletPagepath: q.AppletPagepath,
		QiweiFilePath:  q.QiweiFilePath,
	}).QiweiNoticeInterface
}

/**
 * @description: CreateQiweiBotNoticeInterface
 * @param {notice.QiweiBotNoticeInterface} qiweiBotNoticeInterface
 * @author: Jerry.Yang
 * @date: 2022-09-23 15:41:14
 * @return {*}
 */
func (q *QiweiBotNoticePkg) CreateQiweiBotNoticeInterface(qiweiBotNoticeInterface notice.QiweiBotNoticeInterface) *NoticePkg {
	return &NoticePkg{QiweiBotNoticeInterface: qiweiBotNoticeInterface}
}

/**
 * @description: CreateQiweiBotNoticeInstance
 * @author: Jerry.Yang
 * @date: 2022-09-23 15:41:20
 * @return {*}
 */
func (q *QiweiBotNoticePkg) CreateQiweiBotNoticeInstance() notice.QiweiBotNoticeInterface {
	return q.CreateQiweiBotNoticeInterface(&notice.QiweiBotNotice{
		MsgType:             q.MsgType,
		BotUrl:              q.BotUrl,
		SendMsg:             q.SendMsg,
		ImageData:           q.ImageData,
		Title:               q.Title,
		Description:         q.Description,
		Url:                 q.Url,
		PicUrl:              q.PicUrl,
		MentionedList:       q.MentionedList,
		MentionedMobileList: q.MentionedMobileList,
	}).QiweiBotNoticeInterface
}
