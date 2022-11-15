/*
 * @Author: Jerry.Yang
 * @Date: 2022-11-10 15:14:40
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-11 18:03:36
 * @Description: conf test
 */
package test

import (
	"fmt"
	"testing"
	"time"

	mytoolPkgConf "github.com/yangjerry110/mytool/conf"
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

	// mysqlTest := &MysqlTest{}
	// confObj1 := mytoolPkgConf.Conf{
	// 	FilePath:  "/Users/admin/go/src/my-tool/test/conf",
	// 	FileName:  "mysql.yaml",
	// 	FileType:  "yaml",
	// 	Intervals: time.Duration(2),
	// 	Data:      mysqlTest,
	// }
	// confObj1.GetNewConf()

	TestVal := &Test{}
	confObj := mytoolPkgConf.PathConf{
		FilePath: "/Users/admin/go/src/my-tool/test/conf",
		//FileName:  "redis.yaml",
		FileType:  "yaml",
		Intervals: time.Duration(2),
		Data:      TestVal,
	}
	// a := confObj.Init(TestVal).GetParseConf()

	confObj.GetNewConf()

	timeTickers := time.NewTicker(time.Second * 1)
	for range timeTickers.C {
		// fmt.Printf("testVal : %+v", a.(*Test))
		// fmt.Print("\r\n")

		fmt.Printf("test : %+v", TestVal)
		fmt.Print("\r\n")

		return
	}

}
