#!/bin/bash

# var
# shellcheck disable=SC2046
outfile=output
module=kaishan

# func
function print() {
    echo "
                       _ooOoo_
                      o8888888o
                      88" . "88
                      (| -_- |)
                       O\ = /O
                   ____/'---'\____
                 .   ' \\| |// '.
                  / \\||| : |||// \\
                / _||||| -:- |||||- \\
                  | | \\\ - /// | |
                | \_| ''\---/'' |_/ |
                 \ .-\__ '-' ___/-. /
              ___'. .' /--.--\ '. . __
           ."" '< '.___\_<|>_/___.' >'"".
          | | : '- \'.;'\ _ /';.'/ - ' : | |
            \ \ '-. \_ __\ /__ _/ .-' / /
    ======'-.____'-.___\_____/___.-'____.-'======
                       '=---='
    .............................................
             佛祖保佑             永无BUG

    "
}

function build() {
  go build -o "${outfile}/bin/${module}"
  ret=$?
  if [ $ret -ne 0 ];then
      echo "===== module build failure ====="
      exit $ret
  fi
  cp -r conf "${outfile}"
  mkdir "${outfile}/logs"
  mkdir "${outfile}/data"

  echo -n "===== module build successfully! ====="
}

#print
build