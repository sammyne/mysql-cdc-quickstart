package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
	flag "github.com/spf13/pflag"
)

var mysqldAddr string

func main() {
	flag.Parse()

	if !strings.Contains(mysqldAddr,":") {
		mysqldAddr+=":3306"
	}

    // Define default dump configuration
    dumpCfg := canal.DumpConfig{
        ExecutionPath: "",
    }

    // Set up canal to connect to MySQL database
    cfg := canal.NewDefaultConfig()
    cfg.Addr = mysqldAddr
    cfg.User = "root"
    cfg.Password = "hello"
    cfg.ServerID = 101
    cfg.Flavor = "mysql"
    cfg.Dump = dumpCfg
    cfg.IncludeTableRegex = []string{"world*"}

    // Create an instance of the eventHandler struct
    eventHandler := &eventHandler{}

    // Create a new Canal instance with the specified configuration and event handler
    can, err := canal.NewCanal(cfg)
    if err != nil {
        log.Fatalf("Failed to create Canal: %v", err)
    }

    // Register the event handler with the Canal instance
    can.SetEventHandler(eventHandler)

    // Start canal and subscribe to all binlog events
    err = can.Run()
    if err != nil {
        log.Fatalf("Failed to start Canal: %v", err)
    }

    // Wait for SIGINT or SIGTERM signals to stop the program
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    <-sigs

    // Stop canal and exit program
    can.Close()
    log.Println("Canal stopped")
}

func init() {
	flag.StringVar(&mysqldAddr, "remote", "", "MySQL 服务地址，格式为 'IP:端口'")
}

// Define a custom event handler to process binlog events
type eventHandler struct{}

func (h *eventHandler) String() string {
    //TODO implement me
    panic("implement me")
}

func (h *eventHandler) OnRotate(header *replication.EventHeader, r *replication.RotateEvent) error {
    // Do nothing
    return nil
}

func (h *eventHandler) OnTableChanged(header *replication.EventHeader, schema string, table string) error {
    // Do nothing
    return nil
}

func (h *eventHandler) OnDDL(header *replication.EventHeader, nextPos mysql.Position, queryEvent *replication.QueryEvent) error {
    // Print the DDL statement to the console
    log.Printf("DDL statement: %v", string(queryEvent.Query))

    return nil
}

func (h *eventHandler) OnRow(e *canal.RowsEvent) error {
    // Print the row event to the console
    log.Printf("Row event: %v", e)

    return nil
}

func (h *eventHandler) OnGTID(*replication.EventHeader, mysql.GTIDSet) error {
    // Do nothing
    return nil
}

func (h *eventHandler) OnPosSynced(header *replication.EventHeader, pos mysql.Position, set mysql.GTIDSet, force bool) error {
    // Do nothing
    return nil
}

func (h *eventHandler) OnXID(*replication.EventHeader, mysql.Position) error {
    // Do nothing
    return nil
}

func (h *eventHandler) OnUnmarshal(data []byte) (interface{}, error) {
    // Do nothing
    return nil, nil
}

func (h *eventHandler) OnRawEvent(event *replication.BinlogEvent) error {
    // Do nothing
    return nil
}
