package core

import (
	"fmt"
	"os/exec"
	"time"
)

func Traceroute(destination string, maxHops int, timeout time.Duration) error {
	cmd := exec.Command("tracert", "-h", fmt.Sprintf("%d", maxHops), "-w", fmt.Sprintf("%d", int(timeout.Seconds()*1000)), destination)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("traceroute failed: %s\nOutput: %s", err, output)
	}

	fmt.Println(string(output))

	return nil
}
