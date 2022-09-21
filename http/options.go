/*
 * @Author: yangjie04@qutoutiao.net
 * @Date: 2022-09-21 16:08:24
 * @LastEditors: yangjie04@qutoutiao.net
 * @LastEditTime: 2022-09-21 16:57:41
 * @Description: options
 */
package http

type (
	HttpOptionsInterface interface{}

	HttpOptions struct{}

	/**
	 * @step
	 * @定义optionVal
	 **/
	HttpOption struct {
		Value interface{}
	}

	/**
	 * @step
	 * @定义options
	 **/
	HttpOptionFunc func(map[string]HttpOption) error
)

/**
 * @定义常量
 * @author Jerry.Yang
 * @date 2022-09-21 15:57:57
 **/
const OPTION_OUT_TIME = "outTime"
const OPTION_HEADERS = "headers"

/**
 * @description: SetHeaders
 * @param {map[string]string} value
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-21 16:32:35
 * @return {*}
 */
func (h *HttpOptions) SetHeaders(value map[string]string) HttpOptionFunc {
	return func(HttpOptionFunc map[string]HttpOption) error {
		HttpOptionFunc[OPTION_HEADERS] = HttpOption{value}
		return nil
	}
}

/**
 * @description: setTimeOut
 * @param {int} value
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-03-08 11:17:32
 * @return {*}
 */
func (h *HttpOptions) SetOutTime(value int) HttpOptionFunc {
	return func(HttpOptionFunc map[string]HttpOption) error {
		HttpOptionFunc[OPTION_OUT_TIME] = HttpOption{value}
		return nil
	}
}

/**
 * @description: GetHeaders
 * @param {map[string]HttpOption} httpOptions
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-21 16:59:40
 * @return {*}
 */
func (h *HttpOptions) GetHeaders(httpOptions map[string]HttpOption) map[string]string {

	/**
	 * @step
	 * @定义
	 **/
	headers := map[string]string{}

	/**
	 * @step
	 * @获取设置的headers
	 **/
	setHeaders, ok := httpOptions[OPTION_HEADERS]
	if ok {
		return setHeaders.Value.(map[string]string)
	}
	return headers
}

/**
 * @description: GetOutTime
 * @param {map[string]HttpOption} httpOptions
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-21 16:49:46
 * @return {*}
 */
func (h *HttpOptions) GetOutTime(httpOptions map[string]HttpOption) int {

	/**
	 * @step
	 * @定义初始值
	 **/
	outTime := 60

	/**
	 * @step
	 * @获取设置的outTime
	 **/
	setOutTime, ok := httpOptions[OPTION_OUT_TIME]
	if !ok {
		return outTime
	}

	/**
	 * @step
	 * @获取设置的过期时间的真实的数据
	 **/
	setOutTimeInt := setOutTime.Value.(int)

	/**
	 * @step
	 * @判断设置的过期时间是否有效
	 **/
	if setOutTimeInt <= 0 {
		return outTime
	}
	return setOutTimeInt
}

/**
 * @description: SetOptions
 * @param {[]Option} options
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-03-01 15:44:58
 * @return {*}
 */
func (h *HttpOptions) SetOptions(httpOptionFuncs []HttpOptionFunc) map[string]HttpOption {

	/**
	 * @step
	 * @定义
	 **/
	setHttpOption := map[string]HttpOption{}

	/**
	 * @step
	 * @获取set的functions
	 **/
	for _, option := range httpOptionFuncs {
		if option != nil {
			err := option(setHttpOption)
			if err != nil {
				return nil
			}
		}
	}
	return setHttpOption
}
