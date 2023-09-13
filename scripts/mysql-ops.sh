#!/bin/bash

cmds='create table a(v int); insert into a(v) values(1),(2),(3);'

docker exec -it mysqld mysql -uroot -phello world -e "$cmds"
