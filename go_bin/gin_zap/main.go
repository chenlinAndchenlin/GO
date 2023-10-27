package main

//
//import (
//	"go.uber.org/zap/zapcore"
//	"net/http"
//)
//import "go.uber.org/zap"
//
//var logger *zap.Logger
//var sugarLogger *zap.SugaredLogger
//
//func main() {
//	InitLogger()
//	defer sugarLogger.Sync()
//	simpleHttpGet("www.sogo.com")
//
//	println("***************************************")
//	simpleHttpGet("http://www.baidu.com")
//}
//
////	func InitLogger() {
////		logger, _ = zap.NewProduction()
////	}
////
////	func simpleHttpGet(url string) {
////		resp, err := http.Get(url)
////		if err != nil {
////			logger.Error(
////				"Error fetching url..",
////				zap.String("url", url),
////				zap.Error(err))
////		} else {
////			logger.Info("Success..",
////				zap.String("statusCode", resp.Status),
////				zap.String("url", url))
////			resp.Body.Close()
////		}
////	}
////
////	func InitLogger() {
////		logger, _ = zap.NewProduction()
////		sugarLogger = logger.Sugar()
////	}
//func InitLogger() {
//	writeSyncer := getLogWriter()
//	encoder := getEncoder()
//	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
//
//	//logger := zap.New(core)
//	logger := zap.New(core, zap.AddCaller())
//	sugarLogger = logger.Sugar()
//}
//
////func InitLogger() {
////	encoder := getEncoder()
////	// test.log记录全量日志
////	logF, _ := os.Create("./test.log")
////	c1 := zapcore.NewCore(encoder, zapcore.AddSync(logF), zapcore.DebugLevel)
////	// test.err.log记录ERROR级别的日志
////	errF, _ := os.Create("./test.err.log")
////	c2 := zapcore.NewCore(encoder, zapcore.AddSync(errF), zap.ErrorLevel)
////	// 使用NewTee将c1和c2合并到core
////	core := zapcore.NewTee(c1, c2)
////	logger = zap.New(core, zap.AddCaller())
////	sugarLogger = logger.Sugar()
////}
//
//func getEncoder() zapcore.Encoder {
//	encoderConfig := zap.NewProductionEncoderConfig()
//	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
//	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
//	return zapcore.NewConsoleEncoder(encoderConfig)
//
//	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
//	//return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
//}
//
////	func getLogWriter() zapcore.WriteSyncer {
////		//file, _ := os.Create("./test.log")
////		file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
////		return zapcore.AddSync(file)
////	}
//func getLogWriter() zapcore.WriteSyncer {
//	lumberJackLogger := &lumberjack.Logger{
//		Filename:   "./test.log",
//		MaxSize:    10,
//		MaxBackups: 5,
//		MaxAge:     30,
//		Compress:   false,
//	}
//	return zapcore.AddSync(lumberJackLogger)
//}
//
//func simpleHttpGet(url string) {
//	sugarLogger.Debugf("Trying to hit GET request for %s", url)
//	resp, err := http.Get(url)
//	if err != nil {
//		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
//	} else {
//		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
//		resp.Body.Close()
//	}
//}
