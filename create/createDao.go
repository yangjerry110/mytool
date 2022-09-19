/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-19 14:37:33
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-19 17:41:32
 * @Description: createDao
 */
package create

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type CreateDaoInterface interface {
	CreatedDao(projectName string, daoName string, modelName string, authorName ...string)
	createContent(daoName string, modelName string, projectName string, authorName string) (string, error)
	createDaoInfo(daoName string, modelName string, authorName string) (string, error)
	createDaoList(daoName string, modelName string, authorName string) (string, error)
	createDaoSave(daoName string, modelName string, authorName string) (string, error)
	createDaoDeleted(daoName string, modelName string, authorName string) (string, error)
	checkParam(projectName string, daoName string, modelName string) error
}

type CreateDao struct{}

/**
 * @description: CreatedDao
 * @param {string} projectName
 * @param {string} daoName
 * @param {string} modelName
 * @param {...string} authorName
 * @author: Jerry.Yang
 * @date: 2022-09-19 16:50:28
 * @return {*}
 */
func (c *CreateDao) CreatedDao(projectName string, daoName string, modelName string, authorName ...string) {

	/**
	 * @step
	 * @定义参数
	 **/
	thisAuthorName := ""
	if len(authorName) == 0 {
		thisAuthorName = "Jerry.Yang"
	} else {
		thisAuthorName = authorName[0]
	}

	/**
	 * @step
	 * @检查参数
	 **/
	err := c.checkParam(projectName, daoName, modelName)
	if err != nil {
		fmt.Printf("CreatedDaoErr : %v", err.Error())
		return
	}

	/**
	 * @step
	 * @获取fileContent
	 **/
	fileContent, err := c.createContent(daoName, modelName, projectName, thisAuthorName)
	if err != nil {
		fmt.Printf("CreatedDaoErr : %v", err.Error())
		fmt.Print("\r\n")
		return
	}

	/**
	 * @step
	 * @获取当前目录
	 **/
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("CreatedDaoErr : 获取目录错误!请重试! ErrMsg : %v", err.Error())
		fmt.Print("\r\n")
		return
	}

	/**
	 * @step
	 * @创建文件
	 **/
	createBaseObj := CreateBase{}
	err = createBaseObj.CreateFile(dir, daoName, fileContent)
	if err != nil {
		fmt.Printf("CreatedDaoErr : %v", err.Error())
		fmt.Print("\r\n")
		return
	}

	/**
	 * @step
	 * @成功
	 **/
	fmt.Printf("CreateDao Success!")
	fmt.Print("\r\n")
	return
}

/**
 * @description: CreateContent
 * @param {string} daoName
 * @param {string} modelName
 * @param {string} authorName
 * @author: Jerry.Yang
 * @date: 2022-09-14 16:29:08
 * @return {*}
 */
func (c *CreateDao) createContent(daoName string, modelName string, projectName string, authorName string) (string, error) {
	fileContent := fmt.Sprintf("package dao")
	fileContent += "\r\n"

	/**
	 * @step
	 * @import
	 **/
	fileContent += "import ( \r\n"
	fileContent += fmt.Sprintf("\"%s/internal\"\r\n", projectName)
	fileContent += fmt.Sprintf("\"%s/model\"\r\n", projectName)
	fileContent += ")\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @interface
	 **/
	fileContent += fmt.Sprintf("type %sInterface interface {\r\n", strings.Title(daoName))
	fileContent += fmt.Sprintf("GetInfo(%s *model.%s) (*model.%s, error)\r\n", modelName, strings.Title(modelName), strings.Title(modelName))
	fileContent += fmt.Sprintf("GetList(%s *model.%s) ([]*model.%s, error) \r\n", modelName, strings.Title(modelName), strings.Title(modelName))
	fileContent += fmt.Sprintf("Save(%sModel *model.%s) (bool, error)\r\n", modelName, strings.Title(modelName))
	fileContent += fmt.Sprintf("Delete(%s *model.%s) (bool, error)\r\n", modelName, strings.Title(modelName))
	fileContent += "}\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @struct
	 **/
	fileContent += fmt.Sprintf("type %s struct {}\r\n", strings.Title(daoName))
	fileContent += "\r\n"

	/**
	 * @step
	 * @getInfoById
	 **/
	infoContent := c.createDaoInfo(daoName, modelName, authorName)
	fileContent += infoContent
	fileContent += "\r\n"

	/**
	 * @step
	 * @getList
	 **/
	listContent := c.createDaoList(daoName, modelName, authorName)
	fileContent += listContent
	fileContent += "\r\n"

	/**
	 * @step
	 * @save
	 **/
	saveContent := c.createDaoSave(daoName, modelName, authorName)
	fileContent += saveContent
	fileContent += "\r\n"

	/**
	 * @step
	 * @delete
	 **/
	deleteContent := c.createDaoDeleted(daoName, modelName, authorName)
	fileContent += deleteContent
	fileContent += "\r\n"

	return fileContent, nil
}

/**
 * @description: CreateDaoInfo
 * @param {string} daoName
 * @param {string} modelName
 * @param {string} authorName
 * @author: Jerry.Yang
 * @date: 2022-09-14 16:35:31
 * @return {*}
 */
func (c *CreateDao) createDaoInfo(daoName string, modelName string, authorName string) string {
	fileContent := fmt.Sprintf("/** \r\n * @description: GetInfo \r\n * @param {*model.%s} %s \r\n * @author: %s \r\n * @date: %s \r\n * @return {*} \r\n */ \r\n", strings.Title(modelName), modelName, authorName, time.Now().Format("2006-01-02 15:04:05"))
	fileContent += fmt.Sprintf("func (%s *%s) GetInfo(%s *model.%s) (*model.%s, error) {\r\n", daoName[0:1], strings.Title(daoName), modelName, strings.Title(modelName), strings.Title(modelName))
	fileContent += fmt.Sprintf("result := model.%s{}\r\n", strings.Title(modelName))
	fileContent += fmt.Sprintf("if err := internal.DbClient().Where(%s).First(&result).Error; err != nil { \r\n", modelName)
	fileContent += "internal.LoggorError(err)\r\n"
	fileContent += "return nil, err\r\n"
	fileContent += "}\r\n"
	fileContent += "return &result, nil\r\n"
	fileContent += "}\r\n"
	fileContent += "\r\n"
	return fileContent
}

/**
 * @description: CreateDaoList
 * @param {string} daoName
 * @param {string} modelName
 * @param {string} authorName
 * @author: Jerry.Yang
 * @date: 2022-09-15 14:37:55
 * @return {*}
 */
func (c *CreateDao) createDaoList(daoName string, modelName string, authorName string) string {
	fileContent := fmt.Sprintf("/** \r\n * @description: GetList \r\n * @param {*model.%s} %s \r\n * @author: %s \r\n * @date: %s \r\n * @return {*} \r\n */ \r\n", strings.Title(modelName), modelName, authorName, time.Now().Format("2006-01-02 15:04:05"))
	fileContent += fmt.Sprintf("func (%s *%s) GetList(%s *model.%s) ([]*model.%s, error) {\r\n", daoName[0:1], strings.Title(daoName), modelName, strings.Title(modelName), strings.Title(modelName))
	fileContent += fmt.Sprintf("result := make([]*model.%s, 0)\r\n", strings.Title(modelName))
	fileContent += fmt.Sprintf("if err := internal.DbClient().Where(%s).Find(&result).Error; err != nil {\r\n", modelName)
	fileContent += "internal.LoggorError(err)\r\n"
	fileContent += "return nil, err\r\n"
	fileContent += "}\r\n"
	fileContent += "return result, nil\r\n"
	fileContent += "}\r\n"
	fileContent += "\r\n"
	return fileContent
}

/**
 * @description: CreateDaoSave
 * @param {string} daoName
 * @param {string} modelName
 * @param {string} authorName
 * @author: Jerry.Yang
 * @date: 2022-09-15 14:38:46
 * @return {*}
 */
func (c *CreateDao) createDaoSave(daoName string, modelName string, authorName string) string {
	fileContent := fmt.Sprintf("/** \r\n * @description: Save \r\n * @param {*model.%s} %s \r\n * @author: %s \r\n * @date: %s \r\n * @return {*} \r\n */ \r\n", modelName, strings.Title(modelName), authorName, time.Now().Format("2006-01-02 15:04:05"))
	fileContent += fmt.Sprintf("func (%s *%s) Save(%s *model.%s) (bool, error) {\r\n", daoName[0:1], strings.Title(daoName), modelName, strings.Title(modelName))
	fileContent += "result := false\r\n"
	fileContent += fmt.Sprintf("if err := internal.DbClient().Save(%s).Error; err != nil {\r\n", modelName)
	fileContent += "internal.LoggorError(err)\r\n"
	fileContent += "return result, err\r\n"
	fileContent += "}\r\n"
	fileContent += "result = true\r\n"
	fileContent += "return result, nil\r\n"
	fileContent += "}\r\n"
	fileContent += "\r\n"
	return fileContent
}

/**
 * @description: CreateDaoDeleted
 * @param {string} daoName
 * @param {string} modelName
 * @param {string} authorName
 * @author: Jerry.Yang
 * @date: 2022-09-15 14:42:04
 * @return {*}
 */
func (c *CreateDao) createDaoDeleted(daoName string, modelName string, authorName string) string {
	fileContent := fmt.Sprintf("/** \r\n * @description: Delete \r\n * @param {*model.%s} %s \r\n * @author: %s \r\n * @date: %s \r\n * @return {*} \r\n */ \r\n", strings.Title(modelName), modelName, authorName, time.Now().Format("2006-01-02 15:04:05"))
	fileContent += fmt.Sprintf("func (%s *%s) Delete(%s *model.%s) (bool, error) {\r\n", daoName[0:1], strings.Title(daoName), modelName, strings.Title(modelName))
	fileContent += fmt.Sprintf("result := false\r\n")
	fileContent += fmt.Sprintf("if err := internal.DbClient().Model(&model.%s{}).Where(%s).Update(\"is_deleted\", model.IS_DELETED).Error; err != nil {\r\n", strings.Title(modelName), modelName)
	fileContent += "internal.LoggorError(err)\r\n"
	fileContent += "return result, err\r\n"
	fileContent += "}\r\n"
	fileContent += "result = true\r\n"
	fileContent += "return result, nil\r\n"
	fileContent += "}\r\n"
	fileContent += "\r\n"
	return fileContent
}

/**
 * @description: checkParam
 * @param {string} projectName
 * @param {string} daoName
 * @param {string} modelName
 * @author: Jerry.Yang
 * @date: 2022-09-19 16:51:39
 * @return {*}
 */
func (c *CreateDao) checkParam(projectName string, daoName string, modelName string) error {
	if projectName == "" {
		return errors.New("projectName 为空!")
	}

	if daoName == "" {
		return errors.New("daoName 为空!")
	}

	if modelName == "" {
		return errors.New("modelName 为空!")
	}
	return nil
}
