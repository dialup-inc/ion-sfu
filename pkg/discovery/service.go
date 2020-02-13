package discovery

import (
	"encoding/json"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/pion/ion/pkg/log"
)

// ServiceWatchCallback .
type ServiceWatchCallback func(service string, nodes []Node)

//ServiceRegistry lib
type ServiceRegistry struct {
	Scheme string
	etcd   *Etcd
}

//Node service node info
type Node struct {
	Name string
	Info map[string]string
}

//NewServiceRegistry ServiceRegistry factory method
func NewServiceRegistry(endpoints []string, scheme string) *ServiceRegistry {
	r := &ServiceRegistry{}
	r.Scheme = scheme
	etcd, _ = newEtcd(endpoints)
	r.etcd = etcd
	return r
}

//RegisterServiceNode .
func (r *ServiceRegistry) RegisterServiceNode(serviceName string, node Node) error {
	if serviceName == "" {
		return fmt.Errorf("Service name must be non empty")
	}
	if node.Name == "" {
		return fmt.Errorf("Node name must be non empty")
	}
	go r.keepRegistered(serviceName, node)
	return nil
}

func (r *ServiceRegistry) keepRegistered(serviceName string, node Node) {
	for {
		err := r.etcd.keep(r.nodePath(serviceName, node.Name), encode(node.Info))
		if err != nil {
			log.Warnf("Registration got errors. Restarting. err=%s", err)
			time.Sleep(5 * time.Second)
		} else {
			log.Infof("Registration success!")
			return
		}
	}
}

//GetServiceNodes returns a list of active service nodes
func (r *ServiceRegistry) GetServiceNodes(serviceName string) ([]Node, error) {
	rsp, err := r.etcd.GetResponseByPrefix(r.servicePath(serviceName) + "/")
	if err != nil {
		return nil, err
	}
	nodes := make([]Node, 0)
	if len(rsp.Kvs) == 0 {
		log.Debugf("No services nodes were found under %s", r.servicePath(serviceName)+"/")
		return nodes, nil
	}

	for _, n := range rsp.Kvs {
		node := Node{}
		node.Name = string(n.Key)
		node.Info = decode(n.Value)
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func encode(m map[string]string) string {
	if m != nil {
		b, _ := json.Marshal(m)
		return string(b)
	}
	return ""
}

func decode(ds []byte) map[string]string {
	if ds != nil && len(ds) > 0 {
		var s map[string]string
		json.Unmarshal(ds, &s)
		return s
	}
	return nil
}

func (r *ServiceRegistry) servicePath(serviceName string) string {
	service := strings.Replace(serviceName, "/", "-", -1)
	return path.Join(r.Scheme, service)
}

func (r *ServiceRegistry) nodePath(serviceName string, nodeName string) string {
	service := strings.Replace(serviceName, "/", "-", -1)
	node := strings.Replace(nodeName, "/", "-", -1)
	return path.Join(r.Scheme, service, node)
}
