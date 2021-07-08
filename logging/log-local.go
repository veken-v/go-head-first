//本地日志处理
package logging

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
