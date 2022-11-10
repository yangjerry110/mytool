/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-27 15:59:26
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 15:14:31
 * @Description:
 */
package test

import (
	"testing"

	mytoolLogger "github.com/yangjerry110/mytool/logger"
	"github.com/yangjerry110/mytool/pkg/logger"
)

func TestLogger(t *testing.T) {
	return
	logger.SetOptions([]mytoolLogger.LoggerOptionFunc{
		logger.SetLevel(logger.Level(mytoolLogger.WarnLevel)),
	}).Error("11111")
}
