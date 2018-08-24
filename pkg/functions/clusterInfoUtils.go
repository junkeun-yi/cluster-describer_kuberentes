package functions

import (
	"fmt"
	"strconv"

	core_v1 "k8s.io/api/core/v1"
	apps_v1beta2 "k8s.io/api/apps/v1beta2"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func getNodeCapacities(node core_v1.Node) (float64, float64){
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



func (f FunctionSet) getNodePodNames(nodes *core_v1.NodeList) map[string][]string{
	nodePods := make(map[string][]string, 0)
	for _, node := range nodes.Items {
		nodePods[node.Name] = make([]string, 0)
	}

	pods, podGetErr := f.Client.CoreV1().Pods("").List(meta_v1.ListOptions{})
	if podGetErr != nil {
		fmt.Printf("%v", podGetErr)
	}
	for _, pod := range pods.Items {
		for _, node := range nodes.Items {
			if pod.Status.HostIP == node.Status.Addresses[0].Address {
				nodePods[node.Name] = append(nodePods[node.Name], pod.Name)
				break
			}
		}
	}

	return nodePods
}

func (f FunctionSet) getDeploymentPodNames(deployments *apps_v1beta2.DeploymentList) map[string][]string{
	deploymentPods := make(map[string][]string, 0)

	pods, podGetErr := f.Client.CoreV1().Pods("").List(meta_v1.ListOptions{})
	if podGetErr != nil {
		fmt.Printf("%v", podGetErr)
	}
	podItems := pods.Items
	for _, dep := range deployments.Items {
		selector, selectorErr := meta_v1.LabelSelectorAsSelector(dep.Spec.Selector)
		if selectorErr != nil {
			fmt.Printf("%v", selectorErr)
		}
		for _, pod := range podItems {
			if selector.Matches(labels.Set(pod.GetLabels())) {
				deploymentPods[dep.Name] = append(deploymentPods[dep.Name], pod.Name)
			}
		}
	}

	return deploymentPods
}