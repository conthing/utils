package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig1(t *testing.T) {
	InitLogger(&LoggerConfig{Level: "WARN", SkipCaller: true})

	Log.Error("This is a ERROR")
	Log.Warn("This is a WARNING")
	Log.Info("This is a INFO")
	Log.Debug("This is a DEBUG")
	Log.Errorf("This is a ERROR, %02x", 123)
	Log.Warnf("This is a WARNING, %d", 456)
	Log.Infof("This is a INFO, %s", "String")
	Log.Debugf("This is a DEBUG, %v", Log)

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

}
