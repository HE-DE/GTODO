package utils

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func SnowflakeInit(startTime string, machineID int64) (err error) {
	var st time.Time
	// 格式化 1月2号下午3时4分5秒  2006年
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		fmt.Println(err)
		return
	}

	snowflake.Epoch = st.UnixNano() / 1e6
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
