package main

import (
	"log"
)

const (
	serverAddress = ":3000"
)

func main() {
	srvCfg := &Config{
		ListenAddr: serverAddress,
	}

	server, err := NewServer(srvCfg)
	if err != nil {
		log.Fatal(err)
	}

	server.Start()
}
