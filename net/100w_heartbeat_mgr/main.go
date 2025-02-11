package main

import (
	"container/list"
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	ShardCount      = 256             // 分片数量
	HeartbeatExpire = 5 * time.Second // 心跳超时时间
)

// 分片结构体
type HeartbeatShard struct {
	sync.RWMutex
	nodes map[string]*list.Element // 哈希表
	list  *list.List               // 双向链表
}

// 全局分片数组
var shards [ShardCount]*HeartbeatShard

// 心跳节点数据
type Node struct {
	IP        string
	Timestamp time.Time
}

func init() {
	// 初始化分片
	for i := 0; i < ShardCount; i++ {
		shards[i] = &HeartbeatShard{
			nodes: make(map[string]*list.Element),
			list:  list.New(),
		}
	}
}

// 获取分片索引
func getShard(ip string) int {
	hash := fnv32(ip) % uint32(ShardCount)
	return int(hash)
}

// 处理心跳
func ProcessHeartbeat(ip string) {
	shard := shards[getShard(ip)]

	shard.Lock()
	defer shard.Unlock()

	// 存在则更新
	if elem, exists := shard.nodes[ip]; exists {
		node := elem.Value.(*Node)
		node.Timestamp = time.Now()
		shard.list.MoveToBack(elem)
		return
	}

	// 新增节点
	node := &Node{IP: ip, Timestamp: time.Now()}
	elem := shard.list.PushBack(node)
	shard.nodes[ip] = elem
}

// 定时清理过期节点
func CleanupExpiredNodes() {
	for {
		time.Sleep(1 * time.Second)
		now := time.Now()

		for _, shard := range shards {
			shard.Lock()
			for {
				elem := shard.list.Front()
				if elem == nil {
					break
				}

				node := elem.Value.(*Node)
				if now.Sub(node.Timestamp) < HeartbeatExpire {
					break
				}

				// 删除过期节点
				delete(shard.nodes, node.IP)
				shard.list.Remove(elem)
				triggerOffline(node.IP) // 触发下线逻辑
			}
			shard.Unlock()
		}
	}
}

// UDP服务端
func StartUDPServer() {
	addr, _ := net.ResolveUDPAddr("udp", ":9999")
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, _, _ := conn.ReadFromUDP(buf)
		ip := string(buf[:n])
		fmt.Printf("rcv ip :%s\n", ip)
		go ProcessHeartbeat(ip) // 协程分片处理
	}
}

func main() {
	go StartUDPServer()
	go CleanupExpiredNodes()
	select {}
}

// 哈希函数
func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	for i := 0; i < len(key); i++ {
		hash *= 16777619
		hash ^= uint32(key[i])
	}
	return hash
}

func triggerOffline(ip string) {
	// 处理主机下线逻辑
	fmt.Printf("offline ip:%s\n", ip)
}

