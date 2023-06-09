#!/bin/bash

#代码模板
tpl="../../scripts/template"

# echo "开始创建库模型：${dbname} "
goctl model mysql datasource -url="root:mysql8.0@tcp(127.0.0.1:3306)/bigger_stronger" -table="*"  -dir="./model" -cache=true


#构建api服务代码
#echo "start create api server >>>>>"
#goctl api go -api ./api/desc/*.api -dir ./api --style=go_zero --home=${tpl}