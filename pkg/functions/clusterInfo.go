package functions

import (

	"fmt"
	"strconv"

	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	// "github.com/junkeun-yi/kube-informer/pkg/config"
)

func (f FunctionSet) GetNodesInfo() {
	nodes, nodeGetErr := f.Client.CoreV1().Nodes().List(meta_v1.ListOptions{});
	if (nodeGetErr != nil) {
		fmt.Printf("%v", nodeGetErr)
	}
	nodePods := f.getNodePods(nodes)

	nodeMetrics := f.getNodeMetrics()

	for _, node := range nodes.Items {
		name := node.Name
		nodeMets := nodeMetrics[name]

		fmt.Printf("Node %s\n", name)

		pods := nodePods[name]

		cpuCores, memBytes := getNodeMetricSpecs(node)
		cpuUsage := nodeMets[0]
		memUsage := nodeMets[1]

		fmt.Printf("CPU Cores: %d    ", int(cpuCores))
		fmt.Printf("Usage: %f    ", cpuUsage)
		fmt.Printf("Utilization: %f%%\n", (cpuUsage / cpuCores) * 100)


		fmt.Printf("Memory Bytes: %d    ", int(memBytes))
		fmt.Printf("Usage: %f    ", memUsage)
		fmt.Printf("Utilization: %f%%\n", (memUsage / memBytes) * 100)

		fmt.Printf("%v\n", len(pods))
	}


}

func getNodeMetricSpecs(node core_v1.Node) (float64, float64){
	cpuDec := node.Status.Capacity.Cpu().AsDec()
	cpuStr := fmt.Sprintf("%v", cpuDec)
	cpuCores, cpuErr := strconv.ParseFloat(cpuStr, 64)
	if cpuErr != nil {
		fmt.Printf("%v", cpuErr)
	}
	memDec := node.Status.Capacity.Memory().AsDec()
	memStr := fmt.Sprintf("%v", memDec)
	memBytes, memErr := strconv.ParseFloat(memStr, 64)
	if memErr != nil {
		fmt.Printf("%v", memErr)
	}

	return cpuCores, memBytes
}



func (f FunctionSet) getNodePods(nodes *core_v1.NodeList) map[string][]*core_v1.Pod{
	nodePods := make(map[string][]*core_v1.Pod, 0)
	for _, node := range nodes.Items {
		nodePods[node.Name] = make([]*core_v1.Pod, 0)
	}

	pods, podGetErr := f.Client.CoreV1().Pods("").List(meta_v1.ListOptions{})
	if podGetErr != nil {
		fmt.Printf("%v", podGetErr)
	}
	for _, pod := range pods.Items {
		for _, node := range nodes.Items {
			if pod.Status.HostIP == node.Status.Addresses[0].Address {
				nodePods[node.Name] = append(nodePods[node.Name], &pod)
				break
			}
		}
	}

	return nodePods
}