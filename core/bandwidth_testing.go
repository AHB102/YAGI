package core

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func BandwidthTesting() {
	url := "https://www.google.com/"

	startTime := time.Now()

	bytesTransferred, err := downloadFile(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	elapsedTime := time.Since(startTime)
	bandwidth := float64(bytesTransferred) / (1024.0 * 1024.0) / elapsedTime.Seconds()

	fmt.Printf("Downloaded %d bytes in %s\n", bytesTransferred, elapsedTime)
	fmt.Printf("Estimated Bandwidth: %.2f Mbps\n", bandwidth*8)
}

func downloadFile(url string) (int64, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to download file, status code: %d", resp.StatusCode)
	}

	var bytesTransferred int64
	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			bytesTransferred += int64(n)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return bytesTransferred, err
		}
	}

	return bytesTransferred, nil
}
