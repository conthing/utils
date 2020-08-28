package common

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Config 配置文件基本组成
type Config struct {
	LockType    string
	CameraType  string
	PlateIDType string // CameraType 自带识别功能时，PlateIDType = ""
	LedType     string
	HTTP        HTTP
	Redis       Redis
}

// Redis 配置
type Redis struct {
	Host string //sqlite3时，表示存储文件的路径和文件名
	Port int
}

// HTTP 配置
type HTTP struct {
	Port int
}

var config Config

func TestLoadYaml(t *testing.T) {
	LoadYaml("", &config)

	log.Printf("config:%+v", config)
	// assert equality
	assert.Equal(t, 52032, config.HTTP.Port, "they should be equal")
}
