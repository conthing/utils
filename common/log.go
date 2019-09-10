package common

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// LoggerConfig 配置结构体定义
type LoggerConfig struct {
	Level      string
	File       string
	SkipCaller bool
	Service    string
}

// Logger 结构体定义
type Logger struct {
	ZeroLog zerolog.Logger
}

// Log 全局的loggerclient
var Log = Logger{ZeroLog: log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMilli}).Level(zerolog.DebugLevel).With().CallerWithSkipFrameCount(3).Logger()}

func levelFromString(levelString string) zerolog.Level {
	switch strings.ToUpper(levelString) {
	case "ERROR", "ERR":
		return zerolog.ErrorLevel
	case "WARN", "WARNING":
		return zerolog.WarnLevel
	case "INFO":
		return zerolog.InfoLevel
	default:
		return zerolog.DebugLevel
	}
}

// InitLogger 初始化logger，可以不调用此Init，Log会初始化成默认值
func InitLogger(config *LoggerConfig) {
	if config == nil {
		return
	}

	var file *os.File
	var err error
	var out string

	level := levelFromString(config.Level)
	context := log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMilli}).Level(level).With()
	if config.File == "" {
		out = "Log to stderr"
	} else {
		file, err = os.OpenFile(config.File, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			out = fmt.Sprintf("Failed to log to file \"%s\", using stderr", config.File)
		} else {
			context = log.Output(zerolog.ConsoleWriter{Out: file, TimeFormat: time.StampMilli, NoColor: true}).Level(level).With()
			out = fmt.Sprintf("Log to file \"%s\"", config.File)
		}
	}

	if !config.SkipCaller {
		context = context.CallerWithSkipFrameCount(3)
	}
	if config.Service != "" {
		context = context.Str("service", config.Service)
	}
	Log.ZeroLog = context.Logger()

	Log.Infof("Log level:%s. %s", level, out)
}

// Errorf format方式输出error日志
func (logger Logger) Errorf(format string, v ...interface{}) {
	logger.ZeroLog.Error().Msgf(format, v...)
}

// Warnf format方式输出warn日志
func (logger Logger) Warnf(format string, v ...interface{}) {
	logger.ZeroLog.Warn().Msgf(format, v...)
}

// Infof format方式输出info日志
func (logger Logger) Infof(format string, v ...interface{}) {
	logger.ZeroLog.Info().Msgf(format, v...)
}

// Debugf format方式输出debug日志
func (logger Logger) Debugf(format string, v ...interface{}) {
	logger.ZeroLog.Debug().Msgf(format, v...)
}

// Error string方式输出error日志
func (logger Logger) Error(v ...interface{}) {
	logger.ZeroLog.Error().Msg(fmt.Sprint(v...))
}

// Warn string方式输出warn日志
func (logger Logger) Warn(v ...interface{}) {
	logger.ZeroLog.Warn().Msg(fmt.Sprint(v...))
}

// Info string方式输出info日志
func (logger Logger) Info(v ...interface{}) {
	logger.ZeroLog.Info().Msg(fmt.Sprint(v...))
}

// Debug string方式输出debug日志
func (logger Logger) Debug(v ...interface{}) {
	logger.ZeroLog.Debug().Msg(fmt.Sprint(v...))
}
