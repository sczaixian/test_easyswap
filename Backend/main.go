package main

import (
	"backend/config"
	"backend/service/svc"
	"flag"
)

const (
	defaultConfigPath = "config.json"
)

func main() {
	conf := flag.String("conf", defaultConfigPath, "config file path")
	flag.Parse()
	c, err := config.UnmarshalConfig(*conf)
	if err != nil {
		panic(err)
	}

	for _, chain := range c.ChainSupported {
		if chain.ChainID == 0 || chain.Name == "" {
			panic("invalid chain_suffix config")
		}
	}

	serverCtx, err := svc.NewServiceContext(c)
	if err != nil {
		panic(err)
	}
	print(serverCtx)
	//r := router.New(serverCtx)
	//app, err := server.NewApp(r, serverCtx)
}
