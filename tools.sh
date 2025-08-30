#!/bin/bash

if [ $# -eq 0 ]; then
  echo "没有输入参数"
  exit 1
fi

case $1 in
  "git")
    git pull
    git add .
    git commit -m "-"
    git push origin main
    echo "git 提交成功"
    ;;
  "go")
    cd backend/ && go mod tidy && cd ..
    cd base/ && go mod tidy && cd ..
    cd sync/ && go mod tidy && cd ..
  ;;
  *)
    echo " 无效参数 "
  ;;
esac
