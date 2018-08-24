package functions

import (
	"fmt"
	"strconv"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// prints the utilization for nodes
func (f FunctionSet) getNodeMetrics() map[string][]float64{
	mets := make(map[string][]float64, 0)

	metrics, err := f.MetClient.MetricsV1beta1().NodeMetricses().List(meta_v1.ListOptions{});
	if (err != nil) {
		fmt.Printf("%v", err)
	}
	for _, metric := range metrics.Items {

		cpuDec := metric.Usage.Cpu().AsDec()
		cpuStr := fmt.Sprintf("%v", cpuDec)
		cpu, cpuErr := strconv.ParseFloat(cpuStr, 64)
		if cpuErr != nil {
			fmt.Printf("%v", cpuErr)
		}

		memDec := metric.Usage.Memory().AsDec()
		memStr := fmt.Sprintf("%v", memDec)
		mem, memErr := strconv.ParseFloat(memStr, 64)
		if memErr != nil {
			fmt.Printf("%v", memErr)
		}

		mets[metric.ObjectMeta.Name] = []float64{cpu, mem}
	}

	return mets
}

// returns (CPU USAGE, MEMORY USAGE)
func (f FunctionSet) getPodMetrics() map[string][]float64{
	mets := make(map[string][]float64, 0)

	metrics, err := f.MetClient.MetricsV1beta1().PodMetricses("").List(meta_v1.ListOptions{});
	if (err != nil) {
		fmt.Printf("%v", err)
	}
	for _, metric := range metrics.Items {
		cpu := 0.0
		mem := 0.0
		for _, cont := range metric.Containers {
			cpuDec := cont.Usage.Cpu().AsDec()
			cpuStr := fmt.Sprintf("%v", cpuDec)
			cpuContainer, cpuErr := strconv.ParseFloat(cpuStr, 64)
			if cpuErr != nil {
				fmt.Printf("%v", cpuErr)
			}
			memDec := cont.Usage.Memory().AsDec()
			memStr := fmt.Sprintf("%v", memDec)
			memContainer, memErr := strconv.ParseFloat(memStr, 64)
			if memErr != nil {
				fmt.Printf("%v", memErr)
			}
			cpu += cpuContainer
			mem += memContainer
		}

		mets[metric.ObjectMeta.Name] = []float64{cpu, mem}
	}

	return mets
}