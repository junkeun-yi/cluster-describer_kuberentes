package functions

import (

	"fmt"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	// "github.com/junkeun-yi/kube-informer/pkg/config"
)

func (f FunctionSet) GetAllInfo() {


	deployments, depGetErr := f.Client.AppsV1beta2().Deployments("").List(meta_v1.ListOptions{})
	if depGetErr != nil {
		fmt.Printf("%v", depGetErr)
	}
	deploymentPods := f.getDeploymentPodNames(deployments)
	for _, dep := range deployments.Items {
		
	}

}

func (f FunctionSet) getNodesInfo(){
	nodes, nodeGetErr := f.Client.CoreV1().Nodes().List(meta_v1.ListOptions{});
	if (nodeGetErr != nil) {
		fmt.Printf("%v", nodeGetErr)
	}
	nodePods := f.getNodePodNames(nodes)

	nodeMetrics := f.getNodeMetrics()
	podMetrics := f.getPodMetrics()

	fmt.Printf("========================================================================\n")
	for _, node := range nodes.Items {
		name := node.Name
		nodeMets := nodeMetrics[name]

		fmt.Printf("Node %s\n", name)

		pods := nodePods[name]

		cpuCores, memBytes := getNodeCapacities(node)
		cpuUsage := nodeMets[0]
		memUsage := nodeMets[1]

		fmt.Printf("CPU Cores: %d    ", int(cpuCores))
		fmt.Printf("Usage: %f    ", cpuUsage)
		fmt.Printf("Utilization: %f%%\n", (cpuUsage / cpuCores) * 100)
		fmt.Printf("Memory Bytes: %d    ", int(memBytes))
		fmt.Printf("Usage: %f    ", memUsage)
		fmt.Printf("Utilization: %f%%\n", (memUsage / memBytes) * 100)

		fmt.Printf("Utilization by pod in node:\n")
		fmt.Printf("----------------------------------------------------------------------\n")
		for _, pod := range pods {
			cpu := podMetrics[pod][0]
			mem := podMetrics[pod][1]
			fmt.Printf("%70s:    ", pod)
			fmt.Printf("CPU: %f, %f%% of node cpu usage   ", cpu, (cpu / cpuUsage) * 100)
			fmt.Printf("Memory Bytes: %f, %f%% of node mem usage    ", mem, (mem / memUsage) * 100)
			fmt.Printf("\n")
		}
		fmt.Printf("----------------------------------------------------------------------\n")
	}
	fmt.Printf("========================================================================\n")
}