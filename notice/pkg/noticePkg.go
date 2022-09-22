/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-21 18:34:04
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 14:27:28
 * @Description: qiwei notice
 */
package pkg

import "github.com/yangjerry110/mytool/notice"

type NoticePkgInterface interface {
	QiweiNotifyMessage(qiweiNotice *notice.QiweiNotice) (bool, error)
	QiweiBotNotifyMessage(qiweiBotNotice *notice.QiweiBotNotice) (bool, error)
}

type NoticePkg struct{}

/**
 * @description: QiweiNotifyMessage
 * @param {*notice.QiweiNotice} qiweiNotice
 * @author: Jerry.Yang
 * @date: 2022-09-22 11:58:32
 * @return {*}
 */
func QiweiNotifyMessage(qiweiNotice *notice.QiweiNotice) (bool, error) {
	return qiweiNotice.NotifyMessage()
}

/**
 * @description: QiweiBotNotifyMessage
 * @param {*notice.QiweiBotNotice} qiweiBotNotice
 * @author: Jerry.Yang
 * @date: 2022-09-22 11:58:21
 * @return {*}
 */
func QiweiBotNotifyMessage(qiweiBotNotice *notice.QiweiBotNotice) (bool, error) {
	return qiweiBotNotice.NotifyMessage()
}
