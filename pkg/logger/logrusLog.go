/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-27 18:32:50
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-10-09 19:35:33
 * @Description: logrus
 */
package logger

import (
	"fmt"
	"os"

	"github.com/yangjerry110/mytool/logger"
)

type LogrusLogPkg struct {
	LoggerPkg
}

/**
 * @description: CreatePkgInterface
 * @param {LoggerPkgInterface} loggerPkgInterface
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:07:39
 * @return {*}
 */
func (l *LogrusLogPkg) CreatePkgInterface(loggerPkgInterface LoggerPkgInterface) *LoggerPkg {
	return &LoggerPkg{LoggerPkgInterface: l}
}

/**
 * @description: CreateInterface
 * @param {logger.LoggerInterface} loggerInterface
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:07:47
 * @return {*}
 */
func (l *LogrusLogPkg) CreateInterface(loggerInterface logger.LoggerInterface) *LoggerPkg {
	return &LoggerPkg{LoggerInterface: loggerInterface}
}

/**
 * @description: CreateOptionInterface
 * @param {logger.LoggerOptionInterface} LoggerOptionInterface
 * @author: Jerry.Yang
 * @date: 2022-09-30 14:36:55
 * @return {*}
 */
func (l *LogrusLogPkg) CreateOptionInterface(LoggerOptionInterface logger.LoggerOptionInterface) *LoggerPkg {
	return &LoggerPkg{LoggerOptionInterface: LoggerOptionInterface}
}

/**
 * @description: CreateOptionPkgInterface
 * @param {LoggerOptionPkgInterface} LoggerOptionPkgInterface
 * @author: Jerry.Yang
 * @date: 2022-09-30 14:44:25
 * @return {*}
 */
func (l *LogrusLogPkg) CreateOptionPkgInterface(LoggerOptionPkgInterface LoggerOptionPkgInterface) *LoggerPkg {
	return &LoggerPkg{LoggerOptionPkgInterface: LoggerOptionPkgInterface}
}

/**
 * @description: Log
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:09:01
 * @return {*}
 */
func (l *LogrusLogPkg) Log(args ...interface{}) error {
	return l.SetDefaultOptions().SetLoggerOptions().CreateInterface(&logger.LogrusLog{Options: &logger.LogrusOption{LoggerOption: logger.LoggerOption{Options: l.Options}}}).LoggerInterface.SetLevel().SetCallDept().SetWithFields().SetIsReportcaller().SetFormatter().SetOutput().SetLogger().WriteLog(args...)
}

/**
 * @description: Logf
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:09:08
 * @return {*}
 */
func (l *LogrusLogPkg) Logf(format string, args ...interface{}) error {
	return l.SetDefaultOptions().SetOptions(l.OptionFuns).SetLoggerOptions().CreateInterface(&logger.LogrusLog{Options: &logger.LogrusOption{LoggerOption: logger.LoggerOption{Options: l.Options}}}).LoggerInterface.SetLevel().SetCallDept().SetWithFields().SetIsReportcaller().SetFormatter().SetOutput().SetLogger().WriteLog(fmt.Sprintf(format, args...))
}

/**
 * @description: Info
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:09:19
 * @return {*}
 */
func (l *LogrusLogPkg) Info(args ...interface{}) error {
	return l.SetOptions([]logger.LoggerOptionFunc{CreateOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetLevel(Level(logger.InfoLevel))}).Log(args...)
}

/**
 * @description: Infof
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:09:27
 * @return {*}
 */
func (l *LogrusLogPkg) Infof(format string, args ...interface{}) error {
	return l.SetOptions([]logger.LoggerOptionFunc{CreateOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetLevel(Level(logger.InfoLevel))}).Logf(format, args...)
}

/**
 * @description: Warn
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:09:33
 * @return {*}
 */
func (l *LogrusLogPkg) Warn(args ...interface{}) error {
	return l.SetOptions([]logger.LoggerOptionFunc{CreateOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetLevel(Level(logger.WarnLevel))}).Log(args...)
}

/**
 * @description: Warnf
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:09:41
 * @return {*}
 */
func (l *LogrusLogPkg) Warnf(format string, args ...interface{}) error {
	return l.SetOptions([]logger.LoggerOptionFunc{CreateOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetLevel(Level(logger.WarnLevel))}).Logf(format, args...)
}

/**
 * @description: Trace
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:09:48
 * @return {*}
 */
func (l *LogrusLogPkg) Trace(args ...interface{}) error {
	return l.SetOptions([]logger.LoggerOptionFunc{CreateOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetLevel(Level(logger.TraceLevel))}).Log(args...)
}

/**
 * @description: Tracef
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:09:55
 * @return {*}
 */
func (l *LogrusLogPkg) Tracef(format string, args ...interface{}) error {
	return l.SetOptions([]logger.LoggerOptionFunc{CreateOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetLevel(Level(logger.TraceLevel))}).Logf(format, args...)
}

/**
 * @description: Debug
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:10:03
 * @return {*}
 */
func (l *LogrusLogPkg) Debug(args ...interface{}) error {
	return l.SetOptions([]logger.LoggerOptionFunc{CreateOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetLevel(Level(logger.DebugLevel))}).Log(args...)
}

/**
 * @description: Debugf
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:10:10
 * @return {*}
 */
func (l *LogrusLogPkg) Debugf(format string, args ...interface{}) error {
	return l.SetOptions([]logger.LoggerOptionFunc{CreateOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetLevel(Level(logger.DebugLevel))}).Logf(format, args...)
}

/**
 * @description: Error
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:10:18
 * @return {*}
 */
func (l *LogrusLogPkg) Error(args ...interface{}) error {
	return l.SetOptions([]logger.LoggerOptionFunc{CreateOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetLevel(Level(logger.ErrorLevel))}).Log(args...)
}

/**
 * @description: Errorf
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:10:26
 * @return {*}
 */
func (l *LogrusLogPkg) Errorf(format string, args ...interface{}) error {
	return l.SetOptions([]logger.LoggerOptionFunc{CreateOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetLevel(Level(logger.ErrorLevel))}).Logf(format, args...)
}

/**
 * @description: Fatal
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:10:32
 * @return {*}
 */
func (l *LogrusLogPkg) Fatal(args ...interface{}) error {
	return l.SetOptions([]logger.LoggerOptionFunc{CreateOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetLevel(Level(logger.FatalLevel))}).Log(args...)
}

/**
 * @description: Fatalf
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:10:39
 * @return {*}
 */
func (l *LogrusLogPkg) Fatalf(format string, args ...interface{}) error {
	return l.SetOptions([]logger.LoggerOptionFunc{CreateOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetLevel(Level(logger.FatalLevel))}).Logf(format, args...)
}

/**
 * @description: Panic
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:10:46
 * @return {*}
 */
func (l *LogrusLogPkg) Panic(args ...interface{}) error {
	return l.SetOptions([]logger.LoggerOptionFunc{CreateOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetLevel(Level(logger.PanicLevel))}).Log(args...)
}

/**
 * @description: Panicf
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:10:53
 * @return {*}
 */
func (l *LogrusLogPkg) Panicf(format string, args ...interface{}) error {
	return l.SetOptions([]logger.LoggerOptionFunc{CreateOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetLevel(Level(logger.PanicLevel))}).Logf(format, args...)
}

/**
 * @description: WithField
 * @param {string} key
 * @param {interface{}} value
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:11:01
 * @return {*}
 */
func (l *LogrusLogPkg) WithField(key string, value interface{}) LoggerPkgInterface {
	return l.SetOptions([]logger.LoggerOptionFunc{CreateOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetWithFields(map[string]interface{}{key: value})})
}

/**
 * @description: WithFields
 * @param {map[string]interface{}} fields
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:11:07
 * @return {*}
 */
func (l *LogrusLogPkg) WithFields(fields map[string]interface{}) LoggerPkgInterface {
	return l.SetOptions([]logger.LoggerOptionFunc{CreateOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetWithFields(fields)})
}

/**
 * @description: WithError
 * @param {error} err
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:11:14
 * @return {*}
 */
func (l *LogrusLogPkg) WithError(err error) LoggerPkgInterface {
	return l.SetOptions([]logger.LoggerOptionFunc{CreateOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetWithFields(map[string]interface{}{logger.LOGRUS_ERROR_KEY: err})})
}

/**
 * @description: SetOptions
 * @param {[]logger.LoggerOptionFunc} options
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:11:21
 * @return {*}
 */
func (l *LogrusLogPkg) SetOptions(options []logger.LoggerOptionFunc) LoggerPkgInterface {
	l.OptionFuns = append(l.OptionFuns, options...)
	return l
}

/**
 * @description: SetLoggerOptions
 * @author: Jerry.Yang
 * @date: 2022-09-30 14:39:48
 * @return {*}
 */
func (l *LogrusLogPkg) SetLoggerOptions() LoggerPkgInterface {
	optionFuncs := append(l.DefaultOptionFuns, l.OptionFuns...)
	logrusOptions := &logger.LogrusOption{}
	CreateOptionInterface(logrusOptions).LoggerOptionInterface.SetOptions(optionFuncs)
	l.Options = logrusOptions.Options
	return l
}

/**
 * @description: SetDefaultOptions
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:11:30
 * @return {*}
 */
func (l *LogrusLogPkg) SetDefaultOptions() LoggerPkgInterface {
	setDeafaultOptionFuncs := []logger.LoggerOptionFunc{
		CreateLoggerOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetLevel(Level(logger.WarnLevel)),
		CreateLoggerOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetFormatter(logger.LOGRUS_FORMATTER_JSON),
		CreateLoggerOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetIsReportcaller(true),
		CreateLoggerOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetOutput(os.Stdout),
		CreateLoggerOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetFormatterDisableHtmlEscap(true),
		CreateLoggerOptionPkgInterface(&LogrusOptionsPkg{}).LoggerOptionPkgInterface.SetCallerDept(5),
	}

	/**
	 * @step
	 * @首先创建默认的option
	 **/
	l.DefaultOptionFuns = setDeafaultOptionFuncs
	return l
}
