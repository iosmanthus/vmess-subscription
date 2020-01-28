package subscribe

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/v2fly/vmessping/miniv2ray"
	"github.com/v2fly/vmessping/vmess"
	"v2ray.com/core/infra/conf"
)

type Subscription struct {
	url string
}

func NewSubscription(url string) Subscription {
	return Subscription{url}
}

func (s *Subscription) Fetch(options ...string) ([]*conf.OutboundDetourConfig, error) {
	url := s.url
	for _, option := range options {
		url += "&" + option
	}

	links, err := getVmessLinks(url)
	if err != nil {
		return nil, err
	}

	outbounds := parseVmessLinks(links)
	if outbounds == nil {
		return nil, fmt.Errorf("No links parse successfully")
	}

	return outbounds, nil
}

func getVmessLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	encodedLinks, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// Decode
	rawLinks, err := base64.StdEncoding.DecodeString(string(encodedLinks))
	links := strings.Split(string(rawLinks), "\n")

	return links, nil
}

func parseVmessLinks(links []string) []*conf.OutboundDetourConfig {
	outbounds := make([]*conf.OutboundDetourConfig, 0, len(links))
	for _, link := range links {
		subscription, err := vmess.ParseVmess(link)
		if err != nil {
			continue
		}
		conf := miniv2ray.Vmess2OutboundConfig(subscription)
		outbounds = append(outbounds, conf)
	}
	if len(outbounds) == 0 {
		return nil
	}
	return outbounds
}
