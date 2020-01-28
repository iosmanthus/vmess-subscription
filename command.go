package main

import "flag"

var injectFile = flag.String("inject", "", "file that inject proxy settings")
var subscribeURL = flag.String("subscribe", "", "url of vmess subscription")
var outputFile = flag.String("output", "", "output file")
