package log

import (
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loglv = zap.DebugLevel
var logio = false
var logpath = "./"

// 设置日志级别
func SetLevel(level zapcore.Level) {
	loglv = level
	intLog()
}
func SetPath(path string) {
	logpath = path
	logio = true
	intLog()
}

// 实现Log接口中的方法
func Debug(args ...interface{}) {
	sugarLogger.Debug(args...)
}

func Info(args ...interface{}) {
	sugarLogger.Info(args...)
}

func Warn(args ...interface{}) {
	sugarLogger.Warn(args...)
}

func Error(args ...interface{}) {
	sugarLogger.Error(args...)
}

func Fatal(args ...interface{}) {
	sugarLogger.Fatal(args...)
}

func Debugf(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	sugarLogger.Fatalf(template, args...)
}

var sugarLogger *zap.SugaredLogger

func init() {
	intLog()
}

func intLog() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, loglv)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)) //显示调用函数行数 封装后需要加1才能显示正确
	sugarLogger = logger.Sugar()
	defer sugarLogger.Sync()
}

func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = formatEncodeTime
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //这里可以指定颜色
	//encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	if !logio {
		return zapcore.AddSync(os.Stdout)
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logpath + "log.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
