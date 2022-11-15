/*
 * @Author: Jerry.Yang
 * @Date: 2022-11-10 15:14:40
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-15 19:02:28
 * @Description: conf test
 */
package test

import (
	"fmt"
	"testing"
	"time"

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

	type TestConf struct {
		Addr    string `yaml:"addr"`
		Network string `yaml:"network"`
	}

	testConf := &TestConf{}
	mytoolPkgConf.Init("/Users/admin/go/src/my-tool/test/conf", "redis.yaml", "yaml", time.Second*2, testConf)

	result := mytoolPkgConf.GetParseConf("/Users/admin/go/src/my-tool/test/conf", "redis.yaml")

	fmt.Printf("result: %+v", result)
	fmt.Print("\r\n")

}
