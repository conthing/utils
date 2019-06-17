package common

import (
	"github.com/sirupsen/logrus"
)

// Log 全局的loggerclient
var Log = logrus.New()

// InitLoggerClient 初始化loggerclient
func InitLoggerClient(logLevel string, logTarget string, serviceName string, isRemote bool) {
	fmt.Printf("%s %s %s %v",logLevel,logTarget,serviceName,isRemote)
	if logTarget == ""{
		Log.Out = os.Stdout
	}else{
		file, err := os.OpenFile(logTarget, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.Out = file
		} else {
			log.Info("Failed to log to file, using default stderr")
		}
	}
}
