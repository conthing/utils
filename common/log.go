package common

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Log 全局的loggerclient
var Log = logrus.New()
var logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).With().Str("app", "znet").Caller().Logger()

// InitLoggerClient 初始化loggerclient
func InitLoggerClient(logLevel string, logTarget string, serviceName string, isRemote bool) {
	fmt.Printf("%s %s %s %v",logLevel,logTarget,serviceName,isRemote)
	var file io.Writer
	if logTarget == ""{
		file = os.Stderr
	}else{
		file, err := os.OpenFile(logTarget, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.Out = file
		} else {
			log.Info("Failed to log to file, using default stderr")
			file = os.Stderr
		}
	}
	if serviceName != "" {
		Log = log.Output(zerolog.ConsoleWriter{Out: file, TimeFormat: time.RFC3339}).With().Str("app", serviceName).Caller().Logger()
	} else {
		Log = log.Output(zerolog.ConsoleWriter{Out: file, TimeFormat: time.RFC3339}).With().Caller().Logger()
	}
}
