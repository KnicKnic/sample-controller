package main

import (

// "github.co`m/yyyar/gobetween/config
// "github.com/yyyar/gobetween/core"
// "github.com/yyyar/gobetween/launch"
// "github.com/yyyar/gobetween/manager"
)

type LBProtocol string

const (
	LBProtocolUDP = "udp"
	LBProtocolTCP = "tcp"
)

type LbElem struct {
	Name       string
	IPAddress  string
	Protocol   LBProtocol
	Port       int
	TargetPort int
}

type BackendList struct {
	IPAddresses []string
}
