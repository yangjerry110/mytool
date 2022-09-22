/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-19 17:28:57
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-22 15:55:33
 * @Description: createApp
 */
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type CreateAppInterface interface {
	CreateApp(projectName string, appName string, method string)
	createContent(appDir string, projectName string, appName string, fileName string, method string) string
	checkParams(projectName string, appName string, method string) error
}

type (
	CreateApp struct{}

	CreateAppInputVO struct {
		DirName  string
		FileName string
	}

	CreateAppOutputVO struct {
		DirName  string
		FileName string
	}

	CreateAppRoute struct {
		DirName  string
		FileName string
	}

	CreateAppService struct {
		DirName  string
		FileName string
	}
)

/**
 * @description: CreateApp
 * @param {string} projectName
 * @param {string} appName
 * @param {string} method
 * @author: Jerry.Yang
 * @date: 2022-09-19 17:31:25
 * @return {*}
 */
func (c *CreateApp) CreateApp(projectName string, appName string, method string) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := c.checkParams(projectName, appName, method)
	if err != nil {
		fmt.Printf("CreateApp Err : %v", err.Error())
		fmt.Print("\r\n")
		return
	}

	/**
	 * @step
	 * @获取当前目录
	 **/
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取目录错误!请重试!")
		fmt.Print("\r\n")
		return
	}

	/**
	 * @step
	 * @判断目录是否存在
	 **/
	needCreateAppDir := fmt.Sprintf("%s/%s", dir, appName)
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
			return
		}
	}

	/**
	 * @step
	 * @获取base
	 **/
	createBaseObj := CreateBase{}

	/**
	 * @step
	 * @获取fileContent inputVO
	 **/
	createAppInputVO := CreateAppInputVO{FileName: "InputVO", DirName: fmt.Sprintf("%s/%s%s.go", needCreateAppDir, appName, "InputVO")}
	createAppInputVOFileContent := createAppInputVO.createContent(needCreateAppDir, appName, createAppInputVO.FileName)
	err = createBaseObj.CreateFile(createAppInputVO.DirName, createAppInputVO.FileName, createAppInputVOFileContent)
	if err != nil {
		fmt.Printf("CreateApp Err : %v", err.Error())
		fmt.Print("\r\n")
	}

	/**
	 * @step
	 * @获取fileContent OutputVO
	 **/
	createAppOutputVO := CreateAppOutputVO{FileName: "OutputVO", DirName: fmt.Sprintf("%s/%s%s.go", needCreateAppDir, appName, "OutputVO")}
	createAppOutputVOFileContent := createAppOutputVO.createContent(needCreateAppDir, projectName, appName, createAppOutputVO.FileName)
	err = createBaseObj.CreateFile(createAppInputVO.DirName, createAppOutputVO.FileName, createAppOutputVOFileContent)
	if err != nil {
		fmt.Printf("CreateApp Err : %v", err.Error())
		fmt.Print("\r\n")
	}

	/**
	 * @step
	 * @获取fileContent route
	 **/
	createAppRoute := CreateAppRoute{FileName: "Route", DirName: fmt.Sprintf("%s/%s%s.go", needCreateAppDir, appName, "Route")}
	createAppRouteFileContent := createAppRoute.createContent(needCreateAppDir, projectName, appName, createAppRoute.FileName, method)
	err = createBaseObj.CreateFile(createAppInputVO.DirName, createAppRoute.FileName, createAppRouteFileContent)
	if err != nil {
		fmt.Printf("CreateApp Err : %v", err.Error())
		fmt.Print("\r\n")
	}

	/**
	 * @step
	 * @获取fileContent service
	 **/
	createAppService := CreateAppService{FileName: "Service", DirName: fmt.Sprintf("%s/%s%s.go", needCreateAppDir, appName, "Service")}
	CreateAppServiceFileContent := createAppService.createContent(needCreateAppDir, projectName, appName, createAppService.FileName, method)
	err = createBaseObj.CreateFile(createAppInputVO.DirName, createAppService.FileName, CreateAppServiceFileContent)
	if err != nil {
		fmt.Printf("CreateApp Err : %v", err.Error())
		fmt.Print("\r\n")
	}

	return
}

/**
 * @description: createContent
 * @param {string} appDir
 * @param {string} appName
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-09-07 14:52:41
 * @return {*}
 */
func (c *CreateAppInputVO) createContent(appDir string, appName string, fileName string) string {

	/**
	 * @step
	 * @定义写入内容
	 **/
	fileContent := fmt.Sprintf("package %s", appName)
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @写入interface
	 **/
	interfaceContent := fmt.Sprintf("type %s%sInterface interface {}", strings.Title(appName), strings.Title(fileName))
	fileContent += interfaceContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @写入struct
	 **/
	structContent := fmt.Sprintf("type %s%s struct{}", strings.Title(appName), strings.Title(fileName))
	fileContent += structContent
	return fileContent
}

/**
 * @description: createContent
 * @param {string} appDir
 * @param {string} projectName
 * @param {string} appName
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-09-07 16:04:00
 * @return {*}
 */
func (c *CreateAppOutputVO) createContent(appDir string, projectName string, appName string, fileName string) string {
	fileContent := fmt.Sprintf("package %s", appName)
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @import
	 **/
	importContent := fmt.Sprintf("import (\r\n")
	importContent += "	\"github.com/gin-gonic/gin\"\r\n"
	importContent += fmt.Sprintf("	\"%s/internal\"\r\n", projectName)
	importContent += ")\r\n"
	importContent += "\r\n"
	fileContent += importContent

	/**
	 * @step
	 * @interface
	 **/
	interfaceContent := fmt.Sprintf("type %s%sInterface interface {}", strings.Title(appName), strings.Title(fileName))
	fileContent += interfaceContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @structContent
	 **/
	structContent := fmt.Sprintf("type %s%s struct{\r\n	HttpStatus int\r\n RetCode int\r\n RetMsg string\r\n RetResult interface{}\r\n}", strings.Title(appName), strings.Title(fileName))
	fileContent += structContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @RenderOutputVOSimple
	 **/
	renderOutputVOSimpleContent := fmt.Sprintf("func RenderOutputVOSimple(ctx *gin.Context,retCode int,retMsg string,httpStatus ...int) error {\r\n")
	renderOutputVOSimpleContent += fmt.Sprintf("%s%s := %s%s{ \r\n", strings.Title(appName), strings.Title(fileName), strings.Title(appName), strings.Title(fileName))
	renderOutputVOSimpleContent += "RetCode : retCode,\r\n"
	renderOutputVOSimpleContent += "RetMsg : retMsg,\r\n"
	renderOutputVOSimpleContent += "} \r\n"
	renderOutputVOSimpleContent += "\r\n"
	renderOutputVOSimpleContent += "if len(httpStatus) == 0 {\r\n"
	renderOutputVOSimpleContent += fmt.Sprintf("%s%s.RenderOutputVO(ctx)\r\n", strings.Title(appName), strings.Title(fileName))
	renderOutputVOSimpleContent += "return nil\r\n"
	renderOutputVOSimpleContent += "}\r\n"
	renderOutputVOSimpleContent += "\r\n"
	renderOutputVOSimpleContent += fmt.Sprintf("%s%s.HttpStatus = httpStatus[0]\r\n", strings.Title(appName), strings.Title(fileName))
	renderOutputVOSimpleContent += fmt.Sprintf("%s%s.RenderOutputVO(ctx)\r\n", strings.Title(appName), strings.Title(fileName))
	renderOutputVOSimpleContent += "return nil\r\n"
	renderOutputVOSimpleContent += "}\r\n"
	fileContent += renderOutputVOSimpleContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @RenderOutputVO
	 **/
	renderOutputVOContent := fmt.Sprintf("func (%s%s *%s%s) RenderOutputVO(ctx *gin.Context) error {\r\n", appName, fileName, strings.Title(appName), strings.Title(fileName))
	renderOutputVOContent += fmt.Sprintf("internalOutput := internal.Output{}\r\n")
	renderOutputVOContent += fmt.Sprintf("internalOutput.OutputFunc(ctx, %s%s.RetCode, %s%s.RetMsg, %s%s.RetResult, %s%s.HttpStatus)\r\n", appName, fileName, appName, fileName, appName, fileName, appName, fileName)
	renderOutputVOContent += "return nil\r\n"
	renderOutputVOContent += "}\r\n"
	fileContent += renderOutputVOContent
	return fileContent
}

/**
 * @description: createContent
 * @param {string} appDir
 * @param {string} projectName
 * @param {string} appName
 * @param {string} fileName
 * @param {string} method
 * @author: Jerry.Yang
 * @date: 2022-09-07 16:42:13
 * @return {*}
 */
func (c *CreateAppRoute) createContent(appDir string, projectName string, appName string, fileName string, method string) string {
	fileContent := fmt.Sprintf("package %s", appName)
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @import
	 **/
	importContent := fmt.Sprintf("import (\r\n")
	importContent += "\"github.com/gin-gonic/gin\"\r\n"
	importContent += "\"net/http\"\r\n"
	importContent += fmt.Sprintf("\"%s/config\"\r\n", projectName)
	importContent += ")\r\n"
	importContent += "\r\n"
	fileContent += importContent

	/**
	 * @step
	 * @interface
	 **/
	interfaceContent := fmt.Sprintf("type %s%sInterface interface {}", strings.Title(appName), strings.Title(fileName))
	fileContent += interfaceContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @structContent
	 **/
	structContent := fmt.Sprintf("type %s%s struct{}", strings.Title(appName), strings.Title(fileName))
	fileContent += structContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @routeFunc
	 **/
	routeFuncContent := fmt.Sprintf("func %sRouteFunc(ctx *gin.Context) {\r\n", strings.Title(appName))
	routeFuncContent += fmt.Sprintf("%sInputVO := %sInputVO{} \r\n", appName, strings.Title(appName))

	/**
	 * @step
	 * @判断什么请求方式
	 **/
	if method == "GET" {
		routeFuncContent += fmt.Sprintf("if err := ctx.ShouldBindQuery(&%sInputVO); err != nil {\r\n", appName)
	}

	if method == "POST" {
		routeFuncContent += fmt.Sprintf("if err := ctx.ShouldBind(%sInputVO); err != nil {\r\n", appName)
	}
	routeFuncContent += fmt.Sprintf("RenderOutputVOSimple(ctx, config.COMMON_ERROR, err.Error())\r\n")
	routeFuncContent += "return \r\n"
	routeFuncContent += "}\r\n"
	routeFuncContent += "\r\n"

	routeFuncContent += fmt.Sprintf("err := %sInputVO.%sServiceFunc(ctx)\r\n", appName, strings.Title(appName))
	routeFuncContent += "if err != nil {\r\n"
	routeFuncContent += "return"
	routeFuncContent += "}\r\n"
	routeFuncContent += "\r\n"

	routeFuncContent += fmt.Sprintf("%sOutputVO := %sOutputVO{} \r\n", appName, strings.Title(appName))
	routeFuncContent += fmt.Sprintf("%sOutputVO.HttpStatus = http.StatusOK \r\n", appName)
	routeFuncContent += fmt.Sprintf("%sOutputVO.RetCode = config.NO_ERROR \r\n", appName)
	routeFuncContent += fmt.Sprintf("%sOutputVO.RetResult = true \r\n", appName)
	routeFuncContent += fmt.Sprintf("%sOutputVO.RenderOutputVO(ctx) \r\n", appName)
	routeFuncContent += "return \r\n"
	routeFuncContent += "}\r\n"
	fileContent += routeFuncContent
	return fileContent
}

/**
 * @description: createContent
 * @param {string} appDir
 * @param {string} projectName
 * @param {string} appName
 * @param {string} fileName
 * @param {string} method
 * @author: Jerry.Yang
 * @date: 2022-09-07 17:16:40
 * @return {*}
 */
func (c *CreateAppService) createContent(appDir string, projectName string, appName string, fileName string, method string) string {
	fileContent := fmt.Sprintf("package %s", appName)
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
	interfaceContent := fmt.Sprintf("type %s%sInterface interface {}", strings.Title(appName), strings.Title(fileName))
	fileContent += interfaceContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @structContent
	 **/
	structContent := fmt.Sprintf("type %s%s struct{}", strings.Title(appName), strings.Title(fileName))
	fileContent += structContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	serviceContent := fmt.Sprintf("func (vo *%sInputVO) %sServiceFunc(ctx *gin.Context) error {\r\n", strings.Title(appName), strings.Title(appName))
	serviceContent += "return nil \r\n"
	serviceContent += "}\r\n"
	fileContent += serviceContent
	return fileContent
}

/**
 * @description: checkParams
 * @param {string} projectName
 * @param {string} appName
 * @param {string} method
 * @author: Jerry.Yang
 * @date: 2022-09-19 17:33:25
 * @return {*}
 */
func (c *CreateApp) checkParams(projectName string, appName string, method string) error {

	/**
	 * @step
	 * @判断参数
	 **/
	if projectName == "" {
		return errors.New("projectName缺失!")
	}

	if appName == "" {
		return errors.New("appName确实!")
	}

	if method == "" {
		return errors.New("请求方式缺失!")
	}
	return nil
}
