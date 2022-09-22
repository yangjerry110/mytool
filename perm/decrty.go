/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-20 19:28:19
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-20 19:37:25
 * @Description: decrty 解密
 */
package perm

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

type (
	DecrtyInterface interface {
		Decrty(permPath string, inputStr string) (string, error)
		doRasDecrty(inputStr string) (string, error)
	}

	RasDecrty struct{}
)

/**
 * @description: Decrty
 * @param {string} permPath
 * @param {string} inputStr
 * @author: Jerry.Yang
 * @date: 2022-09-20 19:32:21
 * @return {*}
 */
func (r *RasDecrty) Decrty(permPath string, inputStr string) (string, error) {

	/**
	 * @step
	 * @判断参数
	 **/
	if permPath == "" {
		return "", errors.New("permPath must be not empty!")
	}

	if inputStr == "" {
		return "", errors.New("inputStr must be not empty!")
	}

	/**
	 * @step
	 * @判断文件是否存在
	 **/
	_, err := os.Stat(fmt.Sprintf("%s/private.perm", permPath))
	if err != nil {
		if os.IsNotExist(err) {
			return "", errors.New("permPath is err; private.perm is not exist!")
		}
		return "", err
	}

	/**
	 * @step
	 * @创建通道
	 **/
	doRasDecrtyChan := make(chan string)

	/**
	 * @step
	 * @协程处理
	 **/
	go func() {
		outputStr, err := r.doRasDecrty(inputStr)
		if err != nil {
			doRasDecrtyChan <- ""
		} else {
			doRasDecrtyChan <- outputStr
		}
		close(doRasDecrtyChan)
	}()

	/**
	 * @step
	 * @获取通道数据，并且判断
	 **/
	outputStr := <-doRasDecrtyChan
	if outputStr == "" {
		return "", errors.New("decrty fail!")
	}
	return outputStr, nil
}

/**
 * @description: doRasDecrty
 * @param {string} inputStr
 * @author: Jerry.Yang
 * @date: 2022-09-20 19:32:12
 * @return {*}
 */
func (r *RasDecrty) doRasDecrty(inputStr string) (string, error) {
	/**
	 * @step
	 * @进行base64 decode
	 **/
	inputStrByte, err := base64.StdEncoding.DecodeString(inputStr)
	if err != nil {
		return inputStr, err
	}

	/**
	 * @step
	 * @打开公钥文件
	 **/
	file, err := os.Open(fmt.Sprintf("%s/private.perm"))
	if err != nil {
		return inputStr, err
	}

	/**
	 * @step
	 * @读取文件内容
	 **/
	fileInfo, err := file.Stat()
	if err != nil {
		return inputStr, err
	}

	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	file.Close()

	/**
	 * @step
	 * @pem编码
	 **/
	block, _ := pem.Decode(buf)
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @使用私钥解密
	 **/
	outputByte, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, inputStrByte)
	if err != nil {
		return inputStr, err
	}
	return string(outputByte), nil
}
