package discovery

import "context"

type Instance struct {
	//service name
	Name string `json:"name"`
	// Hostname is hostname from docker.
	HostName string `json:"host_name"`
	// Addrs is the address of app instance
	// format: scheme://host
	Addrs []string `json:"addrs"`
	// Version is publishing version.
	Version string `json:"version"`
	// LastTs is instance latest updated timestamp
	LastTs int64 `json:"latest_timestamp"`
	// Metadata is the information associated with Addr, which may be used
	// to make load balancing decision.
	Metadata map[string]string `json:"metadata"`
	// Status instance status, eg: 1UP 2Waiting
	Status int64 `json:"status"`
}

// InstancesInfo instance info.
type InstancesInfo struct {
	Instances []*Instance `json:"instances"`
	LastTs    int64       `json:"latest_timestamp"`
}

type Builder interface {
	Build(serviceName string) Resolver                                                  //返回一个服务resolver，用于拉取该服务名下的节点
	Register(ctx context.Context, ins *Instance) (cancel context.CancelFunc, err error) //注册服务
	Close() error                                                                       //反注册服务，一般直接调Register返回的cancel即可
}

type Resolver interface {
	Fetch(ctx context.Context) (*InstancesInfo, bool) //拉取某个服务的全部节点
	Watch() <-chan struct{}                           //监听服务节点信息变化
	Close() error                                     //关闭resolver
}
