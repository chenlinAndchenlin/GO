package main

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	//cfg := &zap.Config{}
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./myproject.log",
		"stderr",
		//"stdout",
	}
	return cfg.Build()
}

func main() {
	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	url := "www.google.com"
	sugar.Infow("failed to connect:",
		"url:", url,
		"attempting to connect", 3)
	sugar.Infof("Failed to fetch URL: %s", url)
}

/*func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	url := "www.google.com"
	sugar.Infow("failed to connect:",
		"url:", url,
		"attempting to connect", 3)
	sugar.Infof("Failed to fetch URL: %s", url)
	println("***********************************************")
	logger, _ = zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
*/
