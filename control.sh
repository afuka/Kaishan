#!/bin/bash

# ./control action args
action=$1
env=$2

# env
outfile=output
module=kaishan

case $action in
    "start" )
        if [ "${env}" == "prod" ]; then
            config="${outfile}/conf/app_prod.yaml"
        else
            config="${outfile}/conf/app_dev.yaml"
        fi
        # 启动服务, 以前台方式启动, 否则无法托管
        exec "${outfile}/bin/${module}" -c="${config}"
        ;;
    "stop" )
        # 停止应用
        kill -s SIGSTOP $(<"${outfile}/app.pid")
        ;;
    "reload" )
        # 从新加载配置文件
        ;;
    * )
        # 非法命令, 已非0码退出
        echo "unknown command"
        exit 1
        ;;
esac
