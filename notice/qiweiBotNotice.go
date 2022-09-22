/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 11:28:53
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 11:59:17
 * @Description: qiwei bot notice
 */
package notice

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"

	mytoolHttp "github.com/yangjerry110/mytool/http"
	mytoolHttpPkg "github.com/yangjerry110/mytool/http/pkg"
)

type (
	QiweiBotNoticeInterface interface {
		CheckParams() error
		NotifyMessage() (bool, error)
		FormatNotifyParams(qiweiNotice *QiweiNotice) ([]byte, error)
		DoNotify(accessToken string, qiweiMsg []byte) error
	}

	QiweiBotNotice struct {
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
	 * @step
	 * @定义botTextMessage
	 **/
	NotifyBotTextMessage struct{}

	/**
	 * @step
	 * @定义botMarkdownMessage
	 **/
	NotifyBotMarkdownMessage struct{}

	/**
	 * @step
	 * @定义botImageMessage
	 **/
	NotifyBotImageMessage struct{}

	/**
	 * @step
	 * @定义botNewsMessage
	 **/
	NotifyBotNewsMessage struct{}
)

/**
 * @description: NotifyMessage
 * @author: Jerry.Yang
 * @date: 2022-09-22 11:34:14
 * @return {*}
 */
func (q *QiweiBotNotice) NotifyMessage() (bool, error) {

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
	 * @开启一个协程，获取不同的发送的参数
	 **/
	go func() {
		switch q.MsgType {
		case "text":
			notifyBotTextMessage := NotifyBotTextMessage{}
			notifyQiweiReq, err := notifyBotTextMessage.FormatNotifyBotParams(q)
			notifyQiweiReqChan <- notifyQiweiReq
			notifyQiweiReqErrChan <- err
		case "markdown":
			notifyBotMarkdownMessage := NotifyBotMarkdownMessage{}
			notifyQiweiReq, err := notifyBotMarkdownMessage.FormatNotifyBotParams(q)
			notifyQiweiReqChan <- notifyQiweiReq
			notifyQiweiReqErrChan <- err
		case "image":
			notifyBotImageMessage := NotifyBotImageMessage{}
			notifyQiweiReq, _ := notifyBotImageMessage.FormatNotifyBotParams(q)
			notifyQiweiReqChan <- notifyQiweiReq
			notifyQiweiReqErrChan <- err
		case "news":
			notifyBotNewsMessage := NotifyBotNewsMessage{}
			notifyQiweiReq, err := notifyBotNewsMessage.FormatNotifyBotParams(q)
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
		return false, errors.New("notifyQiweiReq is nil")
	}

	/**
	 * @step
	 * @发送企微消息
	 **/
	err = q.DoNotify(notifyQiweiReq)
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
 * @description: NotifyBotTextMessage FormatNotifyBotParams
 * @param {*QiweiBotNotice} qiweiBotNotice
 * @author: Jerry.Yang
 * @date: 2022-09-22 11:50:18
 * @return {*}
 */
func (n *NotifyBotTextMessage) FormatNotifyBotParams(qiweiBotNotice *QiweiBotNotice) ([]byte, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := n.CheckParams(qiweiBotNotice)
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @text的结构体
	 **/
	type ParamsTextStruct struct {
		Content             string   `json:"content"`
		MentionedList       []string `json:"mentioned_list"`
		MentionedMobileList []string `json:"mentioned_mobile_list"`
	}

	/**
	 * @step
	 * @构建发送的结构体
	 **/
	type ParamsStruct struct {
		Msgtype string           `json:"msgtype"`
		Text    ParamsTextStruct `json:"text"`
	}

	/**
	 * @step
	 * @复制参数
	 **/
	jsonParams, err := json.Marshal(ParamsStruct{
		Msgtype: "text",
		Text: ParamsTextStruct{
			Content:             qiweiBotNotice.SendMsg,
			MentionedList:       qiweiBotNotice.MentionedList,
			MentionedMobileList: qiweiBotNotice.MentionedMobileList,
		},
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
 * @description: NotifyBotMarkdownMessage FormatNotifyBotParams
 * @param {*QiweiBotNotice} qiweiBotNotice
 * @author: Jerry.Yang
 * @date: 2022-09-22 11:50:39
 * @return {*}
 */
func (n *NotifyBotMarkdownMessage) FormatNotifyBotParams(qiweiBotNotice *QiweiBotNotice) ([]byte, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := n.CheckParams(qiweiBotNotice)
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @text的结构体
	 **/
	type ParamsMarkdownStruct struct {
		Content             string   `json:"content"`
		MentionedList       []string `json:"mentioned_list"`
		MentionedMobileList []string `json:"mentioned_mobile_list"`
	}

	/**
	 * @step
	 * @构建发送的结构体
	 **/
	type ParamsStruct struct {
		Msgtype  string               `json:"msgtype"`
		Markdown ParamsMarkdownStruct `json:"markdown"`
	}

	/**
	 * @step
	 * @复制参数
	 **/
	jsonParams, err := json.Marshal(ParamsStruct{
		Msgtype: "markdown",
		Markdown: ParamsMarkdownStruct{
			Content:             qiweiBotNotice.SendMsg,
			MentionedList:       qiweiBotNotice.MentionedList,
			MentionedMobileList: qiweiBotNotice.MentionedMobileList,
		},
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
 * @description: NotifyBotImageMessage FormatNotifyBotParams
 * @param {*QiweiBotNotice} qiweiBotNotice
 * @author: Jerry.Yang
 * @date: 2022-09-22 11:50:52
 * @return {*}
 */
func (n *NotifyBotImageMessage) FormatNotifyBotParams(qiweiBotNotice *QiweiBotNotice) ([]byte, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := n.CheckParams(qiweiBotNotice)
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @先对mediaData解码
	 **/
	decodeMediaData, err := base64.StdEncoding.DecodeString(qiweiBotNotice.ImageData)
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @获取imageData的md5
	 **/
	imageDataMd5Obj := md5.New()
	imageDataMd5Obj.Write(decodeMediaData)
	imageDataMd5Str := hex.EncodeToString(imageDataMd5Obj.Sum(nil))

	/**
	 * @step
	 * @text的结构体
	 **/
	type ParamsImageStruct struct {
		Base64 string `json:"base64"`
		Md5    string `json:"md5"`
	}

	/**
	 * @step
	 * @构建发送的结构体
	 **/
	type ParamsStruct struct {
		Msgtype string            `json:"msgtype"`
		Image   ParamsImageStruct `json:"image"`
	}

	/**
	 * @step
	 * @复制参数
	 **/
	jsonParams, err := json.Marshal(ParamsStruct{
		Msgtype: "image",
		Image: ParamsImageStruct{
			Base64: qiweiBotNotice.ImageData,
			Md5:    imageDataMd5Str,
		},
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
 * @description: NotifyBotNewsMessage FormatNotifyBotParams
 * @param {*QiweiBotNotice} qiweiBotNotice
 * @author: Jerry.Yang
 * @date: 2022-09-22 11:51:05
 * @return {*}
 */
func (n *NotifyBotNewsMessage) FormatNotifyBotParams(qiweiBotNotice *QiweiBotNotice) ([]byte, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := n.CheckParams(qiweiBotNotice)
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @article的结构体
	 **/
	type ParamsArticlesStruct struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		PicUrl      string `json:"picurl"`
	}

	type ParamsNewsStruct struct {
		News []ParamsArticlesStruct `json:"articles"`
	}

	/**
	 * @step
	 * @构建发送的结构体
	 **/
	type ParamsStruct struct {
		Msgtype string           `json:"msgtype"`
		News    ParamsNewsStruct `json:"news"`
	}

	/**
	 * @step
	 * @append article
	 **/
	paramsArticles := make([]ParamsArticlesStruct, 0)
	paramsArticles = append(paramsArticles, ParamsArticlesStruct{
		Title:       qiweiBotNotice.Title,
		Description: qiweiBotNotice.Description,
		Url:         qiweiBotNotice.Url,
		PicUrl:      qiweiBotNotice.PicUrl,
	})

	/**
	 * @step
	 * @复制参数
	 **/
	jsonParams, err := json.Marshal(ParamsStruct{
		Msgtype: "news",
		News:    ParamsNewsStruct{paramsArticles},
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
 * @param {[]byte} QiweiBotNoticeReq
 * @author: Jerry.Yang
 * @date: 2022-09-22 11:57:00
 * @return {*}
 */
func (q *QiweiBotNotice) DoNotify(QiweiBotNoticeReq []byte) error {

	/**
	 * @step
	 * @构建返回值
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
		Url:     q.BotUrl,
		Body:    bytes.NewBuffer(QiweiBotNoticeReq),
		Options: httpOptions,
		Output:  resp,
	}
	mytoolHttpPkg.HttpRequest(&httpClient)

	/**
	 * @step
	 * @判断结果
	 **/
	if resp.ErrCode != 0 {
		return errors.New(resp.ErrMsg)
	}
	return nil
}

/**
 * @description: CheckParams
 * @author: Jerry.Yang
 * @date: 2022-09-22 11:37:09
 * @return {*}
 */
func (q *QiweiBotNotice) CheckParams() error {
	if q.MsgType == "" {
		return errors.New("QiweiBotNotice Err : msgType is empty!")
	}
	return nil
}

/**
 * @description: CheckParams
 * @param {*QiweiBotNotice} qiweiBotNotice
 * @author: Jerry.Yang
 * @date: 2022-09-22 11:40:19
 * @return {*}
 */
func (n *NotifyBotTextMessage) CheckParams(qiweiBotNotice *QiweiBotNotice) error {
	if qiweiBotNotice.SendMsg == "" {
		return errors.New("NotifyBotTextMessage Err : sendMsg is empty!")
	}
	return nil
}

/**
 * @description: NotifyBotMarkdownMessage CheckParams
 * @param {*QiweiBotNotice} qiweiBotNotice
 * @author: Jerry.Yang
 * @date: 2022-09-22 11:42:12
 * @return {*}
 */
func (n *NotifyBotMarkdownMessage) CheckParams(qiweiBotNotice *QiweiBotNotice) error {
	if qiweiBotNotice.SendMsg == "" {
		return errors.New("NotifyBotMarkdownMessage Err : sendMsg is empty!")
	}
	return nil
}

/**
 * @description: NotifyBotImageMessage CheckParams
 * @param {*QiweiBotNotice} qiweiBotNotice
 * @author: Jerry.Yang
 * @date: 2022-09-22 11:44:35
 * @return {*}
 */
func (n *NotifyBotImageMessage) CheckParams(qiweiBotNotice *QiweiBotNotice) error {
	if qiweiBotNotice.ImageData == "" {
		return errors.New("NotifyBotImageMessage Err : ImageData is empty!")
	}
	return nil
}

/**
 * @description: NotifyBotNewsMessage CheckParams
 * @param {*QiweiBotNotice} qiweiBotNotice
 * @author: Jerry.Yang
 * @date: 2022-09-22 11:48:06
 * @return {*}
 */
func (n *NotifyBotNewsMessage) CheckParams(qiweiBotNotice *QiweiBotNotice) error {
	if qiweiBotNotice.Title == "" {
		return errors.New("NotifyBotNewsMessage Err : Title is empty!")
	}

	if qiweiBotNotice.Description == "" {
		return errors.New("NotifyBotNewsMessage Err : Description is empty!")
	}

	if qiweiBotNotice.Url == "" {
		return errors.New("NotifyBotNewsMessage Err : Url is empty!")
	}

	if qiweiBotNotice.PicUrl == "" {
		return errors.New("NotifyBotNewsMessage Err : PicUrl is empty!")
	}
	return nil
}
