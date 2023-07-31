package core

import (
	"fmt"
	"net"
	"os"
)

// NetworkTopology retrieves the network topology by getting the list of all network interfaces and their addresses.
func NetworkTopology() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	topology := make(map[string][]string)

	for _, iface := range interfaces {
		addresses, err := iface.Addrs()
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, address := range addresses {
			if ipnet, ok := address.(*net.IPNet); ok && ipnet.IP.IsGlobalUnicast() {
				topology[address.String()] = append(topology[address.String()], iface.Name)
			}
		}
	}

	fmt.Println("Network Topology:")
	for address, ifaces := range topology {
		fmt.Println("*", address)
		for _, iface := range ifaces {
			fmt.Println("  ", iface)
		}
	}
}
