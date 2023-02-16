/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-27 15:59:26
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-02-15 17:03:19
 * @Description:
 */
package test

import (
	"fmt"
	"testing"

	mytoolLogger "github.com/yangjerry110/mytool/logger"
	"github.com/yangjerry110/mytool/pkg/logger"
)

func TestLogger(t *testing.T) {
	return
	a()

	var loggerPkg = &logger.LoggerPkg{}
	fmt.Printf("logger : %+v", loggerPkg)
	fmt.Printf("\r\n")

	logger.Errorf("11111")
	fmt.Printf("logger: %+v", loggerPkg)
	fmt.Print("\r\n")
}

func a() logger.LoggerPkgInterface {
	return logger.SetOptions([]mytoolLogger.LoggerOptionFunc{
		logger.SetLevel(logger.Level(mytoolLogger.PanicLevel)),
	})
}
