package core

import (
	"fmt"
	"net"
)

// PrintNetworkInfo prints the network information for all interfaces on the system.
func PrintNetworkInfo() {
	// Get the network interfaces on the system.
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// For each network interface, print the IP addresses, subnet masks, default gateways, and MAC addresses.
	for _, iface := range interfaces {
		fmt.Println("Interface:", iface.Name)
		fmt.Println("IP addresses:")
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		for _, addr := range addrs {
			fmt.Println("  -", addr)
		}

		fmt.Println("Subnet mask:", getSubnetMask(&iface))
		fmt.Println("Default gateway:", getDefaultGateway(&iface))
		fmt.Println("MAC address:", iface.HardwareAddr)
		fmt.Println()
	}
}

// getSubnetMask returns the subnet mask for a given network interface.
func getSubnetMask(iface *net.Interface) string {
	addrs, err := iface.Addrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		// Parse CIDR to get the subnet mask.
		_, mask, _ := net.ParseCIDR(addr.String())
		return mask.String()
	}
	return ""
}

// getDefaultGateway returns the default gateway IP address for a given network interface.
func getDefaultGateway(iface *net.Interface) string {
	ifaceAddrs, err := iface.Addrs()
	if err != nil {
		return ""
	}
	for _, ifaceAddr := range ifaceAddrs {
		ip, _, err := net.ParseCIDR(ifaceAddr.String())
		if err != nil {
			continue
		}

		if ip.To4() != nil {
			routes, err := net.InterfaceAddrs()
			if err != nil {
				return ""
			}

			for _, route := range routes {
				if route.String() == "0.0.0.0/0" {
					gatewayIP := getIPv4Gateway(route.(*net.IPNet))
					return gatewayIP.String()
				}
			}
		}
	}

	return ""
}

// getIPv4Gateway returns the IPv4 gateway IP address for a given network route.
func getIPv4Gateway(ipnet *net.IPNet) net.IP {
	if ipnet == nil {
		return nil
	}

	ip := ipnet.IP.To4()
	if ip == nil {
		return nil
	}

	ip[3] |= 1
	return ip
}
