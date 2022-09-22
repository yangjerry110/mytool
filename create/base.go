/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-19 16:57:50
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-20 10:55:19
 * @Description: base
 */
package create

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type CreateBaseInterface interface {
	CreateFile(dirName string, fileName string, fileContent string) error
}

type CreateBase struct{}

/**
 * @description: CreateFile
 * @param {string} dirName
 * @param {string} fileName
 * @param {string} fileContent
 * @author: Jerry.Yang
 * @date: 2022-09-19 17:22:46
 * @return {*}
 */
func (c *CreateBase) CreateFile(dirName string, fileName string, fileContent string) error {

	/**
	 * @step
	 * @定义文件名称
	 **/
	inputFile := fmt.Sprintf("%s/%s.go", dirName, fileName)

	/**
	 * @step
	 * @判断文件是否存在
	 **/
	_, err := os.Stat(inputFile)
	if err == nil {
		return errors.New(fmt.Sprintf("文件存在! filename = %s", inputFile))
	} else {
		if os.IsExist(err) {
			return errors.New(fmt.Sprintf("文件存在! filename = %s", inputFile))
		}
	}

	/**
	 * @step
	 * @定义osFile
	 **/
	var osFile *os.File

	/**
	 * @step
	 * @创建input文件
	 **/
	osFile, err = os.Create(inputFile)
	if err != nil {
		return errors.New(fmt.Sprintf("创建文件错误! filename = %s", inputFile))
	}
	defer osFile.Close()

	/**
	 * @step
	 * @写入内容
	 **/
	_, err = io.WriteString(osFile, fileContent)
	if err != nil {
		return errors.New(fmt.Sprintf("写入文件错误! filename = %s", inputFile))
	}

	fmt.Printf("%s 创建成功!", inputFile)
	fmt.Print("\r\n")

	/**
	 * @step
	 * @格式化代码
	 **/
	cmd := exec.Command("gofmt", "-w", inputFile)
	//cmd := exec.Command("ls", "-al", "/Users/admin/go/src/go.test.app/test")
	if err := cmd.Run(); err != nil {
		fmt.Printf("格式化错误! Err : %v", err.Error())
		fmt.Print("\r\n")
	}
	return nil
}
