package srun

import (
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func CustomIfaceGetRequest(url string, ifaceName string) (*http.Response, string, error) {
	var client *http.Client

	if ifaceName != "" {
		iface, err := net.InterfaceByName(ifaceName)
		if err != nil {
			log.Fatalf("failed to find network interface: %s", ifaceName)
		}

		addrs, err := iface.Addrs()
		if err != nil {
			log.Fatalf("failed to get addresses for interface: %s", ifaceName)
		}

		var localAddr string
		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					localAddr = ipNet.IP.String()
					break
				}
			}
		}

		if localAddr == "" {
			log.Fatalf("failed to get IPv4 address for interface: %s", ifaceName)
		}

		// 创建自定义HTTP客户端，绑定到指定网络接口的IP地址
		localTCPAddr := net.TCPAddr{
			IP: net.ParseIP(localAddr),
		}
		transport := &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				LocalAddr: &localTCPAddr,
				Timeout:   10 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
		}
		client = &http.Client{Transport: transport}

	} else {
		// 如果未指定接口名称，使用系统默认路由表
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
