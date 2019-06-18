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
var Log = Logger{ZeroLog: log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).Level(zerolog.DebugLevel).With().CallerWithSkipFrameCount(3).Logger()}

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

	if config.File == "" {
		file = os.Stderr
		out = "Log to stderr"
	} else {
		file, err = os.OpenFile(config.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			file = os.Stderr
			out = fmt.Sprintf("Failed to log to file \"%s\", using stderr", config.File)
		} else {
			out = fmt.Sprintf("Log to file \"%s\"", config.File)
		}
	}
	level := levelFromString(config.Level)
	if config.Service != "" {
		Log.ZeroLog = log.Output(zerolog.ConsoleWriter{Out: file, TimeFormat: time.RFC3339}).Level(level).With().CallerWithSkipFrameCount(3).Str("service", config.Service).Logger()
	} else {
		Log.ZeroLog = log.Output(zerolog.ConsoleWriter{Out: file, TimeFormat: time.RFC3339}).Level(level).With().CallerWithSkipFrameCount(3).Logger()
	}
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
func (logger Logger) Error(msg string) {
	logger.ZeroLog.Error().Msg(msg)
}

// Warn string方式输出warn日志
func (logger Logger) Warn(msg string) {
	logger.ZeroLog.Warn().Msg(msg)
}

// Info string方式输出info日志
func (logger Logger) Info(msg string) {
	logger.ZeroLog.Info().Msg(msg)
}

// Debug string方式输出debug日志
func (logger Logger) Debug(msg string) {
	logger.ZeroLog.Debug().Msg(msg)
}
