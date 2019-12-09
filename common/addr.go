package common

import (
	"fmt"
	"net"
)

var ethName string
var serialNumber string

func init() {
	netInterface, err := net.InterfaceByName("eth0")
	if err != nil || netInterface.HardwareAddr.String() == "08:00:3e:26:0a:5b" { //有些板子eth0能获取出这个地址
		netInterface, err = net.InterfaceByName("eth1")
		if err != nil {
			Log.Errorf("can not get MAC of eth0/eth1")
			ethName = ""
			serialNumber = ""
			return
		}
		ethName = "eth1"
	} else {
		ethName = "eth0"
	}
	serialNumber = fmt.Sprintf("%x", []byte(netInterface.HardwareAddr))
}

func GetMajorInterface() string {
	if ethName == "" {
		init()
	}
	return ethName
}

func GetSerialNumber() string {
	if serialNumber == "" {
		init()
	}
	return serialNumber
}

func GetMacAddrByName(name string) (macAddr string) {
	netInterface, err := net.InterfaceByName(name)
	if err != nil {
		Log.Errorf("fail to get net interface by name %s: %v", name, err)
		return
	}

	return netInterface.HardwareAddr.String()
}

func GetMacHexStringByName(name string) (macAddr string) {
	netInterface, err := net.InterfaceByName(name)
	if err != nil {
		Log.Errorf("fail to get net interface by name %s: %v", name, err)
		return
	}

	return fmt.Sprintf("%x", []byte(netInterface.HardwareAddr))
}

func GetMacAddrs() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		Log.Errorf("fail to get net interfaces: %v", err)
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}

		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}

func GetIPs() (ips []string) {

	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		Log.Errorf("fail to get net interface addrs: %v", err)
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}
