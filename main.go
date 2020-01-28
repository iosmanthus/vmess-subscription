package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/iosmathus/vmess-subscription/subscribe"
	uuid "github.com/satori/go.uuid"
	"v2ray.com/core/infra/conf"
)

func checkFlags() {
	switch {
	case *subscribeURL == "":
		log.Fatal("missing subscription url")
	case *injectFile == "":
		log.Fatal("missing inject file")
	}
}

func main() {
	flag.Parse()
	checkFlags()

	sub := subscribe.NewSubscription(*subscribeURL)
	outbounds, err := sub.Fetch(flag.Args()...)
	if err != nil {
		log.Fatal(err)
	}

	selected := selectOutbound(outbounds)
	transitTag := tagOutBound(selected)

	config, err := parseInjectFile(*injectFile)
	if err != nil {
		log.Fatal(err)
	}

	injectProxySettings(config, transitTag)

	config.OutboundConfigs = append(config.OutboundConfigs, *selected)

	if err := outputConfig(config); err != nil {
		log.Fatal(err)
	}
}

func selectOutbound(outbounds []*conf.OutboundDetourConfig) *conf.OutboundDetourConfig {
	rand.Seed(time.Now().UTC().UnixNano())
	return outbounds[rand.Intn(len(outbounds))]
}

func parseInjectFile(path string) (*conf.Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config conf.Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func tagOutBound(outbound *conf.OutboundDetourConfig) string {
	outbound.Tag = uuid.NewV4().String()
	return outbound.Tag
}

func injectProxySettings(config *conf.Config, transitTag string) {
	for i := range config.OutboundConfigs {
		if config.OutboundConfigs[i].StreamSetting == nil {
			config.OutboundConfigs[i].ProxySettings = &conf.ProxyConfig{Tag: transitTag}
		}
	}
}

func outputConfig(config *conf.Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	if *outputFile == "" {
		fmt.Println(string(data))
	} else {
		if err := ioutil.WriteFile(*outputFile, data, os.ModeAppend); err != nil {
			return nil
		}
	}
	return nil
}
