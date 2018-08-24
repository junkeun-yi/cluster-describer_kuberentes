package functions

import (

	"fmt"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	// "github.com/junkeun-yi/kube-informer/pkg/config"
	"github.com/junkeun-yi/kube-informer/pkg/config"
)

func (f FunctionSet) GetAllInfo() {
	fmt.Printf("========================================================================\n")
	f.getNodesInfo()
	fmt.Printf("========================================================================\n")
	f.getDeploymentsInfo()
	fmt.Printf("========================================================================\n")
}

func (f FunctionSet) getNodesInfo(){
	fmt.Printf("<<<<Nodes Info>>>:\n\n")

	nodes, nodeGetErr := f.Client.CoreV1().Nodes().List(meta_v1.ListOptions{});
	if (nodeGetErr != nil) {
		fmt.Printf("%v", nodeGetErr)
	}
	nodePods := f.getNodePodNames(nodes)

	nodeMetrics := f.getNodeMetrics()
	podMetrics := f.getPodMetrics()

	nodeHttpRequests := f.query(config.NodeHTTPRequestsPerMin)

	for _, node := range nodes.Items {
		name := node.Name
		nodeMets := nodeMetrics[name]
		pods := nodePods[name]

		fmt.Printf("Node: %s\n", name)

		cpuCores, memBytes := getNodeCapacities(node)
		cpuUsage := nodeMets[0]
		memUsage := nodeMets[1]

		fmt.Printf("CPU Cores: %d    ", int(cpuCores))
		fmt.Printf("Usage: %f    ", cpuUsage)
		fmt.Printf("Utilization: %f%%\n", (cpuUsage / cpuCores) * 100)
		fmt.Printf("Memory Bytes: %d    ", int(memBytes))
		fmt.Printf("Usage: %f    ", memUsage)
		fmt.Printf("Utilization: %f%%\n", (memUsage / memBytes) * 100)

		httpReqs := nodeHttpRequests[name]
		fmt.Printf("http requests per min: %f\n", httpReqs)

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
}

func (f FunctionSet) getDeploymentsInfo(){
	fmt.Printf("<<<<Deployments Info>>>:\n")
	fmt.Printf("(format: average usage of item (by pod) of requested capacity... percent usage)\n\n")

	podMetrics := f.getPodMetrics()

	deployments, depGetErr := f.Client.AppsV1beta2().Deployments("").List(meta_v1.ListOptions{})
	if depGetErr != nil {
		fmt.Printf("%v", depGetErr)
	}
	deploymentPods := f.getDeploymentPodNames(deployments)
	for _, dep := range deployments.Items {
		name := dep.Name
		pods := deploymentPods[name]

		fmt.Printf("%30s: ", name)
		fmt.Printf("%2v pods    ", len(pods))

		cpuRequest := getCPURequest(&dep)
		memRequest := getMemRequest(&dep)
		sumCPU := 0.0
		sumMem := 0.0
		for _, pod := range pods {
			sumCPU += podMetrics[pod][0]
			sumMem += podMetrics[pod][1]
		}
		avgCPU := sumCPU / float64(len(pods))
		avgMem := sumMem / float64(len(pods))
		fmt.Printf("cpu: %f of %f... %f%%   ", avgCPU, cpuRequest, (avgCPU / cpuRequest) * 100)
		fmt.Printf("mem: %d of %d... %f%%   ", int(avgMem), int(memRequest), (avgCPU / cpuRequest) * 100)
		fmt.Printf("\n")

	}

}