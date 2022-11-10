/*
 * @Author: Jerry.Yang
 * @Date: 2022-11-10 15:14:40
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 18:14:57
 * @Description: conf test
 */
package test

import (
	"fmt"
	"testing"
	"time"

	mytoolPkgConf "github.com/yangjerry110/mytool/conf"
)

func TestConf(t *testing.T) {

	type Test struct {
		NetWork string `yaml:"network"`
		Addr    string `yaml:"addr"`
	}

	TestVal := &Test{}
	confObj := mytoolPkgConf.Conf{
		FilePath:  "/Users/admin/go/src/my-tool/test/conf",
		FileName:  "redis.yaml",
		FileType:  "yaml",
		Intervals: time.Duration(5),
		Data:      TestVal,
	}

	confObj.GetHotUpdateConf()

	//mytoolPkgConf.GetYamlConf("/Users/admin/go/src/my-tool/test/conf/redis.yaml", TestVal)

	timeTickers := time.NewTicker(time.Second * 1)
	for range timeTickers.C {
		fmt.Printf("testVal : %+v", TestVal)
		fmt.Print("\r\n")
	}

}
