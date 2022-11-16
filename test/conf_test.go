/*
 * @Author: Jerry.Yang
 * @Date: 2022-11-10 15:14:40
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-16 17:13:52
 * @Description: conf test
 */
package test

import (
	"fmt"
	"testing"
	"time"

	mytoolConf "github.com/yangjerry110/mytool/conf"
	mytoolPkgConf "github.com/yangjerry110/mytool/pkg/conf"
)

type Redis struct {
	NetWork string `yaml:"network"`
	Addr    string `yaml:"addr"`
}

type MysqlTest struct {
	Dsn string `yaml:"dsn"`
}

type Test struct {
	Redis     Redis     `yaml:"redis.yaml"`
	MysqlTest MysqlTest `yaml:"mysql.yaml"`
}

func TestConf(t *testing.T) {
	return
	type TestConf struct {
		Addr    string `yaml:"addr"`
		Network string `yaml:"network"`
	}

	testConf := &TestConf{}

	testConf1 := &TestConf{}

	testConf2 := &TestConf{}
	// mytoolPkgConf.Init("/Users/admin/go/src/my-tool/test/conf", "redis.yaml", "yaml", time.Second*2, testConf)

	// result := mytoolPkgConf.GetParseConf("/Users/admin/go/src/my-tool/test/conf", "redis.yaml")

	// fmt.Printf("result: %+v", result)
	// fmt.Print("\r\n")

	mytoolPkgConf.GetConf("/Users/admin/go/src/my-tool/test/conf", "redis.yaml", "yaml", time.Duration(2), testConf)

	mytoolConf := mytoolConf.Conf{
		FilePath:  "/Users/admin/go/src/my-tool/test/conf",
		FileName:  "redis.yaml",
		FileType:  "yaml",
		Intervals: time.Duration(2),
		Data:      testConf1,
	}
	mytoolConf.GetNewConf()

	mytoolPkgConf.Init("/Users/admin/go/src/my-tool/test/conf", "redis.yaml", "yaml", time.Duration(2), testConf2)
	result2 := mytoolPkgConf.GetParseConf("/Users/admin/go/src/my-tool/test/conf", "redis.yaml")

	timeTickers := time.NewTicker(1 * time.Second)

	for range timeTickers.C {

		fmt.Printf("result: %+v", testConf)
		fmt.Print("\r\n")

		fmt.Printf("result1: %+v", testConf1)
		fmt.Print("\r\n")

		fmt.Printf("result2 : %v", result2)
		fmt.Print("\r\n")
	}

}
