package common

import (
	logger "github.com/edgexfoundry/go-mod-core-contracts/clients/logging"
)

// Log 全局的loggerclient
var Log logger.LoggingClient

// InitLoggerClient 初始化loggerclient
func InitLoggerClient(logLevel string, logTarget string, serviceName string, isRemote bool) {
	Log = logger.NewClient(serviceName, isRemote, logTarget, logLevel)
}
