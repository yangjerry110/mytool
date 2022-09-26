package cmdPkg

import "github.com/yangjerry110/mytool/cmd"

type CreateAppPkg struct{}

/**
 * @description: CreateAppContent
 * @param {string} projectName
 * @param {string} appName
 * @param {string} method
 * @author: Jerry.Yang
 * @date: 2022-09-19 17:43:23
 * @return {*}
 */
func CreateAppContent(projectName string, appName string, method string) (string, error) {
	return CreateCmdInterface(&cmd.CreateApp{ProjectName: projectName, AppName: appName, Method: method}).cmdInterface.CreateContent()
}

/**
 * @description: CreateAppInputVoContent
 * @param {string} projectName
 * @param {string} appName
 * @param {string} method
 * @param {string} dirName
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-09-23 19:05:15
 * @return {*}
 */
func CreateAppInputVoContent(projectName string, appName string, method string, dirName string, fileName string) (string, error) {
	return CreateCmdInterface(&cmd.CreateAppInputVO{CreateApp: cmd.CreateApp{ProjectName: projectName, AppName: appName, Method: method}, DirName: dirName, FileName: fileName}).cmdInterface.CreateContent()
}

/**
 * @description: CreateAppOutputVoContent
 * @param {string} projectName
 * @param {string} appName
 * @param {string} method
 * @param {string} dirName
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-09-23 23:57:23
 * @return {*}
 */
func CreateAppOutputVoContent(projectName string, appName string, method string, dirName string, fileName string) (string, error) {
	return CreateCmdInterface(&cmd.CreateAppOutputVO{CreateApp: cmd.CreateApp{ProjectName: projectName, AppName: appName, Method: method}, DirName: dirName, FileName: fileName}).cmdInterface.CreateContent()
}

/**
 * @description: CreateAppRouteContent
 * @param {string} projectName
 * @param {string} appName
 * @param {string} method
 * @param {string} dirName
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-09-23 23:57:33
 * @return {*}
 */
func CreateAppRouteContent(projectName string, appName string, method string, dirName string, fileName string) (string, error) {
	return CreateCmdInterface(&cmd.CreateAppRoute{CreateApp: cmd.CreateApp{ProjectName: projectName, AppName: appName, Method: method}, DirName: dirName, FileName: fileName}).cmdInterface.CreateContent()
}

/**
 * @description: CreateAppServiceContent
 * @param {string} projectName
 * @param {string} appName
 * @param {string} method
 * @param {string} dirName
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-09-23 23:57:41
 * @return {*}
 */
func CreateAppServiceContent(projectName string, appName string, method string, dirName string, fileName string) (string, error) {
	return CreateCmdInterface(&cmd.CreateAppService{CreateApp: cmd.CreateApp{ProjectName: projectName, AppName: appName, Method: method}, DirName: dirName, FileName: fileName}).cmdInterface.CreateContent()
}
