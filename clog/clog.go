/*
 * @Author: 杨小灿jian308@qq.com
 * @Date: 2023-05-12 16:24:48
 * @LastEditors: 杨小灿jian308@qq.com
 * @LastEditTime: 2023-05-13 10:00:23
 */
package clog

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger zerolog.Logger

func NewDevelopment() {
	//logger = log.With().Caller().Logger()
	//开发模式不在乎性能 增加了控制台颜色/日期格式/文件位置
	logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05"}).With().Caller().Logger()
}

func NewProduction(logfile string) {
	// 使用 lumberjack 实现 log rotate
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logfile,
		MaxSize:    100, // 单个文件最大100M
		MaxBackups: 60,  // 多于 60 个日志文件后，清理较旧的日志
		MaxAge:     1,   // 一天一切割
		Compress:   false,
	}
	fileWriteSyncer := zerolog.SyncWriter(lumberJackLogger)
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	logger = log.Level(zerolog.InfoLevel).Output(fileWriteSyncer).With().Caller().Logger()
}

func Info() *zerolog.Event {
	return logger.Info()
}

func Debug() *zerolog.Event {
	return logger.Debug()
}

func Error() *zerolog.Event {
	return logger.Error()
}

func Warn() *zerolog.Event {
	return logger.Warn()
}
