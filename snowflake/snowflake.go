package snowflake

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

var snowNode *snowflake.Node

func Init() {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal("init snowflake failed", err)
	}
	snowNode = node
}

func NextID() int64 {
	return snowNode.Generate().Int64()
}
