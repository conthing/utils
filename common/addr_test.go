package common

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMajorInterface(t *testing.T) {
	i := GetMajorInterface()
	log.Println(i)
	assert.NotEqual(t, nil, i, "failed to get major interface")
}

func TestGetSerialNumber(t *testing.T) {
	sn := GetSerialNumber()
	log.Println(sn)
	assert.NotEqual(t, nil, sn, "failed to get sn")
}

func TestGetMajorInterfaceIP(t *testing.T) {
	SetMajorInterface("无线网络连接")
	ip := GetMajorInterfaceIP()
	log.Println(ip)
	assert.NotEqual(t, nil, ip, "failed to get ip")
}

func TestSetMajorInterface(t *testing.T) {
	SetMajorInterface("无线网络连接")
	sn1 := GetSerialNumber()
	log.Println(sn1)
	SetMajorInterface("本地连接")
	sn2 := GetSerialNumber()
	log.Println(sn2)
	assert.NotEqual(t, sn1, sn2, "SetMajorInterface failed")
}
