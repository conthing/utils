package common

import (
	"fmt"
	"net"
)

var ethName string
var serialNumber string

// Init 依据既定的顺序查询网卡是否存在（不需要Up），如果存在就作为 MajorInterface ，即将其MAC作为设备序列号
func Init() {
	ethNameList := []string{"eth0", "eth1", "wlan0", "本地连接", "无线网络连接", "en0"}

	for _, name := range ethNameList {
		netInterface, err := net.InterfaceByName(name)
		if err == nil {
			ethName = name
			serialNumber = fmt.Sprintf("%x", []byte(netInterface.HardwareAddr))
			return
		}
	}
	Log.Errorf("can not get MAC of %v", ethNameList)
}

// SetMajorInterface 设置 MajorInterface ，设置后就不再依赖既定的顺序去查询了，原来的MajorInterface会被改写
func SetMajorInterface(name string) error {
	netInterface, err := net.InterfaceByName(name)
	if err == nil {
		ethName = name
		serialNumber = fmt.Sprintf("%x", []byte(netInterface.HardwareAddr))
		return nil
	}
	return fmt.Errorf("can not get MAC of %v, error:%v", name, err)
}

// GetMajorInterface 读取 MajorInterface ，如果原来没设置过，会调用Init
func GetMajorInterface() string {
	if ethName == "" {
		Init()
	}
	return ethName
}

// GetSerialNumber 读取 SerialNumber ，如果原来没设置过，会调用Init
func GetSerialNumber() string {
	if serialNumber == "" {
		Init()
	}
	return serialNumber
}

// GetMacAddrByName 根据网口名称查询MAC地址
func GetMacAddrByName(name string) (macAddr string) {
	netInterface, err := net.InterfaceByName(name)
	if err != nil {
		Log.Errorf("fail to get net interface by name %s: %v", name, err)
		return
	}

	return netInterface.HardwareAddr.String()
}

// 下面的函数没有使用过

// GetMacHexStringByName 根据网口名称查询MAC地址（不带冒号）
func GetMacHexStringByName(name string) (macAddr string) {
	netInterface, err := net.InterfaceByName(name)
	if err != nil {
		Log.Errorf("fail to get net interface by name %s: %v", name, err)
		return
	}

	return fmt.Sprintf("%x", []byte(netInterface.HardwareAddr))
}
