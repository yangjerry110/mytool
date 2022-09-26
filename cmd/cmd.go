/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-23 18:35:52
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 16:31:34
 * @Description: cmd
 */
package cmd

type CmdInterface interface {
	CreateContent() (string, error)
	CheckParams() error
}

type Cmd struct {
	CmdInterface CmdInterface
}

type CreateApp struct {
	ProjectName string
	AppName     string
	Method      string
}

type CreateAppInputVO struct {
	CreateApp
	DirName  string
	FileName string
}

type CreateAppOutputVO struct {
	CreateApp
	DirName  string
	FileName string
}

type CreateAppRoute struct {
	CreateApp
	DirName  string
	FileName string
}

type CreateAppService struct {
	CreateApp
	DirName  string
	FileName string
}

type CreateDao struct {
	ProjectName string
	DaoName     string
	ModelName   string
	AuthorName  []string
}

type CreateDaoInfo struct {
	ProjectName string
	DaoName     string
	ModelName   string
	AuthorName  string
}

type CreateDaoList struct {
	ProjectName string
	DaoName     string
	ModelName   string
	AuthorName  string
}

type CreateDaoSave struct {
	ProjectName string
	DaoName     string
	ModelName   string
	AuthorName  string
}

type CreateDaoDeleted struct {
	ProjectName string
	DaoName     string
	ModelName   string
	AuthorName  string
}
