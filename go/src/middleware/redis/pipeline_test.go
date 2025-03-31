package redis

import (
	"fmt"
	"testing"
)

func Test_pipeline(t *testing.T) {
	client := initRedis()
	err := pipeline(client)
	if err != nil {
		fmt.Printf("err:%v", err)
	}
}
