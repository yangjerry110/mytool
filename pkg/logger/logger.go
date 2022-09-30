/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-27 18:26:53
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-30 09:44:52
 * @Description: logger
 */
package logger

import (
	"github.com/yangjerry110/mytool/logger"
)

type LoggerPkgInterface interface {
	CreatePkgInterface(loggerPkgInterface LoggerPkgInterface) *LoggerPkg
	CreateInterface(loggerInterface logger.LoggerInterface) *LoggerPkg
	CreateOptionInterface(loggerOptionInterface logger.LoggerOptionInterface) *LoggerPkg
	SetOptions(options []logger.LoggerOptionFunc) LoggerPkgInterface
	SetDefaultOptions() LoggerPkgInterface
	Log(args ...interface{}) error                               // 记录对应级别的日志
	Logf(format string, args ...interface{}) error               // 记录对应级别的日志
	Info(args ...interface{}) error                              // 记录 InfoLevel 级别的日志
	Infof(format string, args ...interface{}) error              // 格式化并记录infof级别的日志
	Trace(args ...interface{}) error                             // 记录 TraceLevel 级别的日志
	Tracef(format string, args ...interface{}) error             // 格式化并记录 TraceLevel 级别的日志
	Debug(args ...interface{}) error                             // 记录 DebugLevel 级别的日志
	Debugf(format string, args ...interface{}) error             // 格式化并记录 DebugLevel 级别的日志
	Warn(args ...interface{}) error                              // 记录 WarnLevel 级别的日志
	Warnf(format string, args ...interface{}) error              // 格式化并记录 WarnLevel 级别的日志
	Error(args ...interface{}) error                             // 记录 ErrorLevel 级别的日志
	Errorf(format string, args ...interface{}) error             // 格式化并记录 ErrorLevel 级别的日志
	Fatal(args ...interface{}) error                             // 记录 FatalLevel 级别的日志
	Fatalf(format string, args ...interface{}) error             // 格式化并记录 FatalLevel 级别的日志
	Panic(args ...interface{}) error                             // 记录 PanicLevel 级别的日志
	Panicf(format string, args ...interface{}) error             // 格式化并记录 PanicLevel 级别的日志
	WithField(key string, value interface{}) LoggerPkgInterface  // 为日志添加一个上下文数据
	WithFields(fields map[string]interface{}) LoggerPkgInterface // 为日志添加多个上下文数据
	WithError(err error) LoggerPkgInterface                      // 为日志添加标准错误上下文数据
}

type LoggerPkg struct {
	OptionFunc            logger.LoggerOptionFunc
	Options               map[string]logger.LoggerOptionVal
	LoggerPkgInterface    LoggerPkgInterface
	LoggerInterface       logger.LoggerInterface
	LoggerOptionInterface logger.LoggerOptionInterface
}

// Level type
type Level uint32

/**
 * @description: CreateLogger
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:12:05
 * @return {*}
 */
var loggerNew = CreateLogger(&LogrusLogPkg{})

/**
 * @description: CreatePkgInterface
 * @param {LoggerPkgInterface} loggerPkgInterface
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:12:12
 * @return {*}
 */
func CreatePkgInterface(loggerPkgInterface LoggerPkgInterface) *LoggerPkg {
	return &LoggerPkg{LoggerPkgInterface: loggerPkgInterface}
}

/**
 * @description: CreateInterface
 * @param {logger.LoggerInterface} loggerInterface
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:12:18
 * @return {*}
 */
func CreateInterface(loggerInterface logger.LoggerInterface) *LoggerPkg {
	return &LoggerPkg{LoggerInterface: loggerInterface}
}

/**
 * @description: CreateOptionInterface
 * @param {logger.LoggerOptionInterface} loggerOptionInterface
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:12:25
 * @return {*}
 */
func CreateOptionInterface(loggerOptionInterface logger.LoggerOptionInterface) *LoggerPkg {
	return &LoggerPkg{LoggerOptionInterface: loggerOptionInterface}
}

/**
 * @description: CreateLogger
 * @param {LoggerPkgInterface} loggerPkgInterface
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:12:31
 * @return {*}
 */
func CreateLogger(loggerPkgInterface LoggerPkgInterface) LoggerPkgInterface {
	return loggerPkgInterface
}

/**
 * @description: SetLogger
 * @param {LoggerPkgInterface} LoggerPkgInterface
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:12:38
 * @return {*}
 */
func SetLogger(LoggerPkgInterface LoggerPkgInterface) LoggerPkgInterface {
	loggerNew = CreateLogger(LoggerPkgInterface)
	return LoggerPkgInterface
}

/**
 * @description: Log
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:13:06
 * @return {*}
 */
func Log(args ...interface{}) {
	loggerNew.Log(args...)
}

/**
 * @description: Logf
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:13:13
 * @return {*}
 */
func Logf(format string, args ...interface{}) {
	loggerNew.Logf(format, args...)
}

/**
 * @description: Trace
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:13:19
 * @return {*}
 */
func Trace(args ...interface{}) {
	loggerNew.Trace(args...)
}

/**
 * @description: Tracef
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:13:26
 * @return {*}
 */
func Tracef(format string, args ...interface{}) {
	loggerNew.Tracef(format, args...)
}

/**
 * @description: Debug
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:13:34
 * @return {*}
 */
func Debug(args ...interface{}) {
	loggerNew.Debug(args...)
}

/**
 * @description: Debugf
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:13:41
 * @return {*}
 */
func Debugf(format string, args ...interface{}) {
	loggerNew.Debugf(format, args...)
}

/**
 * @description: Info
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:13:47
 * @return {*}
 */
func Info(args ...interface{}) {
	loggerNew.Info(args...)
}

/**
 * @description: Infof
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:13:54
 * @return {*}
 */
func Infof(format string, args ...interface{}) {
	loggerNew.Infof(format, args...)
}

/**
 * @description: Warn
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:14:01
 * @return {*}
 */
func Warn(args ...interface{}) {
	loggerNew.Warn(args...)
}

/**
 * @description: Warnf
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:14:07
 * @return {*}
 */
func Warnf(format string, args ...interface{}) {
	loggerNew.Warnf(format, args...)
}

/**
 * @description: Error
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:14:14
 * @return {*}
 */
func Error(args ...interface{}) {
	loggerNew.Error(args...)
}

/**
 * @description: Errorf
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:14:21
 * @return {*}
 */
func Errorf(format string, args ...interface{}) {
	loggerNew.Errorf(format, args...)
}

/**
 * @description: Fatal
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:14:28
 * @return {*}
 */
func Fatal(args ...interface{}) {
	loggerNew.Fatal(args...)
}

/**
 * @description: Fatalf
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:14:35
 * @return {*}
 */
func Fatalf(format string, args ...interface{}) {
	loggerNew.Fatalf(format, args...)
}

/**
 * @description: Panic
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:14:42
 * @return {*}
 */
func Panic(args ...interface{}) {
	loggerNew.Panic(args...)
}

/**
 * @description: Panicf
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:14:49
 * @return {*}
 */
func Panicf(format string, args ...interface{}) {
	loggerNew.Panicf(format, args...)
}

/**
 * @description: WithField
 * @param {string} key
 * @param {interface{}} value
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:14:56
 * @return {*}
 */
func WithField(key string, value interface{}) LoggerPkgInterface {
	return loggerNew.WithField(key, value)
}

/**
 * @description: WithFields
 * @param {map[string]interface{}} fields
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:15:03
 * @return {*}
 */
func WithFields(fields map[string]interface{}) LoggerPkgInterface {
	return loggerNew.WithFields(fields)
}

/**
 * @description: WithError
 * @param {error} err
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:15:11
 * @return {*}
 */
func WithError(err error) LoggerPkgInterface {
	return loggerNew.WithError(err)
}
