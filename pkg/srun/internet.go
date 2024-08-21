package srun

import (
	"fmt"
	"net/url"
	"strings"
)

// Internet 通过Http请求是否被302到 srun登陆地址 来判断是否能访问互联网
func (s PortalServer) Internet() bool {
	reqUrl, _ := url.ParseRequestURI(s.internetCheck)
	// resp, _, errs := gorequest.New().Get(reqUrl.String()).End()
	resp, body, errs := CustomIfaceGetRequest(reqUrl.String(), s.iface)
	if errs != nil {
		fmt.Println(errs)
		return false
	}
	if (resp.Request.URL.Hostname() == reqUrl.Hostname()) && !(strings.Contains(body, "Authentication is required. Click ")) {
		return true
	}
	return false
}
