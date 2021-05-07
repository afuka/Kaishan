package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"kaishan/core/handlers/log"
	"os"
	"os/signal"
	"path"
	"syscall"
)

var confPath string
var quit chan bool

func main(){
	flag.StringVar(&confPath, "c", "output/conf/app.yaml", "-c set config file path")
	flag.Parse()

	setPidFile() // 进程号

	quit = make(chan bool, 1) // 关闭标识

	initHandler()

	// 停止应用
	ch := make(chan os.Signal)
	signal.Notify(ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
		syscall.SIGUSR1,
		syscall.SIGPIPE,
	)

	for {
		sig := <-ch
		switch sig {
		case syscall.SIGUSR1:
			// 重新加载配置文件
			continue
		case syscall.SIGPIPE:
			continue
		default:
			log.Info("kaishan got a signal", log.Field{"sign": sig.String()})
		}
		break
	}

	closeHandler()
	remPidFile()
}

func setPidFile() {
	file := path.Join("output/app.pid")
	pid := os.Getpid()
	ioutil.WriteFile(file, []byte(fmt.Sprintf("%d", pid)), 0644)
}

func remPidFile() {
	file := path.Join("output/app.pid")
	os.Remove(file)
}