package uuid

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
)

// var node

func NewSnowflake(startTime string, machineID int64) (*snowflake.Node, error) {
	var st time.Time
	// 格式化 1月2号下午3时4分5秒  2006年
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	snowflake.Epoch = st.UnixNano() / 1e6
	node, err := snowflake.NewNode(machineID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return node, err
}

// 生成 64 位的 雪花 ID
// func GenID() int64 {
// 	// return node.Generate().Int64()
// 	return node.Generate().Int64()
// }

// func main() {
// 	if err := Init("2021-12-03", 1); err != nil {
// 		fmt.Println("Init() failed, err = ", err)
// 		return
// 	}
// 	id := GenID()
// 	fmt.Println(id)
// }
