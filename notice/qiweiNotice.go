/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-21 15:30:12
 * @LastEditors: yangjie04@qutoutiao.net
 * @LastEditTime: 2022-09-22 14:51:06
 * @Description: 企微通知
 */
package notice

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	mytoolCommon "github.com/yangjerry110/mytool/common"
	mytoolCommonPkg "github.com/yangjerry110/mytool/common/pkg"
	mytoolHttp "github.com/yangjerry110/mytool/http"
	mytoolHttpPkg "github.com/yangjerry110/mytool/http/pkg"
	mytoolUpload "github.com/yangjerry110/mytool/upload"
	mytoolUploadPkg "github.com/yangjerry110/mytool/upload/pkg"
)

type (
	QiweiNoticeInterface interface {
		CheckParams() error
		NotifyMessage() (bool, error)
		FormatNotifyParams(qiweiNotice *QiweiNotice) ([]byte, error)
		DoNotify(accessToken string, qiweiMsg []byte) (bool, error)
	}

	QiweiNotice struct {
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

	/**
	 * @step
	 * @定义基础的请求结构
	 **/
	NotifyMessage struct {
		Touser                 string `json:"touser"`
		Toparty                string `json:"toparty"`
		Totag                  string `json:"totag"`
		Msgtype                string `json:"msgtype"`
		Agentid                string `json:"agentid"`
		EnableIdTrans          int32  `json:"enable_id_trans"`
		EnableDuplicateCheck   int32  `json:"enable_duplicate_check"`
		DuplicateCheckInterval int32  `json:"duplicate_check_interval"`
	}

	/**
	 * @step
	 * @定义textMessage的结构
	 **/
	NotifyTextMessage struct{}

	/**
	 * @step
	 * @定义markdownMessage的结构
	 **/
	NotifyMarkdownMessage struct{}

	/**
	 * @step
	 * @定义imageMessage的结构
	 **/
	NotifyImageMessage struct{}

	/**
	 * @step
	 * @定义cardMessage的结构
	 **/
	NotifyCardMessage struct{}

	/**
	 * @step
	 * @定义newsMessage的结构
	 **/
	NotifyNewsMessage struct{}
)

/**
 * @description: NotifyMessage
 * @author: Jerry.Yang
 * @date: 2022-09-21 18:33:53
 * @return {*}
 */
func (q *QiweiNotice) NotifyMessage() (bool, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := q.CheckParams()
	if err != nil {
		return false, err
	}

	/**
	 * @step
	 * @初始化一个通道,获取企微请求参数
	 **/
	notifyQiweiReqChan := make(chan []byte)
	notifyQiweiReqErrChan := make(chan error)

	/**
	 * @step
	 * @获取accessToken
	 **/
	qiweiCommon := mytoolCommon.QiweiCommon{AppId: q.AppId, CropId: q.CropId, CropSecret: q.CropSecret}
	err = mytoolCommonPkg.GetQiweiAccessToken(&qiweiCommon)
	if err != nil {
		return false, err
	}

	/**
	 * @step
	 * @开启一个协程处理
	 **/
	go func() {

		/**
		 * @step
		 * @根据msgType的不同，调用不同的方法
		 **/
		switch q.MsgType {
		case "text":
			notifyTextMessage := NotifyTextMessage{}
			notifyQiweiReq, err := notifyTextMessage.FormatNotifyParams(q)
			notifyQiweiReqChan <- notifyQiweiReq
			notifyQiweiReqErrChan <- err
		case "markdown":
			notifyMarkdownMessage := NotifyMarkdownMessage{}
			notifyQiweiReq, err := notifyMarkdownMessage.FormatNofifyParams(q)
			notifyQiweiReqChan <- notifyQiweiReq
			notifyQiweiReqErrChan <- err
		case "image":
			notifyImageMessage := NotifyImageMessage{}
			notifyQiweiReq, err := notifyImageMessage.FormatNofifyParams(q)
			notifyQiweiReqChan <- notifyQiweiReq
			notifyQiweiReqErrChan <- err
		case "card":
			notifyCardMessage := NotifyCardMessage{}
			notifyQiweiReq, err := notifyCardMessage.FormatNofifyParams(q)
			notifyQiweiReqChan <- notifyQiweiReq
			notifyQiweiReqErrChan <- err
		case "news":
			notifyNewsMessage := NotifyNewsMessage{}
			notifyQiweiReq, err := notifyNewsMessage.FormatNofifyParams(q)
			notifyQiweiReqChan <- notifyQiweiReq
			notifyQiweiReqErrChan <- err
		default:
			notifyQiweiReqChan <- nil
			notifyQiweiReqErrChan <- errors.New("QiweiNotice Err : no match msgType!")
		}
		close(notifyQiweiReqChan)
	}()

	/**
	 * @step
	 * @获取请求企微通知的参数
	 **/
	notifyQiweiReq := <-notifyQiweiReqChan

	/**
	 * @step
	 * @获取组装参数的错误
	 **/
	notifyQiweiReqErr := <-notifyQiweiReqErrChan
	if notifyQiweiReqErr != nil {
		return false, notifyQiweiReqErr
	}

	/**
	 * @step
	 * @判断参数
	 **/
	if notifyQiweiReq == nil {
		return false, errors.New("QiweiNotice Err : notifyQiweiReq is nil")
	}

	/**
	 * @step
	 * @发送企微消息
	 **/
	sendMsgResult, err := q.DoNotify(qiweiCommon.AccessToken, notifyQiweiReq)
	if err != nil {
		return sendMsgResult, err
	}
	return true, nil
}

/**
 * @description: FormatNotifyParams
 * @param {*QiweiNotice} qiweiNotice
 * @author: Jerry.Yang
 * @date: 2022-09-21 18:14:49
 * @return {*}
 */
func (n *NotifyTextMessage) FormatNotifyParams(qiweiNotice *QiweiNotice) ([]byte, error) {

	/**
	 * @step
	 * @判断参数
	 **/
	err := n.CheckParams(qiweiNotice)
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @定义参数格式,text
	 **/
	type StructTextParams struct {
		Context string `json:"content"`
	}

	type StructParams struct {
		NotifyMessage
		Safe int32            `json:"safe"`
		Text StructTextParams `json:"text"`
	}

	/**
	 * @step
	 * @组合参数
	 **/
	jsonParams, err := json.Marshal(StructParams{
		NotifyMessage: NotifyMessage{
			Touser:  qiweiNotice.UserIds,
			Toparty: qiweiNotice.DepartmentIds,
			Totag:   qiweiNotice.TagIds,
			Msgtype: "text",
			Agentid: qiweiNotice.AgentId,
		},
		Safe: qiweiNotice.Safe,
		Text: StructTextParams{qiweiNotice.SendMsg},
	})

	/**
	 * @step
	 * @判断
	 **/
	if err != nil {
		return nil, err
	}
	return jsonParams, nil
}

/**
 * @description: FormatNofifyParams
 * @param {*QiweiNotice} qiweiNotice
 * @author: Jerry.Yang
 * @date: 2022-09-21 18:16:36
 * @return {*}
 */
func (n *NotifyMarkdownMessage) FormatNofifyParams(qiweiNotice *QiweiNotice) ([]byte, error) {

	/**
	 * @step
	 * @判断参数
	 **/
	err := n.CheckParams(qiweiNotice)
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @定义参数格式,text
	 **/
	type StructTextParams struct {
		Context string `json:"content"`
	}

	type StructParams struct {
		NotifyMessage
		Safe     int32            `json:"safe"`
		Markdown StructTextParams `json:"markdown"`
	}

	/**
	 * @step
	 * @组合参数
	 **/
	jsonParams, err := json.Marshal(StructParams{
		NotifyMessage: NotifyMessage{
			Touser:  qiweiNotice.UserIds,
			Toparty: qiweiNotice.DepartmentIds,
			Totag:   qiweiNotice.TagIds,
			Msgtype: "markdown",
			Agentid: qiweiNotice.AgentId,
		},
		Safe:     qiweiNotice.Safe,
		Markdown: StructTextParams{qiweiNotice.SendMsg},
	})

	/**
	 * @step
	 * @判断
	 **/
	if err != nil {
		return nil, err
	}
	return jsonParams, nil
}

/**
 * @description: FormatNofifyParams
 * @param {*QiweiNotice} qiweiNotice
 * @author: Jerry.Yang
 * @date: 2022-09-21 18:19:21
 * @return {*}
 */
func (n *NotifyImageMessage) FormatNofifyParams(qiweiNotice *QiweiNotice) ([]byte, error) {

	/**
	 * @step
	 * @判断参数
	 **/
	err := n.CheckParams(qiweiNotice)
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @获取mediaId
	 **/
	qiweiUploadMedia := mytoolUpload.QiweiUploadMedia{
		AppId:         qiweiNotice.AppId,
		CropId:        qiweiNotice.CropId,
		CropSecret:    qiweiNotice.CropSecret,
		MediaData:     qiweiNotice.MediaData,
		MediaType:     qiweiNotice.MediaType,
		QiweiFilePath: qiweiNotice.QiweiFilePath,
	}
	err = mytoolUploadPkg.QiweiUploadMedia(&qiweiUploadMedia)
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @定义参数格式
	 **/
	type StructMediaParams struct {
		MediaId string `json:"media_id"`
	}

	type StructParams struct {
		NotifyMessage
		Image StructMediaParams `json:"image"`
		Safe  int32             `json:"safe"`
	}

	/**
	 * @step
	 * @组合参数
	 **/
	jsonParams, err := json.Marshal(StructParams{
		NotifyMessage: NotifyMessage{
			Touser:  qiweiNotice.UserIds,
			Toparty: qiweiNotice.DepartmentIds,
			Totag:   qiweiNotice.TagIds,
			Msgtype: "image",
			Agentid: qiweiNotice.AgentId,
		},
		Safe:  qiweiNotice.Safe,
		Image: StructMediaParams{MediaId: qiweiUploadMedia.MediaId},
	})

	/**
	 * @step
	 * @判断
	 **/
	if err != nil {
		return nil, err
	}
	return jsonParams, nil
}

/**
 * @description: FormatNofifyParams
 * @param {*QiweiNotice} qiweiNotice
 * @author: Jerry.Yang
 * @date: 2022-09-21 18:21:55
 * @return {*}
 */
func (n *NotifyCardMessage) FormatNofifyParams(qiweiNotice *QiweiNotice) ([]byte, error) {

	/**
	 * @step
	 * @判断参数
	 **/
	err := n.CheckParams(qiweiNotice)
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @定义参数格式
	 **/
	type StructCardParams struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		Btntxt      string `json:"btntxt"`
	}

	type StructParams struct {
		NotifyMessage
		Textcard               StructCardParams `json:"textcard"`
		Safe                   int32            `json:"safe"`
		EnableIdTrans          int32            `json:"enable_id_trans"`
		EnableDuplicateCheck   int32            `json:"enable_duplicate_check"`
		DuplicateCheckInterval int32            `json:"duplicate_check_interval"`
	}

	/**
	 * @step
	 * @组合参数
	 **/
	jsonParams, err := json.Marshal(StructParams{
		NotifyMessage: NotifyMessage{
			Touser:  qiweiNotice.UserIds,
			Toparty: qiweiNotice.DepartmentIds,
			Totag:   qiweiNotice.TagIds,
			Msgtype: "testcard",
			Agentid: qiweiNotice.AgentId,
		},
		Textcard: StructCardParams{
			Title:       qiweiNotice.Title,
			Description: qiweiNotice.Description,
			Url:         qiweiNotice.Url,
			Btntxt:      qiweiNotice.Btntxt,
		},
		EnableIdTrans: qiweiNotice.EnableIdTrans,
	})

	/**
	 * @step
	 * @判断
	 **/
	if err != nil {
		return nil, err
	}
	return jsonParams, nil
}

/**
 * @description: FormatNofifyParams
 * @param {*QiweiNotice} qiweiNotice
 * @author: Jerry.Yang
 * @date: 2022-09-21 18:22:46
 * @return {*}
 */
func (n *NotifyNewsMessage) FormatNofifyParams(qiweiNotice *QiweiNotice) ([]byte, error) {

	/**
	 * @step
	 * @判断参数
	 **/
	err := n.CheckParams(qiweiNotice)
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @定义参数格式
	 **/
	type StructNewParams struct {
		Title          string `json:"title"`
		Description    string `json:"description"`
		NewUrl         string `json:"url"`
		PicUrl         string `json:"picurl"`
		AppletId       string `json:"appid"`
		AppletPagepath string `json:"pagepath"`
	}

	/**
	 * @step
	 * @定义articles的结构
	 **/
	type StructArticlesParams struct {
		Articles []StructNewParams `json:"articles"`
	}

	type StructParams struct {
		NotifyMessage
		News                   StructArticlesParams `json:"news"`
		Safe                   int32                `json:"safe"`
		EnableIdTrans          int32                `json:"enable_id_trans"`
		EnableDuplicateCheck   int32                `json:"enable_duplicate_check"`
		DuplicateCheckInterval int32                `json:"duplicate_check_interval"`
	}

	/**
	 * @step
	 * @make structNewParams
	 **/
	newsParams := []StructNewParams{{
		Title:          qiweiNotice.Title,
		Description:    qiweiNotice.Description,
		NewUrl:         qiweiNotice.Url,
		PicUrl:         qiweiNotice.PicUrl,
		AppletId:       qiweiNotice.AppletId,
		AppletPagepath: qiweiNotice.AppletPagepath,
	}}

	/**
	 * @step
	 * @组合参数
	 **/
	jsonParams, err := json.Marshal(StructParams{
		NotifyMessage: NotifyMessage{
			Touser:  qiweiNotice.UserIds,
			Toparty: qiweiNotice.DepartmentIds,
			Totag:   qiweiNotice.TagIds,
			Msgtype: "news",
			Agentid: qiweiNotice.AgentId,
		},
		News:          StructArticlesParams{Articles: newsParams},
		EnableIdTrans: qiweiNotice.EnableIdTrans,
	})

	/**
	 * @step
	 * @判断
	 **/
	if err != nil {
		return nil, err
	}
	return jsonParams, nil
}

/**
 * @description: DoNotify
 * @param {string} accessToken
 * @param {[]byte} qiweiMsg
 * @author: Jerry.Yang
 * @date: 2022-09-21 18:33:24
 * @return {*}
 */
func (q *QiweiNotice) DoNotify(accessToken string, qiweiMsg []byte) (bool, error) {

	/**
	 * @step
	 * @url
	 **/
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", accessToken)

	/**
	 * @step
	 * @返回数据 struct
	 **/
	type NotifyOutput struct {
		ErrCode int32  `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}

	/**
	 * @step
	 * @发送消息
	 **/
	resp := &NotifyOutput{}

	/**
	 * @step
	 * @设置httpOption
	 **/
	mytoolHttpOptions := mytoolHttp.HttpOptions{}
	httpOptions := []mytoolHttp.HttpOptionFunc{
		mytoolHttpOptions.SetHeaders(map[string]string{
			"Content-Type": "application/json",
		}),
	}

	/**
	 * @step
	 * @组装请求参数
	 **/
	httpClient := mytoolHttp.HttpClient{
		Method:  "POST",
		Url:     url,
		Body:    bytes.NewBuffer(qiweiMsg),
		Options: httpOptions,
		Output:  resp,
	}
	mytoolHttpPkg.HttpRequest(&httpClient)

	/**
	 * @step
	 * @判断错误
	 **/
	if resp.ErrCode != 0 {
		return false, errors.New(resp.ErrMsg)
	}
	return true, nil
}

/**
 * @description: CheckParams
 * @author: Jerry.Yang
 * @date: 2022-09-21 19:10:08
 * @return {*}
 */
func (q *QiweiNotice) CheckParams() error {
	if q.AppId == "" {
		return errors.New("QiweiNotice Err : appId is empty!")
	}

	if q.CropId == "" {
		return errors.New("QiweiNotice Err : CropId is empty!")
	}

	if q.CropSecret == "" {
		return errors.New("QiweiNotice Err : CropSecret is empty!")
	}

	if q.UserIds == "" {
		return errors.New("QiweiNotice Err : UserIds is empty!")
	}

	if q.MsgType == "" {
		return errors.New("QiweiNotice Err : MsgType is empty!")
	}

	if q.AgentId == "" {
		return errors.New("QiweiNotice Err : AgentId is empty!")
	}
	return nil
}

/**
 * @description: CheckParams
 * @param {*QiweiNotice} qiweiNotice
 * @author: Jerry.Yang
 * @date: 2022-09-21 19:16:13
 * @return {*}
 */
func (n *NotifyTextMessage) CheckParams(qiweiNotice *QiweiNotice) error {
	if qiweiNotice.SendMsg == "" {
		return errors.New("QiweiNotice NotifyTextMessage Err : sendMsg is empty!")
	}
	return nil
}

/**
 * @description: CheckParams
 * @param {*QiweiNotice} qiweiNotice
 * @author: Jerry.Yang
 * @date: 2022-09-21 19:16:59
 * @return {*}
 */
func (n *NotifyMarkdownMessage) CheckParams(qiweiNotice *QiweiNotice) error {
	if qiweiNotice.SendMsg == "" {
		return errors.New("QiweiNotice NotifyTextMessage Err : sendMsg is empty!")
	}
	return nil
}

/**
 * @description: CheckParams
 * @param {*QiweiNotice} qiweiNotice
 * @author: Jerry.Yang
 * @date: 2022-09-22 10:43:17
 * @return {*}
 */
func (n *NotifyImageMessage) CheckParams(qiweiNotice *QiweiNotice) error {
	if qiweiNotice.MediaData == "" {
		return errors.New("QiweiNotice NotifyImageMessage Err : MediaData is empty!")
	}

	if qiweiNotice.MediaType == "" {
		return errors.New("QiweiNotice NotifyImageMessage Err : MediaType is empty!")
	}

	if qiweiNotice.QiweiFilePath == "" {
		return errors.New("QiweiNotice NotifyImageMessage Err : QiweiFilePath is empty!")
	}
	return nil
}

/**
 * @description: CheckParams
 * @param {*QiweiNotice} qiweiNotice
 * @author: Jerry.Yang
 * @date: 2022-09-22 10:45:18
 * @return {*}
 */
func (n *NotifyCardMessage) CheckParams(qiweiNotice *QiweiNotice) error {
	if qiweiNotice.Title == "" {
		return errors.New("QiweiNotice NotifyCardMessage Err : Title is empty!")
	}

	if qiweiNotice.Description == "" {
		return errors.New("QiweiNotice NotifyCardMessage Err : Description is empty!")
	}

	if qiweiNotice.Url == "" {
		return errors.New("QiweiNotice NotifyCardMessage Err : Url is empty!")
	}

	if qiweiNotice.Btntxt == "" {
		return errors.New("QiweiNotice NotifyCardMessage Err : Btntxt is empty!")
	}
	return nil
}

/**
 * @description: CheckParams
 * @param {*QiweiNotice} qiweiNotice
 * @author: Jerry.Yang
 * @date: 2022-09-22 10:47:29
 * @return {*}
 */
func (n *NotifyNewsMessage) CheckParams(qiweiNotice *QiweiNotice) error {
	if qiweiNotice.Title == "" {
		return errors.New("QiweiNotice NotifyNewsMessage Err : Title is empty!")
	}

	if qiweiNotice.Description == "" {
		return errors.New("QiweiNotice NotifyNewsMessage Err : Description is empty!")
	}

	if qiweiNotice.PicUrl == "" {
		return errors.New("QiweiNotice NotifyNewsMessage Err : PicUrl is empty!")
	}

	if qiweiNotice.Url == "" {
		return errors.New("QiweiNotice NotifyNewsMessage Err : NewUrl is empty!")
	}
	return nil
}
