package common

import (
	"fmt"
	"net"
)

func GetMacAddrByName(name string) (macAddr string) {
	netInterface, err := net.InterfaceByName(name)
	if err != nil {
		fmt.Printf("fail to get net interface by name %s: %v", name, err)
		return
	}

	return netInterface.HardwareAddr.String()
}

func GetSerialNumber() string {
	netInterface, err := net.InterfaceByName("eth0")
	if err != nil {
		netInterface, err = net.InterfaceByName("eth1")
		if err != nil {
			netInterface, err = net.InterfaceByName("en0")
			if err != nil {
				log.Fatal("无法获取mac地址")
			}
		}
	}

	sn := fmt.Sprintf("%x", []byte(netInterface.HardwareAddr))
	return sn
}

func GetMacHexStringByName(name string) (macAddr string) {
	netInterface, err := net.InterfaceByName(name)
	if err != nil {
		fmt.Printf("fail to get net interface by name %s: %v", name, err)
		return
	}

	return fmt.Sprintf("%x", []byte(netInterface.HardwareAddr))
}

func GetMacAddrs() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
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
		fmt.Printf("fail to get net interface addrs: %v", err)
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
