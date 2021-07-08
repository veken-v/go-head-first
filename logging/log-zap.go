//zap 日志处理
package logging

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"os"
	"strings"
)

var ZapLog *zap.Logger
var ZapSugarLog *zap.SugaredLogger

//初始化zap
func initZap() {
	//ZapLog, _ = zap.NewProduction()

	writeSyncer := getLogWriter()
	encoder := getEncoder()
	//日志级别
	var logLevel zapcore.Level
	switch strings.ToLower(loggingConfigurer.Level) {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "info":
		logLevel = zapcore.InfoLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	default:
		logLevel = zapcore.InfoLevel
	}
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)

	ZapLog = zap.New(core, zap.AddCaller())
	defer func(ZapLogger *zap.Logger) {
		err := ZapLogger.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}(ZapLog) // flushes buffer, if any

	ZapSugarLog = ZapLog.Sugar()
	defer func(ZapSugarLog *zap.SugaredLogger) {
		err := ZapSugarLog.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}(ZapSugarLog)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//zapcore.NewJSONEncoder(encoderConfig)
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   defaultStr(loggingConfigurer.FilePath, "./log/app.log"), //日志名
		MaxSize:    defaultInt(loggingConfigurer.FileMaxSize, 100),          //每10M会切割
		MaxBackups: defaultInt(loggingConfigurer.MaxBackups, 0),             //最多保存5个备份
		MaxAge:     defaultInt(loggingConfigurer.MaxHistory, 0),             //最多保存30天的数据
		Compress:   loggingConfigurer.Compress,                              //备份日志是否压缩
		LocalTime:  loggingConfigurer.LocalTime,
	}

	//return zapcore.AddSync(lumberJackLogger)
	//打印到控制台,与文件
	writer := io.MultiWriter(lumberJackLogger, os.Stderr)
	return zapcore.AddSync(writer)
}

func defaultStr(targetStr string, defaultStr string) string {
	if strings.Trim(targetStr, "") == "" {
		return defaultStr
	}
	return targetStr
}

func defaultInt(targetInt int, defaultInt int) int {
	if targetInt <= 0 {
		return defaultInt
	}
	return targetInt
}

//sugar.Infow("failed to fetch URL",
//	// Structured context as loosely typed key-value pairs.
//	"url", url,
//	"attempt", 3,
//	"backoff", time.Second,
//)
//sugar.Infof("Failed to fetch URL: %s", url)
