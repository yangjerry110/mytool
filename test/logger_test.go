/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-27 15:59:26
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-10-09 19:19:53
 * @Description:
 */
package test

import (
	"testing"

	"github.com/yangjerry110/mytool/pkg/logger"
)

func TestLogger(t *testing.T) {
	// logger.SetOptions([]toolLogger.LoggerOptionFunc{
	// 	logger.SetIsReportcaller(true),
	// 	logger.SetLevel(logger.Level(toolLogger.DebugLevel)),
	// }).WithField("testFields", "testFieldVal").Info("this is test withField")

	logger.WithField("test", "testVal").WithField("test1", "testval").Infof("this is test infof")
}
