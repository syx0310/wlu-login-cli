package srun

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"syscall"
	"time"
)

func CustomIfaceGetRequest(url string, ifaceName string) (*http.Response, string, error) {
	var client *http.Client

	if ifaceName != "" {
		// Create a custom HTTP client, binding it directly to the specified network interface
		transport := &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				d := &net.Dialer{
					Timeout:   10 * time.Second,
					KeepAlive: 30 * time.Second,
					Control: func(network, address string, c syscall.RawConn) error {
						return c.Control(func(fd uintptr) {
							// Bind to the specified network interface
							err := syscall.SetsockoptString(int(fd), syscall.SOL_SOCKET, syscall.SO_BINDTODEVICE, ifaceName)
							if err != nil {
								log.Fatalf("failed to bind to network interface: %s", ifaceName)
							}
						})
					},
				}
				return d.DialContext(ctx, network, addr)
			},
		}
		client = &http.Client{Transport: transport}

	} else {
		// If no interface name is specified, use the system's default routing table
		client = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	resp, err := client.Get(url)
	if err != nil {
		return resp, "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, "", err
	}

	return resp, string(body), nil
}
