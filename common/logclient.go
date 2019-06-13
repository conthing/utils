package common

import (
	logger "github.com/edgexfoundry/go-mod-core-contracts/clients/logging"
)

// LogClient 全局的loggerclient
var LogClient logger.LoggingClient

// InitLoggerClient 初始化loggerclient
func InitLoggerClient(serviceName string, isRemote bool, logTarget string, logLevel string) {
	LogClient = logger.NewClient(serviceName, isRemote, logTarget, logLevel)
}
