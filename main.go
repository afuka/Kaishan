package main

import "flag"

var confPath string
var quit chan bool

func main(){
	flag.StringVar(&confPath, "c", "./conf/app.yaml", "-c set config file path")
	flag.Parse()

	quit = make(chan bool, 1) // 关闭标识

	initHandler()

}