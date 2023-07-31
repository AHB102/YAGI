package core

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func DHCPManagement() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [release|renew|flush]")
		return
	}

	command := os.Args[1]

	switch command {
	case "release":
		err := ReleaseDHCPLease()
		if err != nil {
			fmt.Println("Failed to release DHCP lease:", err)
		} else {
			fmt.Println("DHCP lease released successfully.")
		}

	case "renew":
		err := RenewDHCPLease()
		if err != nil {
			fmt.Println("Failed to renew DHCP lease:", err)
		} else {
			fmt.Println("DHCP lease renewed successfully.")
		}

	case "flush":
		err := FlushDHCPLease()
		if err != nil {
			fmt.Println("Failed to flush DHCP lease:", err)
		} else {
			fmt.Println("DHCP lease flushed successfully.")
		}

	default:
		fmt.Println("Usage: go run main.go [release|renew|flush]")
	}
}

func ReleaseDHCPLease() error {
	switch runtime.GOOS {
	case "windows":
		return runCommand("ipconfig", "/release")
	case "linux":
		return runCommand("dhclient", "-r")
	default:
		return fmt.Errorf("DHCP lease release is not supported on %s", runtime.GOOS)
	}
}

func RenewDHCPLease() error {
	switch runtime.GOOS {
	case "windows":
		return runCommand("ipconfig", "/renew")
	case "linux":
		return runCommand("dhclient")
	default:
		return fmt.Errorf("DHCP lease renew is not supported on %s", runtime.GOOS)
	}
}

func FlushDHCPLease() error {
	switch runtime.GOOS {
	case "windows":
		return runCommand("ipconfig", "/flushdns")
	case "linux":
		return runCommand("service", "network-manager", "restart")
	default:
		return fmt.Errorf("DHCP lease flush is not supported on %s", runtime.GOOS)
	}
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
