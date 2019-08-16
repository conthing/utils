package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMacAddrs1(t *testing.T) {
	assert.Equal(t, "1", GetMacAddrs(), "Bootstrap should return error")
	assert.Equal(t, "1", GetIPs(), "Bootstrap should return error")
	assert.Equal(t, GetMacHexStringByName("本地连接"), GetMacAddrByName("本地连接"), "Bootstrap should return error")
}
