/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-27 15:59:26
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-29 18:47:43
 * @Description:
 */
package test

import (
	"testing"

	"github.com/yangjerry110/mytool/pkg/logger"
)

func TestLogger(t *testing.T) {
	logger.SetLogger(&logger.TestLog{})
	logger.Info("thisErr : %s", "this is err")
}
