module github.com/iosmathus/vmess-subscription

go 1.13

require (
	github.com/satori/go.uuid v1.2.0
	github.com/v2fly/vmessping v0.2.3
	v2ray.com/core v4.19.1+incompatible
)

replace github.com/v2fly/vmessping => github.com/iosmanthus/vmessping v0.2.4-0.20200127175013-c18f29ffebbc

replace v2ray.com/core => github.com/iosmanthus/v2ray-core v1.24.5-0.20200128120758-58c6fe125af7
