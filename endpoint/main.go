package main

import "runtime"

var server string = "http://127.0.0.1:8443" // TODO: client config file

func register() {
	//var ip string = "10.0.1.39" // TODO: get ip with library(?) system command (?)
	//var hostname string = "bingus!"
}

func main() {
	println("The stones are capped.... Or whatever")
	println("We'll be connecting to " + server)
	println("My hostname is " + GetHostname())
	//println("My IP is : " + GetHostIp())
	println("My host OS is " + runtime.GOOS)
	println("Version info: " + GetOSSubtype())
}
