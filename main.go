package main

import (
	"tqisci/common/logger"
)

func main() {
	logger.InitLogger()
	logger.Info("logger.Info(>>>>>>>>>>>>>&)")
	logger.Error(">>>> Error!")
}
