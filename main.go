package main

import (
	"context"
	zabbix "github.com/0xdeface/zabbix/sender"
)
func main() {
	zabbixSender := zabbix.NewZabbixSender("127.0.0.1", "10051")
	ctx := context.Background()
	go zabbixSender.Start(ctx)
	zabbixSender.MsgCh <- zabbix.CreateMessage("MY_ZABBIX_ITEM_HOST", "MY_DATA_KEY", "MY_DATA_VAL")
}
