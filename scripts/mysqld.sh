#!/bin/bash

repo_tag=mysql:8.0.34-debian

opts="--log-bin=mysql-bin --server-id=1 --binlog-format=ROW --binlog-row-image=FULL --expire-logs-days=10"
opts="$opts --max_allowed_packet=32505856 --gtid-mode=ON --enforce-gtid-consistency=ON --sql_mode=ALLOW_INVALID_DATES" 

name=mysqld

docker stop $name

docker run -td --rm                 \
  -e MYSQL_DATABASE='world'         \
  -e MYSQL_ROOT_PASSWORD='hello'    \
  -e MYSQL_ROOT='root'              \
  --name $name                      \
  -p 9306:3306                      \
  $repo_tag                         \
    $opts
