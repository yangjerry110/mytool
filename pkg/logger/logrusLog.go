/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-27 18:32:50
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-29 19:11:31
 * @Description: logrus
 */
package logger

import (
	"fmt"
	"io"
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
 * @param {logger.LoggerOptionInterface} loggerOptionInterface
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:07:55
 * @return {*}
 */
func (l *LogrusLogPkg) CreateOptionInterface(loggerOptionInterface logger.LoggerOptionInterface) *LoggerPkg {
	return &LoggerPkg{LoggerOptionInterface: loggerOptionInterface}
}

/**
 * @description: SetLevel
 * @param {Level} level
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:08:02
 * @return {*}
 */
func (l *LogrusLogPkg) SetLevel(level Level) logger.LoggerOptionFunc {
	return CreateOptionInterface(&logger.LogrusOption{}).LoggerOptionInterface.SetLevel(logger.Level(level))
}

/**
 * @description: SetWithFields
 * @param {map[string]interface{}} fields
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:08:15
 * @return {*}
 */
func (l *LogrusLogPkg) SetWithFields(fields map[string]interface{}) logger.LoggerOptionFunc {
	return CreateOptionInterface(&logger.LogrusOption{}).LoggerOptionInterface.SetWithFields(fields)
}

/**
 * @description: SetIsReportcaller
 * @param {bool} isOpen
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:08:23
 * @return {*}
 */
func (l *LogrusLogPkg) SetIsReportcaller(isOpen bool) logger.LoggerOptionFunc {
	return CreateOptionInterface(&logger.LogrusOption{}).LoggerOptionInterface.SetIsReportcaller(isOpen)
}

/**
 * @description: SetOutput
 * @param {io.Writer} output
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:08:31
 * @return {*}
 */
func (l *LogrusLogPkg) SetOutput(output io.Writer) logger.LoggerOptionFunc {
	return CreateOptionInterface(&logger.LogrusOption{}).LoggerOptionInterface.SetOutput(output)
}

/**
 * @description: SetFormatter
 * @param {string} format
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:08:39
 * @return {*}
 */
func (l *LogrusLogPkg) SetFormatter(format string) logger.LoggerOptionFunc {
	return CreateOptionInterface(&logger.LogrusOption{}).LoggerOptionInterface.SetFormatter(format)
}

/**
 * @description: SetFormatterDisableHtmlEscap
 * @param {bool} isOpen
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:08:47
 * @return {*}
 */
func (l *LogrusLogPkg) SetFormatterDisableHtmlEscap(isOpen bool) logger.LoggerOptionFunc {
	return CreateOptionInterface(&logger.LogrusOption{}).LoggerOptionInterface.SetFormatterDisableHtmlEscap(isOpen)
}

/**
 * @description: SetFormatterDisableTime
 * @param {bool} isOpen
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:08:54
 * @return {*}
 */
func (l *LogrusLogPkg) SetFormatterDisableTime(isOpen bool) logger.LoggerOptionFunc {
	return CreateOptionInterface(&logger.LogrusOption{}).LoggerOptionInterface.SetFormatterDisableTime(isOpen)
}

/**
 * @description: Log
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:09:01
 * @return {*}
 */
func (l *LogrusLogPkg) Log(args ...interface{}) error {
	return CreateInterface(&logger.LogrusLog{Options: &logger.LogrusOption{LoggerOption: &logger.LoggerOption{Options: l.Options}}}).LoggerInterface.SetLevel().SetFormatter().SetOutput().SetLogger().WriteLog(args...)
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
	return CreateInterface(&logger.LogrusLog{Options: &logger.LogrusOption{LoggerOption: &logger.LoggerOption{Options: l.Options}}}).LoggerInterface.SetLevel().SetFormatter().SetOutput().SetLogger().WriteLog(fmt.Sprintf(format, args...))
}

/**
 * @description: Info
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:09:19
 * @return {*}
 */
func (l *LogrusLogPkg) Info(args ...interface{}) error {
	return l.SetDefaultOptions().SetOptions([]logger.LoggerOptionFunc{l.SetLevel(Level(logger.InfoLevel))}).Log(args...)
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
	return l.SetDefaultOptions().SetOptions([]logger.LoggerOptionFunc{l.SetLevel(Level(logger.InfoLevel))}).Logf(format, args...)
}

/**
 * @description: Warn
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:09:33
 * @return {*}
 */
func (l *LogrusLogPkg) Warn(args ...interface{}) error {
	return l.SetDefaultOptions().SetOptions([]logger.LoggerOptionFunc{l.SetLevel(Level(logger.WarnLevel))}).Log(args...)
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
	return l.SetDefaultOptions().SetOptions([]logger.LoggerOptionFunc{l.SetLevel(Level(logger.WarnLevel))}).Logf(format, args...)
}

/**
 * @description: Trace
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:09:48
 * @return {*}
 */
func (l *LogrusLogPkg) Trace(args ...interface{}) error {
	return l.SetDefaultOptions().SetOptions([]logger.LoggerOptionFunc{l.SetLevel(Level(logger.TraceLevel))}).Log(args...)
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
	return l.SetDefaultOptions().SetOptions([]logger.LoggerOptionFunc{l.SetLevel(Level(logger.TraceLevel))}).Logf(format, args...)
}

/**
 * @description: Debug
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:10:03
 * @return {*}
 */
func (l *LogrusLogPkg) Debug(args ...interface{}) error {
	return l.SetDefaultOptions().SetOptions([]logger.LoggerOptionFunc{l.SetLevel(Level(logger.DebugLevel))}).Log(args...)
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
	return l.SetDefaultOptions().SetOptions([]logger.LoggerOptionFunc{l.SetLevel(Level(logger.DebugLevel))}).Logf(format, args...)
}

/**
 * @description: Error
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:10:18
 * @return {*}
 */
func (l *LogrusLogPkg) Error(args ...interface{}) error {
	return l.SetDefaultOptions().SetOptions([]logger.LoggerOptionFunc{l.SetLevel(Level(logger.ErrorLevel))}).Log(args...)
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
	return l.SetDefaultOptions().SetOptions([]logger.LoggerOptionFunc{l.SetLevel(Level(logger.ErrorLevel))}).Logf(format, args...)
}

/**
 * @description: Fatal
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:10:32
 * @return {*}
 */
func (l *LogrusLogPkg) Fatal(args ...interface{}) error {
	return l.SetDefaultOptions().SetOptions([]logger.LoggerOptionFunc{l.SetLevel(Level(logger.FatalLevel))}).Log(args...)
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
	return l.SetDefaultOptions().SetOptions([]logger.LoggerOptionFunc{l.SetLevel(Level(logger.FatalLevel))}).Logf(format, args...)
}

/**
 * @description: Panic
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:10:46
 * @return {*}
 */
func (l *LogrusLogPkg) Panic(args ...interface{}) error {
	return l.SetDefaultOptions().SetOptions([]logger.LoggerOptionFunc{l.SetLevel(Level(logger.PanicLevel))}).Log(args...)
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
	return l.SetDefaultOptions().SetOptions([]logger.LoggerOptionFunc{l.SetLevel(Level(logger.PanicLevel))}).Logf(format, args...)
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
	return l.SetDefaultOptions().SetOptions([]logger.LoggerOptionFunc{l.SetWithFields(map[string]interface{}{key: value})})
}

/**
 * @description: WithFields
 * @param {map[string]interface{}} fields
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:11:07
 * @return {*}
 */
func (l *LogrusLogPkg) WithFields(fields map[string]interface{}) LoggerPkgInterface {
	return l.SetDefaultOptions().SetOptions([]logger.LoggerOptionFunc{l.SetWithFields(fields)})
}

/**
 * @description: WithError
 * @param {error} err
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:11:14
 * @return {*}
 */
func (l *LogrusLogPkg) WithError(err error) LoggerPkgInterface {
	return l.SetDefaultOptions().SetOptions([]logger.LoggerOptionFunc{l.SetWithFields(map[string]interface{}{logger.LOGRUS_ERROR_KEY: err})})
}

/**
 * @description: SetOptions
 * @param {[]logger.LoggerOptionFunc} options
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:11:21
 * @return {*}
 */
func (l *LogrusLogPkg) SetOptions(options []logger.LoggerOptionFunc) LoggerPkgInterface {
	logrusOption := &logger.LogrusOption{LoggerOption: &logger.LoggerOption{Options: l.Options}}
	CreateOptionInterface(logrusOption).LoggerOptionInterface.SetOptions(options)
	l.Options = logrusOption.Options
	return l
}

/**
 * @description: SetDefaultOptions
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:11:30
 * @return {*}
 */
func (l *LogrusLogPkg) SetDefaultOptions() LoggerPkgInterface {
	setDeafaultOptionFuncs := []logger.LoggerOptionFunc{}

	/**
	 * @step
	 * @判断哪些需要赋值默认配置的
	 * @level
	 **/
	if l.Options[logger.OPTION_LOGLEVEL].Value == nil {
		setDeafaultOptionFuncs = append(setDeafaultOptionFuncs, l.SetLevel(Level(logger.WarnLevel)))
	}

	/**
	 * @step
	 * @formatter
	 **/
	if l.Options[logger.OPTION_FORMATTER_TYPE].Value == nil {
		setDeafaultOptionFuncs = append(setDeafaultOptionFuncs, l.SetFormatter(logger.LOGRUS_FORMATTER_JSON))
	}

	/**
	 * @step
	 * @check isReportcaller
	 **/
	if l.Options[logger.OPTION_ISREPORTCALLER].Value == nil {
		setDeafaultOptionFuncs = append(setDeafaultOptionFuncs, l.SetIsReportcaller(true))
	}

	/**
	 * @step
	 * @check output
	 **/
	if l.Options[logger.OPTION_OUTPUT].Value == nil {
		setDeafaultOptionFuncs = append(setDeafaultOptionFuncs, l.SetOutput(os.Stdout))
	}

	/**
	 * @step
	 * @check FormatterDisableHtmlEscap
	 **/
	if l.Options[logger.OPTION_FORMATTER_DISABLEHTMLESCAP].Value == nil {
		setDeafaultOptionFuncs = append(setDeafaultOptionFuncs, l.SetFormatterDisableHtmlEscap(true))
	}

	/**
	 * @step
	 * @首先创建默认的option
	 **/
	l.SetOptions(setDeafaultOptionFuncs)
	return l
}
