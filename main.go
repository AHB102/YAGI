package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/AHB102/YAGi/core"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("YAGI:")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "exit" {
			break
		}

		// Split the input into command and destination (if available)
		args := strings.Fields(input)
		if len(args) < 1 {
			fmt.Println("Invalid input. Usage: command [destination]")
			continue
		}

		command := args[0]

		switch command {
		case "help", "-h":
			showHelp()
		case "tr":
			if len(args) < 2 {
				fmt.Println("Invalid input for traceroute. Usage: traceroute | destination")
				continue
			}
			destination := args[1]
			err := core.Traceroute(destination, 30, 5*time.Second)
			if err != nil {
				fmt.Println(err)
			}
		case "rDNS":
			if len(args) < 2 {
				fmt.Printf("Invalid input for reverseDNS. \nUsage: reverseDNS | destination")
				continue
			}
			destination := args[1]
			names, err := core.ReverseDNS(destination)
			if err != nil {
				fmt.Println("Error performing reverse DNS lookup:", err)
				continue
			}

			if len(names) == 0 {
				fmt.Printf("No domain names found for the given IP address.")
				continue
			}

			fmt.Printf("Domain names associated with IP address %s:", destination)
			for _, name := range names {
				fmt.Println(name)
			}
		case "bwidth":
			core.BandwidthTesting()
		case "DHCPManage":
			if len(args) < 2 {
				fmt.Printf("Invalid input for DHCPManage. \nUsage: DHCPManage -subCommand [release|renew|flush]")
				continue
			}

			subCommand := args[1]

			switch subCommand {
			case "release":
				err := core.ReleaseDHCPLease()
				if err != nil {
					fmt.Println("Failed to release DHCP lease:", err)
				} else {
					fmt.Printf("DHCP lease released successfully.\n")
				}

			case "renew":
				err := core.RenewDHCPLease()
				if err != nil {
					fmt.Println("Failed to renew DHCP lease:", err)
				} else {
					fmt.Printf("DHCP lease renewed successfully.\n")
				}

			case "flush":
				err := core.FlushDHCPLease()
				if err != nil {
					fmt.Println("Failed to flush DHCP lease:", err)
				} else {
					fmt.Printf("DHCP lease flushed successfully.\n")
				}

			default:
				fmt.Printf("Invalid subcommand for DHCPManage. Usage: DHCPManage -subcommand [release|renew|flush]")
			}
		case "nettopo":
			core.NetworkTopology()
		case "netinfo":
			core.PrintNetworkInfo()
		case "portscan":
			if len(args) < 2 {
				fmt.Printf("Invalid input for portscan. \nUsage: portscan | destination")
				continue
			}
			destination := args[1]
			core.GetOpenPorts(destination, core.PortRange{Start: 1, End: 65535})
		default:
			fmt.Printf("Unknown command:", command, "\n")
		}
	}

}

func showHelp() {
	fmt.Println("YAGI - Yet Another Go-based CLI Tool")
	fmt.Println("Usage: command [destination]")
	fmt.Println("Available commands:")
	fmt.Println("  tr   	    							   - Perform a traceroute to the specified destination.")
	fmt.Println("  rDNS         							   - Perform a reverse DNS lookup for the given IP address.")
	fmt.Println("  bwidth       							   - Perform bandwidth testing.")
	fmt.Println("  DHCPManage-subcommand [release|renew|flush] - Manage DHCP leases.")
	fmt.Println("  nettopo      							   - Display the network topology. * represents root")
	fmt.Println("  netinfo      							   - Print network information.")
	fmt.Println("  portscan  								   - Perform a port scan on the specified destination.")
	fmt.Println("  help, -h     							   - Show this help message.")
	fmt.Println("  exit         							   - Exit the program.")

}
