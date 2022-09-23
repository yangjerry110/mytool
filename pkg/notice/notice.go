/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-23 15:14:48
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-23 15:58:24
 * @Description: notice
 */
package notice

import "github.com/yangjerry110/mytool/notice"

type NoticePkgInterface interface {
	QiweiNotice() (bool, error)
	QiweiBotNotice() (bool, error)
}

type NoticePkg struct {
	QiweiNoticeInterface    notice.QiweiNoticeInterface
	QiweiBotNoticeInterface notice.QiweiBotNoticeInterface
}

/**
 * @description: QiweiNotice
 * @author: Jerry.Yang
 * @date: 2022-09-23 15:42:06
 * @return {*}
 */
func (q *QiweiNoticePkg) QiweiNotice() (bool, error) {
	return q.CreateQiweiNoticeInstance().NotifyMessage()
}

/**
 * @description: QiweiBotNotice
 * @author: Jerry.Yang
 * @date: 2022-09-23 15:42:12
 * @return {*}
 */
func (q *QiweiBotNoticePkg) QiweiBotNotice() (bool, error) {
	return q.CreateQiweiBotNoticeInstance().NotifyMessage()
}
