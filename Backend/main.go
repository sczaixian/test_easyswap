package main

import (
	"flag"
	"fmt"
	"reflect"

	"github.com/ProjectsTask/Backend/config"
	"github.com/ProjectsTask/Backend/service/svc"
	"github.com/zeromicro/go-zero/rest/router"
)

const (
	defaultConfigPath = "G:\\web3\\test_easyswap\\Backend\\config\\config.toml"
)

func PrintUnknownStruct(s interface{}) {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem() // 解引用指针
	}
	if val.Kind() != reflect.Struct {
		fmt.Println("输入非结构体类型")
		return
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)
		fmt.Printf("%s: %v\n", field.Name, value.Interface())
	}
}

func main() {
	conf := flag.String("conf", defaultConfigPath, "config file path")
	flag.Parse()
	c, err := config.UnmarshalConfig(*conf)
	if err != nil {
		panic(err)
	}
	PrintUnknownStruct(c)
	fmt.Printf("-----config------> ")

	for _, chain := range c.ChainSupported {
		if chain.ChainID == 0 || chain.Name == "" {
			panic("invalid chain_suffix config")
		}
	}

	serverCtx, err := svc.NewServiceContext(c)
	if err != nil {
		panic(err)
	}
	// Initialize router
	r := router.NewRouter(serverCtx)
	app, err := app.NewPlatform(c, r, serverCtx)
	if err != nil {
		panic(err)
	}
	app.Start()
}
