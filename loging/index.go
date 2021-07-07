package loging

import (
	"go.uber.org/zap"
	"log"
)

var ZapLog *zap.Logger
var ZapSugarLog *zap.SugaredLogger

func init() {
	initZap()
}

//初始化zap
func initZap() {
	ZapLog, _ = zap.NewProduction()
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

	//sugar.Infow("failed to fetch URL",
	//	// Structured context as loosely typed key-value pairs.
	//	"url", url,
	//	"attempt", 3,
	//	"backoff", time.Second,
	//)
	//sugar.Infof("Failed to fetch URL: %s", url)
}

//var (
//	Trace   *log.Logger // 记录所有日志
//	Info    *log.Logger // 重要的信息
//	Warning *log.Logger // 需要注意的信息
//	Error   *log.Logger // 非常严重的问题
//)
//go 自带的日志配置
//func localLog() {
////设置日志
//log.SetPrefix("TRACE: ")
//log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
//
//file, err := os.OpenFile("errors.txt",
//	os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
//if err != nil {
//	log.Fatalln("Failed to open error log file:", err)
//}
//
//Trace = log.New(ioutil.Discard,
//	"TRACE: ",
//	log.Ldate|log.Lmicroseconds|log.Lshortfile)
//
//Info = log.New(os.Stdout,
//	"INFO: ",
//	log.Ldate|log.Lmicroseconds|log.Lshortfile)
//
//Warning = log.New(os.Stdout,
//	"WARNING: ",
//	log.Ldate|log.Lmicroseconds|log.Lshortfile)
//
//Error = log.New(io.MultiWriter(file, os.Stderr),
//	"ERROR: ",
//	log.Ldate|log.Lmicroseconds|log.Lshortfile)
//}
