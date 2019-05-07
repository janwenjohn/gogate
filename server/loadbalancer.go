package server

import "sync/atomic"

// 负载均衡接口
type LoadBalancer interface {
	// 从instance中选一个对象返回
	Choose(instances []*InstanceInfo) *InstanceInfo
}

// 轮询均衡器实现
type RoundRobinLoadBalancer struct {
	index	int64
}

func (lb *RoundRobinLoadBalancer) Choose(instances []*InstanceInfo) *InstanceInfo {
	total := len(instances)

	target := lb.index % int64(total)
	if target < 0 {
		target = target * -1
	}

	atomic.AddInt64(&lb.index, 1)

	return instances[target]
}

