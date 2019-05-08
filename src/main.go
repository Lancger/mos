package main

import (
	"flag"
	"log"
	"mos/src/glo"
	"mos/src/server"
	"os"
)

var (
	help   bool
	config string
)

func init() {
	flag.BoolVar(&help, "h", false, "帮助信息")
	flag.StringVar(&config, "c", "", "配置文件")
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	}
	err := glo.ParseConfig(config)
	if err != nil {
		log.Panicln(err)
	}

	// init database
	glo.DbConnect()
	glo.RdsInit()
	defer glo.DbDisconnect()
	defer glo.RdsDisConnect()

	go server.Serve()
	select {}
}
