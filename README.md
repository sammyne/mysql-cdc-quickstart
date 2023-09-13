# MySQL CDC 快速开始

## 快速开始

### 1. 启动 MySQL 服务

```bash
bash scripts/mysqld.sh
```

### 2. 启动示例程序

```bash
mysqld_addr=`docker inspect -f '{{.NetworkSettings.Networks.bridge.IPAddress}}' mysqld`

go run main.go --remote $mysqld_addr
```

成功启动后，示例程序的日志样例如下
```bash
[2023/09/13 05:56:50] [info] binlogsyncer.go:173 create BinlogSyncer with config {ServerID:101 Flavor:mysql Host:192.168.10.9 Port:3306 User:root Password: Localhost: Charset:utf8 SemiSyncEnabled:false RawModeEnabled:false TLSConfig:<nil> ParseTime:false TimestampStringLocation:UTC UseDecimal:false RecvBufferSize:0 HeartbeatPeriod:0s ReadTimeout:0s MaxReconnectAttempts:0 DisableRetrySync:false VerifyChecksum:false DumpCommandFlag:0 Option:<nil> Logger:0xc00009a5a0 Dialer:0x61cb00 RowsEventDecodeFunc:<nil> DiscardGTIDSet:false}
[2023/09/13 05:56:50] [info] dump.go:203 skip dump, no mysqldump
[2023/09/13 05:56:50] [info] binlogsyncer.go:410 begin to sync binlog from position (, 0)
[2023/09/13 05:56:50] [info] sync.go:24 start sync binlog at binlog file (, 0)
[2023/09/13 05:56:50] [info] binlogsyncer.go:813 rotate to (mysql-bin.000001, 4)
[2023/09/13 05:56:50] [info] binlogsyncer.go:813 rotate to (mysql-bin.000002, 4)
[2023/09/13 05:56:50] [info] sync.go:63 received fake rotate event, next log name is mysql-bin.000001
[2023/09/13 05:56:50] [info] sync.go:65 log name changed, the fake rotate event will be handled as a real rotate event
[2023/09/13 05:56:50] [info] sync.go:91 rotate binlog to (mysql-bin.000001, 4)
[2023/09/13 05:56:50] [info] sync.go:63 received fake rotate event, next log name is mysql-bin.000002
[2023/09/13 05:56:50] [info] sync.go:65 log name changed, the fake rotate event will be handled as a real rotate event
[2023/09/13 05:56:50] [info] sync.go:91 rotate binlog to (mysql-bin.000002, 4)
[2023/09/13 05:56:50] [info] sync.go:228 table structure changed, clear table cache: mysql.time_zone
2023/09/13 05:56:50 DDL statement: TRUNCATE TABLE time_zone
[2023/09/13 05:56:50] [info] sync.go:228 table structure changed, clear table cache: mysql.time_zone_name
2023/09/13 05:56:50 DDL statement: TRUNCATE TABLE time_zone_name
[2023/09/13 05:56:50] [info] sync.go:228 table structure changed, clear table cache: mysql.time_zone_transition
2023/09/13 05:56:50 DDL statement: TRUNCATE TABLE time_zone_transition
[2023/09/13 05:56:50] [info] sync.go:228 table structure changed, clear table cache: mysql.time_zone_transition_type
2023/09/13 05:56:50 DDL statement: TRUNCATE TABLE time_zone_transition_type
[2023/09/13 05:56:50] [info] binlogsyncer.go:813 rotate to (mysql-bin.000003, 4)
[2023/09/13 05:56:50] [info] sync.go:63 received fake rotate event, next log name is mysql-bin.000003
[2023/09/13 05:56:50] [info] sync.go:65 log name changed, the fake rotate event will be handled as a real rotate event
[2023/09/13 05:56:50] [info] sync.go:91 rotate binlog to (mysql-bin.000003, 4)
```

### 3. 触发 binlog 更新

```bash
bash scripts/mysql-ops.sh
```

可得示例程序新增日志片段如下

```bash
[2023/09/13 05:58:50] [info] sync.go:228 table structure changed, clear table cache: world.a
2023/09/13 05:58:50 DDL statement: create table a(v int)
2023/09/13 05:58:50 Row event: insert world.a [[1] [2] [3]]
```

## 参考文献
- [CDC replication from mysql using go (golang)](https://adam-szpilewicz.pl/cdc-replication-from-mysql-using-go-golang)
