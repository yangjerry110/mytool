/*
 * @Author: Jerry.Yang
 * @Date: 2022-11-10 19:12:26
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-15 10:29:17
 * @Description: pathConf
 */
package conf

import (
	"fmt"
	"reflect"
	"time"
)

type PathConf struct {
	FilePath       string
	FileType       string
	Intervals      time.Duration
	Data           interface{}
	LastModityTime time.Time
}

func (p *PathConf) GetNewConf() error {
	return p.GetConf()
}

func (p *PathConf) GetConf() error {

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
	p.Data = &Test{}

	/**
	 * @step
	 * @定义
	 **/
	files := []string{}

	/**
	 * @step
	 * @反射结构体，获取下一级的struct
	 **/
	dataStruct := reflect.ValueOf(p.Data).Elem()
	// fmt.Printf("dataStuct : %+v", dataStruct)
	// fmt.Print("\r\n")
	for i := 0; i < dataStruct.NumField(); i++ {

		/**
		 * @step
		 * @获取yaml
		 * @这里获取的yaml，就是yaml配置文件的文件名
		 **/
		yamlName := dataStruct.Type().Field(i).Tag.Get("yaml")
		files = append(files, yamlName)

		// p.GetConfInfo(yamlName, dataStruct.Field(i).Interface())
		// return nil

		/**
		 * @step
		 * @获取当前结构体的name
		 **/
		name := dataStruct.Type().Field(i).Name

		fmt.Printf("name : %v", name)
		fmt.Print("\r\n")

		/**
		 * @step
		 * @渲染配置
		 **/
		conf := &InitConf{
			FilePath:  p.FilePath,
			FileName:  yamlName,
			FileType:  p.FileType,
			Intervals: p.Intervals,
			//Data:      dataStruct.Field(i).Interface(),
		}

		fileData := map[interface{}]interface{}{}
		result := conf.Init(&fileData).GetParseConf()
		resultData := reflect.ValueOf(result)

		/**
		 * @step
		 * @获取下一级的struct
		 **/
		dataStruct.FieldByName(name).Set(resultData)

		fmt.Printf("dataStruct : %+v", dataStruct)
		fmt.Print("\r\n")
		return nil
	}

	return nil
}

func (p *PathConf) GetConfInfo(yamlName string, data interface{}) {
	conf := map[interface{}]interface{}{}

	initConfObj := &InitConf{
		FilePath:  p.FilePath,
		FileName:  yamlName,
		FileType:  p.FileType,
		Intervals: p.Intervals,
		//Data:      dataStruct.Field(i).Interface(),
	}
	parseConf := initConfObj.Init(conf).GetParseConf()

	fmt.Printf("data : %+v", data)
	fmt.Print("\r\n")

	fmt.Printf("result : %+v", parseConf)
	fmt.Print("\r\n")
}
