/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-19 14:37:33
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 15:06:21
 * @Description: createDao
 */
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)


/**
 * @description: CreatedDao
 * @param {string} projectName
 * @param {string} c.DaoName
 * @param {string} c.ModelName
 * @param {...string} c.AuthorName
 * @author: Jerry.Yang
 * @date: 2022-09-19 16:50:28
 * @return {*}
 */
func (c *CreateDao) CreateContent() (string, error) {
	/**
	 * @step
	 * @检查参数
	 **/
	err := c.CheckParams()
	if err != nil {
		fmt.Printf("CreatedDaoErr : %v", err.Error())
		return "", err
	}

	/**
	 * @step
	 * @获取fileContent
	 **/
	fileContent, err := c.GetContent()
	if err != nil {
		fmt.Printf("CreatedDaoErr : %v", err.Error())
		fmt.Print("\r\n")
		return "", err
	}

	/**
	 * @step
	 * @获取当前目录
	 **/
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("CreatedDaoErr : 获取目录错误!请重试! ErrMsg : %v", err.Error())
		fmt.Print("\r\n")
		return "", err
	}

	/**
	 * @step
	 * @创建文件
	 **/
	createBase := CreateBase{DirName: dir, FileName: c.DaoName, FileContent: fileContent}
	_, err = createBase.CreateContent()
	if err != nil {
		fmt.Printf("CreatedDaoErr : %v", err.Error())
		fmt.Print("\r\n")
		return "", err
	}

	/**
	 * @step
	 * @成功
	 **/
	fmt.Printf("CreateDao Success!")
	fmt.Print("\r\n")
	return "", nil
}

/**
 * @description: GetContent
 * @param {string} c.DaoName
 * @param {string} c.ModelName
 * @param {string} c.AuthorName
 * @author: Jerry.Yang
 * @date: 2022-09-14 16:29:08
 * @return {*}
 */
func (c *CreateDao) GetContent() (string, error) {

	thisAuthorName := ""
	if len(c.AuthorName) == 0 {
		thisAuthorName = "Jerry.Yang"
	} else {
		thisAuthorName = c.AuthorName[0]
	}

	fileContent := fmt.Sprintf("package dao")
	fileContent += "\r\n"

	/**
	 * @step
	 * @import
	 **/
	fileContent += "import ( \r\n"
	fileContent += fmt.Sprintf("\"%s/internal\"\r\n", c.ProjectName)
	fileContent += fmt.Sprintf("\"%s/c.ModelName\"\r\n", c.ProjectName)
	fileContent += ")\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @interface
	 **/
	fileContent += fmt.Sprintf("type %sInterface interface {\r\n", strings.Title(c.DaoName))
	fileContent += fmt.Sprintf("GetInfo(%s *c.ModelName.%s) (*c.ModelName.%s, error)\r\n", c.ModelName, strings.Title(c.ModelName), strings.Title(c.ModelName))
	fileContent += fmt.Sprintf("GetList(%s *c.ModelName.%s) ([]*c.ModelName.%s, error) \r\n", c.ModelName, strings.Title(c.ModelName), strings.Title(c.ModelName))
	fileContent += fmt.Sprintf("Save(%sc.ModelName *c.ModelName.%s) (bool, error)\r\n", c.ModelName, strings.Title(c.ModelName))
	fileContent += fmt.Sprintf("Delete(%s *c.ModelName.%s) (bool, error)\r\n", c.ModelName, strings.Title(c.ModelName))
	fileContent += "}\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @struct
	 **/
	fileContent += fmt.Sprintf("type %s struct {}\r\n", strings.Title(c.DaoName))
	fileContent += "\r\n"

	/**
	 * @step
	 * @getInfoById
	 **/
	createDaoInfo := CreateDaoInfo{DaoName: c.DaoName, ModelName: c.ModelName, AuthorName: thisAuthorName}
	infoContent, err := createDaoInfo.CreateContent()
	if err != nil {
		return "", err
	}
	fileContent += infoContent
	fileContent += "\r\n"

	/**
	 * @step
	 * @getList
	 **/
	createDaoList := CreateDaoList{DaoName: c.DaoName, ModelName: c.ModelName, AuthorName: thisAuthorName}
	listContent, err := createDaoList.CreateContent()
	if err != nil {
		return "", err
	}
	fileContent += listContent
	fileContent += "\r\n"

	/**
	 * @step
	 * @save
	 **/
	createDaoSave := CreateDaoSave{DaoName: c.DaoName, ModelName: c.ModelName, AuthorName: thisAuthorName}
	saveContent, err := createDaoSave.CreateContent()
	if err != nil {
		return "", err
	}
	fileContent += saveContent
	fileContent += "\r\n"

	/**
	 * @step
	 * @delete
	 **/
	createDaoDeleted := CreateDaoDeleted{DaoName: c.DaoName, ModelName: c.ModelName, AuthorName: thisAuthorName}
	deleteContent, err := createDaoDeleted.CreateContent()
	if err != nil {
		return "", nil
	}
	fileContent += deleteContent
	fileContent += "\r\n"
	return fileContent, nil
}

/**
 * @description: CreateContent
 * @param {string} c.DaoName
 * @param {string} c.ModelName
 * @param {string} c.AuthorName
 * @author: Jerry.Yang
 * @date: 2022-09-14 16:35:31
 * @return {*}
 */
func (c *CreateDaoInfo) CreateContent() (string, error) {
	fileContent := fmt.Sprintf("/** \r\n * @description: GetInfo \r\n * @param {*c.ModelName.%s} %s \r\n * @author: %s \r\n * @date: %s \r\n * @return {*} \r\n */ \r\n", strings.Title(c.ModelName), c.ModelName, c.AuthorName, time.Now().Format("2006-01-02 15:04:05"))
	fileContent += fmt.Sprintf("func (%s *%s) GetInfo(%s *c.ModelName.%s) (*c.ModelName.%s, error) {\r\n", c.DaoName[0:1], strings.Title(c.DaoName), c.ModelName, strings.Title(c.ModelName), strings.Title(c.ModelName))
	fileContent += fmt.Sprintf("result := c.ModelName.%s{}\r\n", strings.Title(c.ModelName))
	fileContent += fmt.Sprintf("if err := internal.DbClient().Where(%s).First(&result).Error; err != nil { \r\n", c.ModelName)
	fileContent += "internal.LoggorError(err)\r\n"
	fileContent += "return nil, err\r\n"
	fileContent += "}\r\n"
	fileContent += "return &result, nil\r\n"
	fileContent += "}\r\n"
	fileContent += "\r\n"
	return fileContent, nil
}

/**
 * @description: CreateDaoList
 * @param {string} c.DaoName
 * @param {string} c.ModelName
 * @param {string} c.AuthorName
 * @author: Jerry.Yang
 * @date: 2022-09-15 14:37:55
 * @return {*}
 */
func (c *CreateDaoList) CreateContent() (string, error) {
	fileContent := fmt.Sprintf("/** \r\n * @description: GetList \r\n * @param {*c.ModelName.%s} %s \r\n * @author: %s \r\n * @date: %s \r\n * @return {*} \r\n */ \r\n", strings.Title(c.ModelName), c.ModelName, c.AuthorName, time.Now().Format("2006-01-02 15:04:05"))
	fileContent += fmt.Sprintf("func (%s *%s) GetList(%s *c.ModelName.%s) ([]*c.ModelName.%s, error) {\r\n", c.DaoName[0:1], strings.Title(c.DaoName), c.ModelName, strings.Title(c.ModelName), strings.Title(c.ModelName))
	fileContent += fmt.Sprintf("result := make([]*c.ModelName.%s, 0)\r\n", strings.Title(c.ModelName))
	fileContent += fmt.Sprintf("if err := internal.DbClient().Where(%s).Find(&result).Error; err != nil {\r\n", c.ModelName)
	fileContent += "internal.LoggorError(err)\r\n"
	fileContent += "return nil, err\r\n"
	fileContent += "}\r\n"
	fileContent += "return result, nil\r\n"
	fileContent += "}\r\n"
	fileContent += "\r\n"
	return fileContent, nil
}

/**
 * @description: CreateDaoSave
 * @param {string} c.DaoName
 * @param {string} c.ModelName
 * @param {string} c.AuthorName
 * @author: Jerry.Yang
 * @date: 2022-09-15 14:38:46
 * @return {*}
 */
func (c *CreateDaoSave) CreateContent() (string, error) {
	fileContent := fmt.Sprintf("/** \r\n * @description: Save \r\n * @param {*c.ModelName.%s} %s \r\n * @author: %s \r\n * @date: %s \r\n * @return {*} \r\n */ \r\n", c.ModelName, strings.Title(c.ModelName), c.AuthorName, time.Now().Format("2006-01-02 15:04:05"))
	fileContent += fmt.Sprintf("func (%s *%s) Save(%s *c.ModelName.%s) (bool, error) {\r\n", c.DaoName[0:1], strings.Title(c.DaoName), c.ModelName, strings.Title(c.ModelName))
	fileContent += "result := false\r\n"
	fileContent += fmt.Sprintf("if err := internal.DbClient().Save(%s).Error; err != nil {\r\n", c.ModelName)
	fileContent += "internal.LoggorError(err)\r\n"
	fileContent += "return result, err\r\n"
	fileContent += "}\r\n"
	fileContent += "result = true\r\n"
	fileContent += "return result, nil\r\n"
	fileContent += "}\r\n"
	fileContent += "\r\n"
	return fileContent, nil
}

/**
 * @description: CreateDaoDeleted
 * @param {string} c.DaoName
 * @param {string} c.ModelName
 * @param {string} c.AuthorName
 * @author: Jerry.Yang
 * @date: 2022-09-15 14:42:04
 * @return {*}
 */
func (c *CreateDaoDeleted) CreateContent() (string, error) {
	fileContent := fmt.Sprintf("/** \r\n * @description: Delete \r\n * @param {*c.ModelName.%s} %s \r\n * @author: %s \r\n * @date: %s \r\n * @return {*} \r\n */ \r\n", strings.Title(c.ModelName), c.ModelName, c.AuthorName, time.Now().Format("2006-01-02 15:04:05"))
	fileContent += fmt.Sprintf("func (%s *%s) Delete(%s *c.ModelName.%s) (bool, error) {\r\n", c.DaoName[0:1], strings.Title(c.DaoName), c.ModelName, strings.Title(c.ModelName))
	fileContent += fmt.Sprintf("result := false\r\n")
	fileContent += fmt.Sprintf("if err := internal.DbClient().c.ModelName(&c.ModelName.%s{}).Where(%s).Update(\"is_deleted\", c.ModelName.IS_DELETED).Error; err != nil {\r\n", strings.Title(c.ModelName), c.ModelName)
	fileContent += "internal.LoggorError(err)\r\n"
	fileContent += "return result, err\r\n"
	fileContent += "}\r\n"
	fileContent += "result = true\r\n"
	fileContent += "return result, nil\r\n"
	fileContent += "}\r\n"
	fileContent += "\r\n"
	return fileContent, nil
}

/**
 * @description: checkParam
 * @param {string} projectName
 * @param {string} c.DaoName
 * @param {string} c.ModelName
 * @author: Jerry.Yang
 * @date: 2022-09-19 16:51:39
 * @return {*}
 */
func (c *CreateDao) CheckParams() error {
	if c.ProjectName == "" {
		return errors.New("projectName 为空!")
	}

	if c.DaoName == "" {
		return errors.New("c.DaoName 为空!")
	}

	if c.ModelName == "" {
		return errors.New("c.ModelName 为空!")
	}
	return nil
}
