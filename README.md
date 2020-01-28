# Use 3rd-party vmess services without fear

This tool aims to enable user use 3rd-party vmess services without any fear. User needs a private server overseas which providing a vmess service without TLS or HTTP support. When providing the configuration file to the vmess service, this tool could parse a vmess subscription links, randomly select a node and inject it into provided configuration file and establish a proxy chain, so we could send message to 3rd-party vmess provider securely.

## Usage

```sh
go build
# flag -inject follows a user provided configuration file path
# flag -subscribe follows a subscription link
# flag -output follows a output file path
./vmess-subscription -inject config.json -subscribe https://subscription.com -output injected.json
```