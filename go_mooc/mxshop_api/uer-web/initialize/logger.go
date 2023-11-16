package initialize

import "go.uber.org/zap"

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	//defer logger.Sync()
	//suger:=logger.Sugar()

	//省略上面三步骤，zap.S()=L 但是需要我们自己设置一个全局logger
	zap.ReplaceGlobals(logger)
}
