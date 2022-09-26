/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-19 17:28:57
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 11:53:14
 * @Description: createApp
 */
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type (
	CreateApp struct {
		ProjectName string
		AppName     string
		Method      string
	}

	CreateAppInputVO struct {
		CreateApp
		DirName  string
		FileName string
	}

	CreateAppOutputVO struct {
		CreateApp
		DirName  string
		FileName string
	}

	CreateAppRoute struct {
		CreateApp
		DirName  string
		FileName string
	}

	CreateAppService struct {
		CreateApp
		DirName  string
		FileName string
	}
)

/**
 * @description: Create
 * @param {string} c.ProjectName
 * @param {string} c.AppName
 * @param {string} c.Method
 * @author: Jerry.Yang
 * @date: 2022-09-19 17:31:25
 * @return {*}
 */
func (c *CreateApp) CreateContent() (string, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := c.CheckParams()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @获取当前目录
	 **/
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取目录错误!请重试!")
		fmt.Print("\r\n")
		return "", err
	}

	/**
	 * @step
	 * @判断目录是否存在
	 **/
	needCreateAppDir := fmt.Sprintf("%s/%s", dir, c.AppName)
	_, err = os.Stat(needCreateAppDir)
	if os.IsNotExist(err) {

		/**
		 * @step
		 * @创建目录
		 **/
		err = os.Mkdir(needCreateAppDir, 0777)
		if err != nil {
			fmt.Printf("创建App目录失败!请重试!")
			fmt.Print("\r\n")
			return "", err
		}
	}

	/**
	 * @step
	 * @获取fileContent inputVO
	 **/
	_, err = c.CreateAppInputVO(needCreateAppDir, "InputVO")
	if err != nil {
		fmt.Printf("CreateApp Err : %v", err.Error())
		fmt.Print("\r\n")
	}

	/**
	 * @step
	 * @获取fileContent OutputVO
	 **/
	_, err = c.CreateAppOutputVO(needCreateAppDir, "OutputVO")
	if err != nil {
		fmt.Printf("CreateApp Err : %v", err.Error())
		fmt.Print("\r\n")
	}

	/**
	 * @step
	 * @获取fileContent route
	 **/
	_, err = c.CreateAppRoute(needCreateAppDir, "Route")
	if err != nil {
		fmt.Printf("CreateApp Err : %v", err.Error())
		fmt.Print("\r\n")
	}

	/**
	 * @step
	 * @获取fileContent service
	 **/
	_, err = c.CreateAppService(needCreateAppDir, "Service")
	if err != nil {
		fmt.Printf("CreateApp Err : %v", err.Error())
		fmt.Print("\r\n")
	}
	return "", nil
}

/**
 * @description: CreateAppInputVO
 * @param {string} dirName
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-09-23 23:52:57
 * @return {*}
 */
func (c *CreateApp) CreateAppInputVO(dirName string, fileName string) (string, error) {

	/**
	 * @step
	 * @获取fileContent inputVO
	 **/
	createAppInputVO := CreateAppInputVO{CreateApp: CreateApp{ProjectName: c.ProjectName, AppName: c.AppName, Method: c.Method}, DirName: fmt.Sprintf("%s/%s%s.go", dirName, c.AppName, fileName), FileName: fileName}
	content, err := createAppInputVO.CreateContent()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @创建文件
	 **/
	createBase := CreateBase{DirName: createAppInputVO.DirName, FileName: createAppInputVO.FileName, FileContent: content}
	return createBase.CreateContent()
}

/**
 * @description: CreateAppOutputVO
 * @param {string} dirName
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-09-23 23:53:06
 * @return {*}
 */
func (c *CreateApp) CreateAppOutputVO(dirName string, fileName string) (string, error) {

	/**
	 * @step
	 * @获取fileContent inputVO
	 **/
	createAppOutputVO := CreateAppOutputVO{CreateApp: CreateApp{ProjectName: c.ProjectName, AppName: c.AppName, Method: c.Method}, DirName: fmt.Sprintf("%s/%s%s.go", dirName, c.AppName, fileName), FileName: fileName}
	content, err := createAppOutputVO.CreateContent()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @创建文件
	 **/
	createBase := CreateBase{DirName: createAppOutputVO.DirName, FileName: createAppOutputVO.FileName, FileContent: content}
	return createBase.CreateContent()
}

/**
 * @description: CreateAppRoute
 * @param {string} dirName
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-09-23 23:53:15
 * @return {*}
 */
func (c *CreateApp) CreateAppRoute(dirName string, fileName string) (string, error) {
	/**
	 * @step
	 * @获取fileContent inputVO
	 **/
	createAppRoute := CreateAppRoute{CreateApp: CreateApp{ProjectName: c.ProjectName, AppName: c.AppName, Method: c.Method}, DirName: fmt.Sprintf("%s/%s%s.go", dirName, c.AppName, fileName), FileName: fileName}
	content, err := createAppRoute.CreateContent()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @创建文件
	 **/
	createBase := CreateBase{DirName: createAppRoute.DirName, FileName: createAppRoute.FileName, FileContent: content}
	return createBase.CreateContent()
}

/**
 * @description: CreateAppService
 * @param {string} dirName
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-09-23 23:53:23
 * @return {*}
 */
func (c *CreateApp) CreateAppService(dirName string, fileName string) (string, error) {
	/**
	 * @step
	 * @获取fileContent inputVO
	 **/
	createAppService := CreateAppService{CreateApp: CreateApp{ProjectName: c.ProjectName, AppName: c.AppName, Method: c.Method}, DirName: fmt.Sprintf("%s/%s%s.go", dirName, c.AppName, fileName), FileName: fileName}
	content, err := createAppService.CreateContent()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @创建文件
	 **/
	createBase := CreateBase{DirName: createAppService.DirName, FileName: createAppService.FileName, FileContent: content}
	return createBase.CreateContent()
}

/**
 * @description: createContent
 * @param {string} appDir
 * @param {string} c.AppName
 * @param {string} c.FileName
 * @author: Jerry.Yang
 * @date: 2022-09-07 14:52:41
 * @return {*}
 */
func (c *CreateAppInputVO) CreateContent() (string, error) {

	/**
	 * @step
	 * @定义写入内容
	 **/
	fileContent := fmt.Sprintf("package %s", c.AppName)
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @写入interface
	 **/
	interfaceContent := fmt.Sprintf("type %s%sInterface interface {}", strings.Title(c.AppName), strings.Title(c.FileName))
	fileContent += interfaceContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @写入struct
	 **/
	structContent := fmt.Sprintf("type %s%s struct{}", strings.Title(c.AppName), strings.Title(c.FileName))
	fileContent += structContent
	return fileContent, nil
}

/**
 * @description: createContent
 * @param {string} appDir
 * @param {string} c.ProjectName
 * @param {string} c.AppName
 * @param {string} c.FileName
 * @author: Jerry.Yang
 * @date: 2022-09-07 16:04:00
 * @return {*}
 */
func (c *CreateAppOutputVO) CreateContent() (string, error) {
	fileContent := fmt.Sprintf("package %s", c.AppName)
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @import
	 **/
	importContent := fmt.Sprintf("import (\r\n")
	importContent += "	\"github.com/gin-gonic/gin\"\r\n"
	importContent += fmt.Sprintf("	\"%s/internal\"\r\n", c.ProjectName)
	importContent += ")\r\n"
	importContent += "\r\n"
	fileContent += importContent

	/**
	 * @step
	 * @interface
	 **/
	interfaceContent := fmt.Sprintf("type %s%sInterface interface {}", strings.Title(c.AppName), strings.Title(c.FileName))
	fileContent += interfaceContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @structContent
	 **/
	structContent := fmt.Sprintf("type %s%s struct{\r\n	HttpStatus int\r\n RetCode int\r\n RetMsg string\r\n RetResult interface{}\r\n}", strings.Title(c.AppName), strings.Title(c.FileName))
	fileContent += structContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @RenderOutputVOSimple
	 **/
	renderOutputVOSimpleContent := fmt.Sprintf("func RenderOutputVOSimple(ctx *gin.Context,retCode int,retMsg string,httpStatus ...int) error {\r\n")
	renderOutputVOSimpleContent += fmt.Sprintf("%s%s := %s%s{ \r\n", strings.Title(c.AppName), strings.Title(c.FileName), strings.Title(c.AppName), strings.Title(c.FileName))
	renderOutputVOSimpleContent += "RetCode : retCode,\r\n"
	renderOutputVOSimpleContent += "RetMsg : retMsg,\r\n"
	renderOutputVOSimpleContent += "} \r\n"
	renderOutputVOSimpleContent += "\r\n"
	renderOutputVOSimpleContent += "if len(httpStatus) == 0 {\r\n"
	renderOutputVOSimpleContent += fmt.Sprintf("%s%s.RenderOutputVO(ctx)\r\n", strings.Title(c.AppName), strings.Title(c.FileName))
	renderOutputVOSimpleContent += "return nil\r\n"
	renderOutputVOSimpleContent += "}\r\n"
	renderOutputVOSimpleContent += "\r\n"
	renderOutputVOSimpleContent += fmt.Sprintf("%s%s.HttpStatus = httpStatus[0]\r\n", strings.Title(c.AppName), strings.Title(c.FileName))
	renderOutputVOSimpleContent += fmt.Sprintf("%s%s.RenderOutputVO(ctx)\r\n", strings.Title(c.AppName), strings.Title(c.FileName))
	renderOutputVOSimpleContent += "return nil\r\n"
	renderOutputVOSimpleContent += "}\r\n"
	fileContent += renderOutputVOSimpleContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @RenderOutputVO
	 **/
	renderOutputVOContent := fmt.Sprintf("func (%s%s *%s%s) RenderOutputVO(ctx *gin.Context) error {\r\n", c.AppName, c.FileName, strings.Title(c.AppName), strings.Title(c.FileName))
	renderOutputVOContent += fmt.Sprintf("internalOutput := internal.Output{}\r\n")
	renderOutputVOContent += fmt.Sprintf("internalOutput.OutputFunc(ctx, %s%s.RetCode, %s%s.RetMsg, %s%s.RetResult, %s%s.HttpStatus)\r\n", c.AppName, c.FileName, c.AppName, c.FileName, c.AppName, c.FileName, c.AppName, c.FileName)
	renderOutputVOContent += "return nil\r\n"
	renderOutputVOContent += "}\r\n"
	fileContent += renderOutputVOContent
	return fileContent, nil
}

/**
 * @description: createContent
 * @param {string} appDir
 * @param {string} c.ProjectName
 * @param {string} c.AppName
 * @param {string} c.FileName
 * @param {string} c.Method
 * @author: Jerry.Yang
 * @date: 2022-09-07 16:42:13
 * @return {*}
 */
func (c *CreateAppRoute) CreateContent() (string, error) {
	fileContent := fmt.Sprintf("package %s", c.AppName)
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @import
	 **/
	importContent := fmt.Sprintf("import (\r\n")
	importContent += "\"github.com/gin-gonic/gin\"\r\n"
	importContent += "\"net/http\"\r\n"
	importContent += fmt.Sprintf("\"%s/config\"\r\n", c.ProjectName)
	importContent += ")\r\n"
	importContent += "\r\n"
	fileContent += importContent

	/**
	 * @step
	 * @interface
	 **/
	interfaceContent := fmt.Sprintf("type %s%sInterface interface {}", strings.Title(c.AppName), strings.Title(c.FileName))
	fileContent += interfaceContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @structContent
	 **/
	structContent := fmt.Sprintf("type %s%s struct{}", strings.Title(c.AppName), strings.Title(c.FileName))
	fileContent += structContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @routeFunc
	 **/
	routeFuncContent := fmt.Sprintf("func %sRouteFunc(ctx *gin.Context) {\r\n", strings.Title(c.AppName))
	routeFuncContent += fmt.Sprintf("%sInputVO := %sInputVO{} \r\n", c.AppName, strings.Title(c.AppName))

	/**
	 * @step
	 * @判断什么请求方式
	 **/
	if c.Method == "GET" {
		routeFuncContent += fmt.Sprintf("if err := ctx.ShouldBindQuery(&%sInputVO); err != nil {\r\n", c.AppName)
	}

	if c.Method == "POST" {
		routeFuncContent += fmt.Sprintf("if err := ctx.ShouldBind(%sInputVO); err != nil {\r\n", c.AppName)
	}
	routeFuncContent += fmt.Sprintf("RenderOutputVOSimple(ctx, config.COMMON_ERROR, err.Error())\r\n")
	routeFuncContent += "return \r\n"
	routeFuncContent += "}\r\n"
	routeFuncContent += "\r\n"

	routeFuncContent += fmt.Sprintf("err := %sInputVO.%sServiceFunc(ctx)\r\n", c.AppName, strings.Title(c.AppName))
	routeFuncContent += "if err != nil {\r\n"
	routeFuncContent += "return"
	routeFuncContent += "}\r\n"
	routeFuncContent += "\r\n"

	routeFuncContent += fmt.Sprintf("%sOutputVO := %sOutputVO{} \r\n", c.AppName, strings.Title(c.AppName))
	routeFuncContent += fmt.Sprintf("%sOutputVO.HttpStatus = http.StatusOK \r\n", c.AppName)
	routeFuncContent += fmt.Sprintf("%sOutputVO.RetCode = config.NO_ERROR \r\n", c.AppName)
	routeFuncContent += fmt.Sprintf("%sOutputVO.RetResult = true \r\n", c.AppName)
	routeFuncContent += fmt.Sprintf("%sOutputVO.RenderOutputVO(ctx) \r\n", c.AppName)
	routeFuncContent += "return \r\n"
	routeFuncContent += "}\r\n"
	fileContent += routeFuncContent
	return fileContent, nil
}

/**
 * @description: createContent
 * @param {string} appDir
 * @param {string} c.ProjectName
 * @param {string} c.AppName
 * @param {string} c.FileName
 * @param {string} c.Method
 * @author: Jerry.Yang
 * @date: 2022-09-07 17:16:40
 * @return {*}
 */
func (c *CreateAppService) CreateContent() (string, error) {
	fileContent := fmt.Sprintf("package %s", c.AppName)
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @import
	 **/
	importContent := fmt.Sprintf("import (\r\n")
	importContent += "\"github.com/gin-gonic/gin\"\r\n"
	importContent += ")\r\n"
	importContent += "\r\n"
	fileContent += importContent

	/**
	 * @step
	 * @interface
	 **/
	interfaceContent := fmt.Sprintf("type %s%sInterface interface {}", strings.Title(c.AppName), strings.Title(c.FileName))
	fileContent += interfaceContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @structContent
	 **/
	structContent := fmt.Sprintf("type %s%s struct{}", strings.Title(c.AppName), strings.Title(c.FileName))
	fileContent += structContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	serviceContent := fmt.Sprintf("func (vo *%sInputVO) %sServiceFunc(ctx *gin.Context) error {\r\n", strings.Title(c.AppName), strings.Title(c.AppName))
	serviceContent += "return nil \r\n"
	serviceContent += "}\r\n"
	fileContent += serviceContent
	return fileContent, nil
}

/**
 * @description: checkParams
 * @param {string} c.ProjectName
 * @param {string} c.AppName
 * @param {string} c.Method
 * @author: Jerry.Yang
 * @date: 2022-09-19 17:33:25
 * @return {*}
 */
func (c *CreateApp) CheckParams() error {

	/**
	 * @step
	 * @判断参数
	 **/
	if c.ProjectName == "" {
		return errors.New("c.ProjectName缺失!")
	}

	if c.AppName == "" {
		return errors.New("c.AppName确实!")
	}

	if c.Method == "" {
		return errors.New("请求方式缺失!")
	}
	return nil
}
